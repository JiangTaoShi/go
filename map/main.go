package main

import (
	"fmt"
	"sync"
)

//多个协程写入会报错
var countryMap = make(map[string]interface{})

func main() {
	//go fmtvalue()
	//go fmtvalue()
	//go fmtvalue()
	//go fmtvalue()

	//time.Sleep(2)

	// countryMap["val1"] = "2938293"

	// var1 := countryMap["val1"]
	// fmt.Println(fmt.Sprintf("%v", var1))

	var syncMap = sync.Map{}
	syncMap.Store(1, "1")
	syncMap.Store(1, "2")

	val, ok := syncMap.Load(1)
	if ok {
		fmt.Println(val)
	}

}

func fmtvalue() {
	v, ok := countryMap["test"]
	if ok {
		fmt.Println(v)
	} else {
		countryMap["test"] = "测试"
	}
}
