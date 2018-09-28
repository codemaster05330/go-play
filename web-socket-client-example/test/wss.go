package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	log.Printf("wss example")

	path := "wss://api.starter-us-east-2a.openshift.com/oapi/v1/namespaces/nvirani-preview/buildconfigs"
	config, err := websocket.NewConfig(path, "https://api.starter-us-east-2a.openshift.com")
	if err != nil {
		log.Fatal(err)
	}
	config.Header = http.Header{
		"Authorization": {"Basic M_i6IWyprxctakDqrB0XphhhG2bs-EsIg43YqPyeTCU"},
	}
	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
}
