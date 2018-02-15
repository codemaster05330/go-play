package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func basiAuthMiddlewareEx() {
	log.Println("basic auth middleware ..")

	handler := http.HandlerFunc(secureHandler)
	authHandler := httpauth.SimpleBasicAuth("testuser", "secretpassword")
	http.Handle("/", authHandler(handler))

	http.ListenAndServe(":8080", nil)
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("secure handler")
	w.Write([]byte("Kool User\n"))
}
