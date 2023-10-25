package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	for {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		base_url := "https://wso2.keys4.local:8243/personsfastapi/0.1.0/api/persons/"
		id := rand.Intn(50)
		url := ""
		if id == 0 {
			url = base_url
		} else {
			url = base_url + strconv.Itoa(id)
		}
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		fmt.Printf("%d - %s - sleeping %d seconds...\n", resp.StatusCode, url, n)
		time.Sleep(time.Duration(n) * time.Second)
	}
}
