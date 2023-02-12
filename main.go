package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn"
	resp,err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url error:%v", err)
		return
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}
	fmt.Println("body:", string(body))
}
