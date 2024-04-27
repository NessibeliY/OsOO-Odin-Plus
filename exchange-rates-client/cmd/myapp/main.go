package main

import (
	"exchange-rates-client/nyeltay/internal/handlers"
	"exchange-rates-client/nyeltay/internal/requestapi"
	"exchange-rates-client/nyeltay/internal/server"
	"exchange-rates-client/nyeltay/internal/templates"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	templateCache, err := templates.NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	apiRequest := requestapi.New()
	handler := handlers.NewApplication(templateCache, apiRequest)

	http.HandleFunc("/ws", handleWebSocket)

	go func() {
		log.Fatal(server.RunServer(handler.Routes()))
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := apiRequest.Api()
				if err != nil {
					log.Println("error refreshing API data:", err)
					continue
				}
			}
		}
	}()

	select {}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading to WebSocket", err)
		return
	}
	defer conn.Close()
}
