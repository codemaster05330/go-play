package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func logMiddlewareEx() {
	log.Println("log middleware ..")

	handler := http.HandlerFunc(greetHandler)

	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	http.Handle("/", handlers.LoggingHandler(logFile, handler))

	http.ListenAndServe(":8080", nil)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("greet handler")
	w.Write([]byte("Hello World !!!\n"))
}
