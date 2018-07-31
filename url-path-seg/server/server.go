package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("server started at localhost:8080")

	http.Handle("/", http.HandlerFunc(handler))
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	log.Printf("req.URL.Path=%s\n", req.URL.Path)
	log.Printf("req.URL.RawQuery=%s\n", req.URL.RawQuery)
	log.Printf("w=%s\n", req.URL.Query().Get("w"))
}

/*
sample run

** server output **
[user1@localhost server]$ go run server.go
2018/07/31 17:15:55 server started at localhost:8080
2018/07/31 17:15:59 req.URL.Path=/api/v1/namespaces/ns;type=che/pods/p;space=997f146d-b0f4-4a97-ab20-6414878d9508
2018/07/31 17:15:59 req.URL.RawQuery=w=true
2018/07/31 17:15:59 w=true

** client output **
[user1@localhost client]$ go run client.go
2018/07/31 17:15:59 req.URL.Path=/api/v1/namespaces/ns;type=che/pods/p;space=997f146d-b0f4-4a97-ab20-6414878d9508
2018/07/31 17:15:59 req.URL.RawQuery=w=true
2018/07/31 17:15:59 res.StatusCode=200

*/
