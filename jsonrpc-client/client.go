package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

func main() {

	client, err := jsonrpc.Dial("tcp", "devel.behterev.su:1234")
	if err != nil {
		fmt.Printf("dialing: %s\n", err.Error())
		os.Exit(1)
	}

	// Synchronous call
	args := new(Args)
	args.A = 7
	args.B = 8
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Printf("arith error: %s", err.Error())
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
