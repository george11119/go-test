package main

import (
	"fmt"
	"net/http"
	"time"
)

var PORT = ":3001"

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func startHTTPServer() {
	s := http.Server{
		Addr:         PORT,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      HelloHandler{},
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func securityMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Secret-Password") != "what" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("password required\n"))
			return
		}
		h.ServeHTTP(w, r)
	})
}

func securityMiddleware2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Another-Password") != "whyanotherpassword" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("another password required ahahah\n"))
			return
		}
		h.ServeHTTP(w, r)
	})
}

func startServeMuxServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello\n"))
		},
	)

	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
	})

	wrappedMux := securityMiddleware(mux)
	wrappedMux2 := securityMiddleware2(wrappedMux)
	http.ListenAndServe(PORT, wrappedMux2)
}

func main() {
	startServeMuxServer()
}
