package main

import (
	"log"
	"net/http"
)

func httpEx() {
	log.Println("http example")

	funcHandlerDemo()
	typeHandlerDemo()

	http.ListenAndServe(":8080", nil) // nil means use http.DefaultServerMux
}

func funcHandlerDemo() {
	http.HandleFunc("/foo1", funcHandler)
}

func funcHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from FuncHandler"))
}

func typeHandlerDemo() {
	http.Handle("/foo2", myHandlerType{})
}

type myHandlerType struct {
}

func (h myHandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from TypeHandler"))
}
