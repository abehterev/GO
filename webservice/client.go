package main;

import (
	"net/http"
	"net/url"
	"crypto/tls"
	"fmt"
)

func main(){

	tr := &http.Transport{
	//	TLSClientConfig:	&tls.Config{RootCAs: pool},
		DisableCompression:	true,
		TLSClientConfig:	&tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
	//	CheckRedirect: redirectPolicyFunc,
		Transport: tr,
	}
	
	resp, err := client.PostForm("https://localhost:8080/", url.Values{"key": {"Value"}, "id": {"123"},})

	//resp, err := client.Get("http://localhost:8080/robots.txt")

	resp.Body.Close()	

	if err != nil {
		fmt.Println("Err: " + err.Error() + "\n")
	}

}
