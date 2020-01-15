package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, I can't believe that worked!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Go Simple Server")
}

func main() {

	const PORT = 5000
	var builder strings.Builder
	builder.WriteString(":") // prefix
	builder.WriteString(strconv.Itoa(PORT))
	var PORT_STRING = builder.String()

	fmt.Println("Server running on port", PORT, "...")
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(PORT_STRING, nil)

}
