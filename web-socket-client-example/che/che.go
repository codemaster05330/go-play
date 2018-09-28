package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	addr := "api.starter-us-east-2a.openshift.com"

	// path := "/oapi/v1/namespaces/nvirani-preview/buildconfigs"
	// query := "watch=true"

	// path := "/oapi/v1/namespaces/nvirani-preview/builds"
	// query := "watch=true"

	path := "/api/v1/namespaces/nvirani-preview-che/pods"
	// query := "fieldSelector=metadata.name=workspacej3rl5cgarxagfcwq&watch=true"
	query := "watch=true"

	u := url.URL{Scheme: "wss", Host: addr, Path: path, RawQuery: query}
	reqHeader := http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", getToken())},
		"Origin":        {"localhost:8080"},
	}

	log.Printf("Connecting to url:%s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), reqHeader)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Printf("Connect OK.")

	for i := 0; ; i++ {
		start := time.Now()
		log.Printf("Reading msg-%d", (i + 1))
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Printf("Got err:%s, after=%f sec", err, time.Since(start).Seconds())
			break
		}
		// log.Printf("msg-%d:%s", (i + 1), string(msg))
		log.Printf("Got msg-%d, with len=%d, after=%f sec", (i + 1), len(msg), time.Since(start).Seconds())
	}
}

func getToken() string {
	f, _ := os.Open("/home/nvirani/temp/token.txt")
	b, _ := ioutil.ReadAll(f)
	return string(b)
}
