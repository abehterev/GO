package main;

import (
	"net/http"
	"fmt"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func main(){
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
