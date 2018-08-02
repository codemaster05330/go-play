package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("server started at localhost:8080")

	http.Handle("/", http.HandlerFunc(handler))
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("req.URL.Path=%s\n", req.URL.Path)
	fmt.Printf("req.URL.RawQuery=%s\n", req.URL.RawQuery)
	printQueryParam(req.URL.Query())
	printHeaders(req.Header)

	writePods(rw)
}

func printQueryParam(params url.Values) {
	fmt.Printf("query_params %d\n", len(params))
	for key, value := range params {
		fmt.Printf("key=%s, value=%s\n", key, value)
	}
}

func printHeaders(headers http.Header) {
	fmt.Printf("headers %d\n", len(headers))
	for key, value := range headers {
		fmt.Printf("key=%s, value=%s\n", key, value)
	}
}

func writePods(rw http.ResponseWriter) {
	res := `{"kind":"PodList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/namespaces/john-preview-stage/pods","resourceVersion":"300455571"},"items":[{"metadata":{"name":"cpl-tools-1-62d9w","generateName":"cpl-tools-1-","namespace":"nvirani-preview-stage","uid":"e2a13ac3-9432-11e8-bf1d-029a5a55534e","creationTimestamp":"2018-07-30T19:58:10Z"}},{"metadata":{"name":"dummy-app1-1-6tb8l","generateName":"dummy-app1-1-","namespace":"john-preview-stage","uid":"772af236-9432-11e8-bf1d-029a5a55534e","creationTimestamp":"2018-07-30T19:55:10Z"}}]}`
	rw.Write([]byte(res))
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
