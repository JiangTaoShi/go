package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
	// 带有结构体tag的字段
	Type int `json:"type" id:"100"`
}

func main() {
	typeofCat := reflect.TypeOf(cat{})
	fmt.Println(typeofCat.Name(), typeofCat.Kind())
	for i := 0; i < typeofCat.NumField(); i++ {
		fildType := typeofCat.Field(i)
		fmt.Println(fildType.Name, fildType.Type, fildType.Tag)
	}
}
