package main

import (
	"log"
	"net/http"
	"strings"

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

	fs := http.FileServer(http.Dir("./example/dist"))
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
		http.ServeFile(w, r, "./example/dist/index.html")
	})

	log.Println("running...")
	return http.ListenAndServe("0.0.0.0:5000", nil)
}
