package main

import (
	"log"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/namespaces/ns.p;type=che/pods/p;space=997f146d-b0f4-4a97-ab20-6414878d9508?w=true", nil)
	log.Printf("req.URL.Path=%s\n", req.URL.Path)
	log.Printf("req.URL.RawQuery=%s\n", req.URL.RawQuery)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Err:", err)
		return
	}
	log.Printf("res.StatusCode=%d\n", res.StatusCode)
}
