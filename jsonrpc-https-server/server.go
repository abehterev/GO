package main;

import (
	"net/http"
	"fmt"
	"time"
	"errors"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func init() {
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterService(new(HelloService), "")
    http.Handle("/rpc", s)
}

var ErrResponseError = errors.New("response error")

type HelloService struct {}

type SumArgs struct {
	a int
	b int
}

type SumReply struct {
	rep int
}

func (h *HelloService) Sum(r *http.Request, args *SumArgs, reply *SumReply) error {
	reply.rep = args.a + args.b
	fmt.Printf("%d + %d = %d\n",args.a,args.b,reply.rep)
	return nil
}

func (h *HelloService) ResponseError(r *http.Request, args *SumArgs, reply *SumReply) error {
        return ErrResponseError
}

func main(){

	ser := &http.Server{
		Addr:           ":4443",
		Handler:        nil,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

//	err := ser.ListenAndServeTLS("test.crt","test.key.nopass")
	err := ser.ListenAndServe()

	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}
}
