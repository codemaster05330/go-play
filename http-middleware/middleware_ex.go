package main

import (
	"log"
	"net/http"
)

func middlewareEx() {
	log.Println("middleware example ..")

	finalHandle := http.HandlerFunc(finalHandler)

	http.Handle("/", firstMiddleware(secondMiddleware(finalHandle)))

	http.ListenAndServe(":8080", nil)
}

func finalHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Final Handler called, request path=", req.URL.Path)
	res.Write([]byte("Final Handler"))
}

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println("In first middleware")
		next.ServeHTTP(res, req)
		log.Println("Out first middleware")
	})
}

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println("In second middleware")
		// if req.URL.Path != "/" {
		// 	return
		// }
		next.ServeHTTP(res, req)
		log.Println("Out second middleware")
	})
}
