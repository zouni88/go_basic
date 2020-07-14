package main

import (
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"os"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		println("收到了什么？")
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/","add user", func(s socketio.Conn,username string) {
		println("add user",username)
		//sio.emit('login', {'username': username, 'numUsers': len(datas)}, room='123')
		//jsons
		ack := AckUserInfo{Username: username ,NumUsers: 2}
		res ,_ := json.Marshal(ack)
		println(string(res))
		//nsp := s.Namespace()
		s.Emit("login",string(res))
		s.Emit("user joined",string(res))
	})

	server.OnEvent("/","new message", func(s socketio.Conn,msg string) string {
		println("new message？")
		s.SetContext(msg)

		return "new message " + msg
	})

	server.OnEvent("/","login", func(s socketio.Conn,msg string) string {
		println("login？")
		s.SetContext(msg)
		return "new message " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()
	pwd, _ := os.Getwd()
	println(pwd)
	http.Handle("/socket.io/", server)
	dir := http.Dir("D:\\WorkProject\\go\\go_basic\\socket.io\\static")
	http.Handle("/", http.FileServer(dir))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
