package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// Help: example run with arguments
// go run main.go api.starter-us-east-2a.openshift.com nvirani-preview xxxx
func main() {
	osoServer := os.Args[1]    // api.starter-us-east-2a.openshift.com
	osoNamespace := os.Args[2] // nvirani-preview
	osoToken := os.Args[3]

	addr := osoServer
	// path := fmt.Sprintf("/oapi/v1/namespaces/%s/buildconfigs", osoNamespace)
	path := fmt.Sprintf("/oapi/v1/namespaces/%s/builds", osoNamespace)

	u := url.URL{Scheme: "wss", Host: addr, Path: path, RawQuery: "watch=true"}
	reqHeader := http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", osoToken)},
		"Origin":        {"localhost:8080"},
	}

	log.Printf("Connecting to url:%s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), reqHeader)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for i := 0; ; i++ {
		start := time.Now()
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Printf("Got err:%s, after=%f sec", err, time.Since(start).Seconds())
			break
		}
		// log.Printf("msg-%d:%s", (i + 1), string(msg))
		log.Printf("Got msg-%d, with len=%d, after=%f sec", (i + 1), len(msg), time.Since(start).Seconds())
	}
}
