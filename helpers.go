package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func MakeRequest(url string, method string, headers []string) (int, string) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if headers[0] != "" {
		for _, header := range headers {
			splitted := strings.Split(header, ":")
			req.Header.Add(splitted[0], splitted[1])
		}
	}

	if err != nil {
		log.Fatalln(err)
	}
	resp, err2 := client.Do(req)

	if err2 != nil {
		log.Fatal("Error reading response. ", err)

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	defer resp.Body.Close()
	return resp.StatusCode, string(body)
}