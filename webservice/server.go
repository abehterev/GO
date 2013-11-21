package main;

import (
	"net/http"
	"io"
	"fmt"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var (
		out string
	)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Server", "WebService/0.1")
	out += "Hello, world\n"
	out += "You are use: " + r.Header.Get("User-Agent") + "\n"
	out += "METHOD: " + r.Method + "\n"
	out += "rawURI: " + r.RequestURI + "\n"
	out += "URI: " + r.URL.Path + "\n"
	io.WriteString(w, out)
	fmt.Println(out)
}

func main(){
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServeTLS(":8080", "test.crt", "test.key.nopass", nil)
//	http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}
}
