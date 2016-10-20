package main

import (
	_ "blog_go/routers"

	"github.com/astaxie/beego"

	_ "blog_go/config"
	"blog_go/dao"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("receive: %s\n", msg[:n])
	sendMsg := "hello client"
	m, err := ws.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("send: ", sendMsg, m)

}

func main() {
	defer dao.Database.Close()
	beego.Run()
	go func() {
		http.Handle("/echo", websocket.Handler(echoHandler))
		http.Handle("/", http.FileServer(http.Dir(".")))

		err := http.ListenAndServe(":8088", nil)
		if err != nil {
			fmt.Println(err)
		}
	}()
}
