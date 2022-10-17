package main

import (
	"fmt"
)

//多个协程写入会报错
var countryMap = make(map[string]interface{})

func main() {
	//go fmtvalue()
	//go fmtvalue()
	//go fmtvalue()
	//go fmtvalue()

	//time.Sleep(2)

	countryMap["val1"] = "2938293"

	var1 := countryMap["val1"]
	fmt.Println(fmt.Sprintf("%v", var1))

}

func fmtvalue() {
	v, ok := countryMap["test"]
	if ok {
		fmt.Println(v)
	} else {
		countryMap["test"] = "测试"
	}
}
