package main

import (
	"crawler/collect"
	doubangroup "crawler/parse/douban"
	"fmt"
)

func main() {
	var workList []*collect.Request
	for i := 25; i <= 100; i += 25 {
		str := fmt.Sprintf("<https://www.douban.com/group/szsh/discussion?start=%d>", i)
		workList = append(workList, &collect.Request{
			Url:       str,
			ParseFunc: doubangroup.ParseURL,
		})
	}
}
