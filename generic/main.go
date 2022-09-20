package main

import "fmt"

func main() {
	fmt.Println(Add[int](1, 2))
}

func Add[T int | int64](a T, b T) T {
	return a + b
}
