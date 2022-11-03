package main

import "fmt"

func main() {
	var list []entity1
	list = append(list, entity1{
		name:  "1",
		age:   10,
		index: 0,
	})
	list = append(list, entity1{
		name:  "2",
		age:   20,
		index: 5,
	})

	list = append(list, entity1{
		name:  "2",
		age:   20,
		index: 4,
	})

	fmt.Println(cap(list))
	fmt.Println(list[:0])

	// sort.Slice(list, func(ii, jj int) bool {
	// 	return list[ii].index < list[jj].index
	// })

	// for _, val := range list {
	// 	fmt.Println(val.index)
	// }

	// var list []string
	// list = append(list, "1")
	// list = append(list, "1")
	// list = append(list, "1")
	// fmt.Println(cap(list))
	// fmt.Println(len(list))

}

type entity1 struct {
	name  string
	age   int
	index int
}
