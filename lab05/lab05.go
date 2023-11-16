package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"html/template"
	"strings"
	"time"
	"encoding/json"
	"github.com/joho/godotenv"
)

// TODO: Please create a struct to include the information of a video
type Output struct {
    Title string
    Id  string
	ChannelTitle string
	LikeCount string
	ViewCount string
	PublishedAt string
	CommentCount string
}


func YouTubePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Get API token from .env file
	// TODO: Get video ID from URL query `v`
	// TODO: Get video information from YouTube API
	// TODO: Parse the JSON response and store the information into a struct
	// TODO: Display the information in an HTML page through `template`
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")	
	baseURL := "https://www.googleapis.com/youtube/v3/videos"
	videoID := r.URL.Query().Get("v")
	if videoID == "" {
		http.ServeFile(w, r, "error.html")
		return
	}	
	url := fmt.Sprintf("%s?part=statistics,snippet&id=%s&key=%s", baseURL, videoID, youtubeAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return 
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return 
	}

	items, ok := data["items"].([]interface{})
	if !ok || len(items) == 0 {
		http.ServeFile(w, r, "error.html")
		return 
	}
	statistics, ok := items[0].(map[string]interface{})["statistics"].(map[string]interface{})
	if !ok {
		http.ServeFile(w, r, "error.html")
		return 
	}
	views := formatNumber(statistics["viewCount"].(string))
	likes := formatNumber(statistics["likeCount"].(string))
	comments := formatNumber(statistics["commentCount"].(string))
	title := items[0].(map[string]interface{})["snippet"].(map[string]interface{})["title"].(string)
	channelTitle := items[0].(map[string]interface{})["snippet"].(map[string]interface{})["channelTitle"].(string)
	publishedAt := items[0].(map[string]interface{})["snippet"].(map[string]interface{})["publishedAt"].(string)
	parsedTime, err := time.Parse(time.RFC3339, publishedAt)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return 
	}
	formattedDate := parsedTime.Format("2006年01月02日")
	var out = Output{
		Title :title,
		Id :videoID,
		ChannelTitle :channelTitle,
		LikeCount :likes,
		ViewCount :views,
		PublishedAt :formattedDate,
		CommentCount :comments,
	}
	template.Must(template.ParseFiles("index.html")).Execute(w,out)	
	// fmt.Fprintf(w, videoID)
	// fmt.Fprintf(w, views)
	// fmt.Fprintf(w, likes)
	// fmt.Fprintf(w, comments)
	// fmt.Fprintf(w, title)
	// fmt.Fprintf(w, channelTitle)
	// fmt.Fprintf(w, formattedDate)
	
}
func formatNumber(number string) string {
	// Format the number with commas every 3 digits
	parts := strings.Split(number, "")
	result := ""
	for i := len(parts) - 1; i >= 0; i-- {
		result = parts[i] + result
		if (len(parts)-i)%3 == 0 && i != 0 {
			result = "," + result
		}
	}
	return result
}
func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
