package main

import (
	"fmt"
	"runtime"
)

func infiniteLoopCall() {
	fmt.Println("infiniteLoopCall")
}

func infiniteLoop() {
	fmt.Println("infiniteLoop goroutine started")
	for {
		// infiniteLoopCall()
	}
}

func main() {
	// runtime.GOMAXPROCS(1)
	fmt.Println("main started")

	go infiniteLoop()

	fmt.Println("prepare to gc")

	runtime.Gosched() // 让出goroutine执行权
	// runtime.GC()      // 手动gc

	fmt.Println("main finished")
}
