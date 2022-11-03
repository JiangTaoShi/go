package main

import (
	"fmt"
	"strings"
)

func main() {
	// u, err := url.ParseRequestURI("https://tunshu.oss-cn-zhangjiakou.aliyuncs.com/tunshu/Picture/431973573217747129.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// u.Host = "tunshu.oss-cn-zhangjiakou-internal.aliyuncs.com"
	// fmt.Println(u.String())

	trHaiCoder := "服务经理：{$}"
	test := "张三"
	test = strings.Replace(trHaiCoder, "{$}", test, 1)
	fmt.Println("StrReplace =", test)

}
