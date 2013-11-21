package main;

import (
	"net/http"
	"io"
	"fmt"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var (
		out string
	)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Server", "WebService/0.1")

	out += "Hello, world\n"
	out += "Remote addr: " + r.RemoteAddr + "\n"
	out += "You are use: " + r.Header.Get("User-Agent") + "\n"
	out += "METHOD: " + r.Method + "\n"
	out += "rawURI: " + r.RequestURI + "\n"
	out += "URI: " + r.URL.Path + "\n"

	io.WriteString(w, out)
	fmt.Println(out)
}

func main(){

	ser := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,	
	}

	http.HandleFunc("/", requestHandler)
	err := ser.ListenAndServeTLS("test.crt","test.key.nopass")

//	http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}
}
