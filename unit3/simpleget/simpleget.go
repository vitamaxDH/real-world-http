package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
	log.Println("StatusCode", resp.StatusCode)
	log.Println("Proto", resp.Proto)
	log.Println("Proto", resp.Header)
	log.Println("Content-Length", resp.Header.Get("Content-Length"))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
