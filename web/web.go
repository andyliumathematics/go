package main

import (
	"net/http"
)

func main1() {
	// http.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("hello world"))
	// })
	// http.ListenAndServe("localhost:8080",nil)
	serv()

}
func serv() {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	server.ListenAndServe()

}

type myHandler struct{}
type helloHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HellosssssWorld"))
}
func main() {
	mh := myHandler{}
	hh :=helloHandler{}
	http.Handle("/hello", &mh)
	http.Handle("/hellome", &hh)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	server.ListenAndServe()

}
