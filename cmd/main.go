package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("running")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
