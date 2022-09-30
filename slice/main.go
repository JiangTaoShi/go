package main

import "fmt"

func main() {
	var list []entity1
	list = append(list, entity1{
		name: "1",
		age:  10,
	})
	list = append(list, entity1{
		name: "2",
		age:  20,
	})

	fmt.Println(len(list))
	for i, _ := range list {
		fmt.Println(i)
	}

}

type entity1 struct {
	name string
	age  int
}
