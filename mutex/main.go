package main

import (
	"sync"
)

var mutex sync.Mutex

func main() {
	// count := 0
	// for i := 0; i < 5000; i++ {
	// 	go func() {
	// 		mutex.Lock()
	// 		count += 1
	// 		mutex.Unlock()
	// 	}()
	// }
	// time.Sleep(time.Millisecond)
	// fmt.Println(count)
	//fmt.Println(os.Getwd())

}
