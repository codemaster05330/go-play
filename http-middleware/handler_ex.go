package main

import (
	"log"
	"net/http"
)

func handlerEx() {
	log.Println("handler example start ..")

	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	http.HandleFunc("/test", nil)

	log.Println("Listening ..")
	http.ListenAndServe(":3000", mux)
}
