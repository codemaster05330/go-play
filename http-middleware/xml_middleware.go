package main

import (
	"bytes"
	"log"
	"net/http"
)

func xmlMiddlewareEx() {
	log.Println("xml middleware ..")

	handler := http.HandlerFunc(okHandler)
	http.Handle("/", xmlMiddleware(handler))

	http.ListenAndServe(":8080", nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ok handler")
	w.Write([]byte("I am OK\n"))
}

func xmlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("In xml middleware")

		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		if http.DetectContentType(buf.Bytes()) != "text/xml; charset=utf-8" {
			http.Error(w, http.StatusText(415), 415)
			return
		}

		next.ServeHTTP(w, r)

		log.Println("Out xml middleware")
	})
}
