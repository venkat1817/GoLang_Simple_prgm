package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the home page")
}
func handleRequest() {

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequest()
}
