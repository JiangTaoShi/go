package main

import (
	"fmt"
	"time"
)

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

	// str := fmt.Sprintf(" AND order.order_id='%%%s%%'", "093840394")
	// fmt.Println(str)

	t := time.Now()
	fmt.Println(t.Weekday().String())

	// i := 101 / 10
	// fmt.Println(i)

	i := new(int)
	fmt.Println(i)

	//make(chan)

	dtime, _ := time.Parse("2006-01-02", "2006-01-03")
	fmt.Println(dtime.Format("01-02"))

}

type entity1 struct {
	name  string
	age   int
	index int
}
