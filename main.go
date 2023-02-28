package main

import (
	"crawler/collect"
	"fmt"
)

func main() {
	url := "https://book.douban.com/subject/1007305/"
	var f collect.Fetcher = collect.BrowserFetch{}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	fmt.Println(string(body))
}
