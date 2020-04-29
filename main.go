package main

import (
	"log"
	"net/http"

	"github.com/tockn/singo/handler"
	"github.com/tockn/singo/manager"
	"github.com/tockn/singo/repository/mem"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	roomRepo := mem.NewRoomRepository()
	man := manager.NewManager(roomRepo)
	han := handler.NewHandler(man)

	http.HandleFunc("/connect", han.CreateConnection)

	log.Println("running...")
	return http.ListenAndServe("0.0.0.0:5000", nil)
}
