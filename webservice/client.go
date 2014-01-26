package main;

import (
	"net/http"
//	"net/url"
	"crypto/tls"
//	"io"
	"io/ioutil"
//	"log"
	"fmt"
	"encoding/json"
	"strings"
)



func main(){

	tr := &http.Transport{
		DisableCompression:	true,
		TLSClientConfig:	&tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	type AC struct {
		Login string `json:"login"`
		Password string `json:"pass"`
	}

	a := &AC{
		Login: "log",
		Password: "pas",
	}

	j_a, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("JSON marshal error: %s\n",err.Error())
	}

	post := strings.NewReader(string(j_a))
	fmt.Println(post)

	req, err := http.NewRequest("POST", "https://localhost:8080/api", post)
	if err != nil {
		fmt.Printf("NewRequest error: %s\n",err.Error())
	}

	req.Header.Set("User-Agent", "TestAgent")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request error: %s\n",err.Error())
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Body error: %s\n",err.Error())
	}

	fmt.Printf("BODY:\n%s", body)
}
