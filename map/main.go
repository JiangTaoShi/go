package main

import (
	"fmt"
	"time"
)

//多个协程写入会报错
var countryMap = make(map[string]string)

func main() {
	go fmtvalue()
	go fmtvalue()
	go fmtvalue()
	go fmtvalue()

	time.Sleep(2)
}

func fmtvalue() {
	v, ok := countryMap["test"]
	if ok {
		fmt.Println(v)
	} else {
		countryMap["test"] = "测试"
	}
}
