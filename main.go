package main

import (
	"crawler/collect"
	"crawler/log"
	"crawler/proxy"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	plugin, c := log.NewFilePlugin("./log.txt", zap.InfoLevel)
	defer c.Close()
	logger := log.NewLogger(plugin)
	logger.Info("log init end")
	proxyURLs := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher failed")
	}
	url := "https://baidu.com"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}

	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read content failed:%v\\n", err)
		return
	}
	fmt.Println(string(body))

}
