package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/wonderivan/logger"
)

func FuncA(nums []int) {
	if nums == nil {
		panic("nums == nil")
	}
	fmt.Println(nums[1])
}

func Process() (err error) {
	logger.Info("Enter function Process")

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			// log
			logger.Warn("There is a panic in Process:", panicInfo)
			panicInfoMsg := fmt.Sprintf("%s", panicInfo)
			if strings.Contains(panicInfoMsg, "nil") {
				// return err
				err = errors.New("FucA panic: nums == nil")
			}
		}
	}()

	// 引发panic
	FuncA(nil)

	logger.Info("Exit function Process")

	return
}

func main() {
	err := Process()
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Printf("err is not nil: %v\n", err)
	}
}
