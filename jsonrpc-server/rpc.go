package main

import (
	"fmt"
//        "net/http"
	"net"
        "net/rpc"
        "net/rpc/jsonrpc"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Printf("listen error: %s\n", e.Error())
	}
	defer l.Close()

	rpc.Register(arith)

	for {
		fmt.Printf("Waiting for connection...\n")
		if conn, err := l.Accept(); err == nil {
			fmt.Printf("Connection started: %v\n", conn.RemoteAddr())
			go jsonrpc.ServeConn(conn)
		} else {
			fmt.Printf("Failed connection acceptance: %s\n", err.Error())
		}
	}

/*	go http.Serve(l, nil)
*/
}
