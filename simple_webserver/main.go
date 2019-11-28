package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer : go run main.go -> http://localhost:8080/Phi
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s!", r.URL.Path[1:])
}
