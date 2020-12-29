package main

import (
	"fmt"
	"github.com/samy1937/req"
)

func main() {

	header := req.Header{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
	}
	r, err := req.Get("http://127.0.0.1:8081/get", header)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.ToString())
}
