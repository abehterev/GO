package main;

import (
	"net/http"
	"io"
	"fmt"
	"time"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func init() {
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterService(new(HelloService), "")
    http.Handle("/rpc", s)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var (
		out string
	)

	vars := mux.Vars(r)
	ver := vars["ver"]
	action := vars["action"]

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Server", "MuxService/0.1")

	out += "Hello, world\n"
	out += "Remote addr: " + r.RemoteAddr + "\n"
	out += "You are use: " + r.Header.Get("User-Agent") + "\n"
	out += "METHOD: " + r.Method + "\n"
	out += "rawURI: " + r.RequestURI + "\n"
	out += "URI: " + r.URL.Path + "\n"
	out += "API ver: " + ver + "\n"
	out += "ACTION: " + action + "\n"

	io.WriteString(w, out)
	fmt.Println(out)
}

type HelloArgs struct {
    Who string
}

type HelloReply struct {
    Message string
}

type HelloService struct {}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
    reply.Message = "Hello, " + args.Who + "!"
    return nil
}


func main(){

	ser := &http.Server{
		Addr:           ":4443",
		Handler:        nil,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	r := mux.NewRouter()
	//r.HandleFunc("/", homeHandl)

	r.HandleFunc("/api/{ver:[0-9]{2}}/function/{action}/", requestHandler).Host("devel.behterev.su")

	http.Handle("/", r)

	err := ser.ListenAndServeTLS("test.crt","test.key.nopass")

	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}
}
