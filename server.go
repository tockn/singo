package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/tockn/singo/handler"
	"github.com/tockn/singo/manager"
	"github.com/tockn/singo/repository/mem"
)

func serve(addr string) error {
	roomRepo := mem.NewRoomRepository()
	man := manager.NewManager(roomRepo)
	han := handler.NewHandler(man)

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		han.CreateConnection(w, r)
	})

	log.Println("running...")
	return http.ListenAndServe(addr, nil)
}

func serveWithDemo(addr string) error {
	roomRepo := mem.NewRoomRepository()
	man := manager.NewManager(roomRepo)
	han := handler.NewHandler(man)

	fs := http.FileServer(http.Dir("./demo/dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		if r.URL.Path == "/connect" {
			han.CreateConnection(w, r)
			return
		}
		sp := strings.Split(r.URL.Path, "/")
		if len(sp) > 2 && (sp[1] == "js" || sp[1] == "css") {
			fs.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, "./demo/dist/index.html")
	})

	log.Println("running...(with demo)")
	return http.ListenAndServe(addr, nil)
}
