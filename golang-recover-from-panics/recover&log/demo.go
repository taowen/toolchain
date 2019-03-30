package main

import (
	"fmt"

	"github.com/wonderivan/logger"
)

func genPanic() {
	logger.Info("Enter function genPanic")

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			logger.Warn("There is a panic in genPanic:", panicInfo)
		}
	}()

	// 引发panic
	nums := make([]int, 0)
	fmt.Println(nums[1])

	logger.Info("Exit function genPanic")
}

func main() {
	genPanic()
}
