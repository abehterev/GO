package main

import (
	"io"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

type FCGIServer struct{}

func (s FCGIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        var (
                out string
        )   

        out += "Hello, world\n"
        out += "Remote addr: " + r.RemoteAddr + "\n"
        out += "You are use: " + r.Header.Get("User-Agent") + "\n"
        out += "METHOD: " + r.Method + "\n"
        out += "rawURI: " + r.RequestURI + "\n"
        out += "URI: " + r.URL.Path + "\n"

        io.WriteString(w, out)
        fmt.Println(out)
}

func main() {
	fcgi_listener, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}
	defer fcgi_listener.Close()
	fcgi_srv := new(FCGIServer)
	fcgi.Serve(fcgi_listener, fcgi_srv)
}

