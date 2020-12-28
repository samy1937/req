package main

import (
	"fmt"
	"req"
)

func main()  {
	r, err := req.Get("https://www.baidu.com/")
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.ToString())
}
