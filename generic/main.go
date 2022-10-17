package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Add[int](1, 2))

	//time, _ := time.Parse("2006-01-02", "2016-01-02")
	//month := int(time.Month())
	var a float32 = 1.5

	fmt.Println(fmt.Sprintf("%v", a))
}

func Add[T int | int64](a T, b T) T {
	return a + b
}
