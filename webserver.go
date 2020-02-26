package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server up and running at :8080. Visit http://localhost:8080/yourname to test")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// greet user echoing the text after the trailing slash eg: http://localhost:8080/username will say Hello, username!
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
