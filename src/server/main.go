package main

import (
	"fmt"
	"log"
	"net/http"

	r "gopkg.in/dancannon/gorethink.v2"
)

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	router.Handle("user edit", editUser)
	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)

	router.Handle("message add", addChannelMessage)
	router.Handle("message subscribe", subscribeChannelMessage)
	router.Handle("message unsubscribe", unsubscribeChannelMessage)
	http.Handle("/ws", router)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.ListenAndServe(":4000", nil)

	fmt.Println("go server rodando na porta 4000")
}
