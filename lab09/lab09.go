package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	maxComments := flag.Int("max",10,"Max number of comments to show")
	// defaultUsage := flag.Usage
	
	flag.Parse()
	// fmt.Println("Max Comments:", *maxComments)
	c := colly.NewCollector()
	
	count := 0
	for count =0 ; count < *maxComments ; count++{
		result := fmt.Sprintf("#main-content > div:nth-child(%d)",count+7)
		currentCount := count
		c.OnHTML(result, func(e *colly.HTMLElement) {
			name := e.DOM.Find("span.f3.hl.push-userid").Text() 
			content := e.DOM.Find("span.f3.push-content").Text()
			time := e.DOM.Find("span.push-ipdatetime").Text()
			name = strings.TrimSpace(name)
			content = strings.TrimSpace(content)
			time = strings.TrimSpace(time)
			fmt.Printf("%d. 名字：%s，留言%s，時間： %s\n",currentCount+1 , name, content, time)

	
		})	
	}
	c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")




	

}
