package main

import (
	"fmt"
	"time"
)

type chanEntity struct {
	A int
	B int
}

var ch = make(chan chanEntity)

func main() {
	//阻塞main 协程会死锁
	go wirte()
	go read()
	time.Sleep(time.Millisecond)
}

func wirte() {
	ch <- chanEntity{A: 0, B: 0}
	//for i := 0; 1 < 10; i++ {
	//	ch <- chanEntity{A: i, B: i}
	//}
}
func read() {
	fmt.Println(<-ch)
	//for i := 0; 1 < 10; i++ {
	//
	//	fmt.Println(<-ch)
	//}
}
