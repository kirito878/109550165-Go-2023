package main

import (
	"bufio"
	"context"
	// "fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/reactivex/rxgo/v2"
)

type client chan<- string // an outgoing message channel

var (
	entering      = make(chan client)
	leaving       = make(chan client)
	messages      = make(chan rxgo.Item) // all incoming client messages
	ObservableMsg = rxgo.FromChannel(messages)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	MessageBroadcast := ObservableMsg.Observe()
	for {
		select {
		case msg := <-MessageBroadcast:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg.V.(string)
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func clientWriter(conn *websocket.Conn, ch <-chan string) {
	for msg := range ch {
		conn.WriteMessage(1, []byte(msg))
	}
}

func wshandle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "你是 " + who + "\n"
	messages <- rxgo.Of(who + " 來到了現場" + "\n")
	entering <- ch

	defer func() {
		log.Println("disconnect !!")
		leaving <- ch
		messages <- rxgo.Of(who + " 離開了" + "\n")
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		messages <- rxgo.Of(who + " 表示: " + string(msg))
	}
}

func filterSwearWords(item interface{}) bool {
	// 讀取 swear_word.txt 中的髒話列表
	swearWords := readSwearWordsFromFile("swear_word.txt")

	// 檢查訊息是否包含髒話
	message := item.(string)
	for _, word := range swearWords {
		if strings.Contains(message, word) {
			return false// 如果包含髒話，過濾掉
		}
	}
	return true // 沒有髒話，保留
}

// 從檔案中讀取髒話列表
func readSwearWordsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		// 可以在這裡進行錯誤處理，不過不再返回錯誤
		return nil
	}
	defer file.Close()

	var swearWords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		swearWords = append(swearWords, scanner.Text())
	}

	return swearWords
}
func modifySensitiveNames(_ context.Context, item interface{}) (interface{}, error) {
	// 讀取 sensitive_name.txt 中的敏感名字列表
	sensitiveNames := readSensitiveNamesFromFile("sensitive_name.txt")

	// 修改訊息中的敏感名字
	message := item.(string)
	for name, replacement := range sensitiveNames {
		message = strings.ReplaceAll(message, name, replacement)
	}
	return message, nil
}


// 從檔案中讀取敏感名字列表
func readSensitiveNamesFromFile(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		// 可以在這裡進行錯誤處理，不過不再返回錯誤
		return nil
	}
	defer file.Close()

	sensitiveNames := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		// 將名字的第二個字修改為 '*'
		if len(name) >= 2 {
			sensitiveNames[name] = string(name[0:3]) + "*" + string(name[6:])
		}
		// fmt.Printf("Name: %s\n", sensitiveNames)
	}

	return sensitiveNames
}
func InitObservable() {
	// TODO: Please create an Observable to handle the messages
	/*
		ObservableMsg = ObservableMsg.Filter(...) ... {
		}).Map(...) {
			...
		})
	*/
	ObservableMsg = ObservableMsg.Filter(rxgo.Predicate(filterSwearWords)).Map((modifySensitiveNames))
}

func main() {
	InitObservable()
	go broadcaster()
	http.HandleFunc("/wschatroom", wshandle)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("server start at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
