package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func MakeRequest(url string, method string, headers []string, parameters string, timeout int) (int, string) {

	if url != "" {
		url = url + "?" + parameters
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

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
