package main;

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"time"
	"encoding/json"
)

func requestAPI(w http.ResponseWriter, r *http.Request) {
	var (
		out string
	)

	if (r.Method != "POST"){
		out = "Err: Method not supported.\n"
	}else{
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Server", "CollogueService/0.1")

		out += "Hello! If you know how, you can use API for connectivity.\n"
		out += "Remote addr: " + r.RemoteAddr + "\n"
		out += "You are use: " + r.Header.Get("User-Agent") + "\n"
		out += "METHOD: " + r.Method + "\n"
		out += "rawURI: " + r.RequestURI + "\n"
		out += "URI: " + r.URL.Path + "\n"
	}
	io.WriteString(w, out)

	defer r.Body.Close()
        body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Body error: %s\n",err.Error())
	}

	type AC struct {
		Login string `json:"login"`
		Password string `json:"pass"`
	}

	var ac AC

	err = json.Unmarshal(body, &ac)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %s\n",err.Error())
	}

	fmt.Println(ac)
}

func main(){

	ser := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	http.HandleFunc("/api", requestAPI)
	err := ser.ListenAndServeTLS("test.crt","test.key.nopass")

	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}

}
