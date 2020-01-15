package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, I can't believe that worked!")
}

func main() {
	fmt.Println("Server running on port 8000...")
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
