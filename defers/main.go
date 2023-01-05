package main

import (
	"fmt"
	"strconv"
)

func main() {
	//exec()
	//time.Sleep(time.Second * 30)
	var id int64
	id = 70802
	idstr := strconv.FormatInt(id, 10)
	text := fmt.Sprintf("%s%s", "tdx", idstr)
	fmt.Println(text)
}

func exec() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}
