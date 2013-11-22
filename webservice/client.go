package main;

import (
	"net/http"
	"net/url"
	"crypto/tls"
	"io/ioutil"
	"log"
	"fmt"
)

func failOnError(err error, msg string) {
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
                panic(fmt.Sprintf("%s: %s", msg, err))
        }
}

func main(){

	tr := &http.Transport{
		DisableCompression:	true,
		TLSClientConfig:	&tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}
	
	resp, err := client.PostForm("https://localhost:8080/", url.Values{"key": {"Value"}, "id": {"123"},})
	failOnError(err, "Failed to connect")
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	failOnError(err, "Failed to get body")

	fmt.Sprintf("BODY:\n%s", body)

}
