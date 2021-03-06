package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, I can't believe that worked!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Go Simple Server")
}

func main() {

	var buffer bytes.Buffer

	const PORT = 8000

	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(PORT))

	var PORT_STRING = buffer.String()

	fmt.Println("Server running on port", PORT, "...")
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(PORT_STRING, nil)

}
