package main

import (
	"fmt"
	"log"
	"net/http"
)

// import "net/http"

func main88() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "alskd	 jflaksdj")
	fmt.Fprint(w, "alskd	 jflaksdj")
}
