package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("GO server running!")
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}

func servePage(writer http.ResponseWriter, reqest *http.Request) {
	io.WriteString(writer, "Hello world!")
}
