package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Say hi to web server!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
