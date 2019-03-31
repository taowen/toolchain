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

type ErrorProcess struct {
	Error      error
	ProcessNum int
}

func Process(processNum int, errChan chan ErrorProcess) (err error) {
	logger.Info("Enter function Process")

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			// log
			logger.Warn("There is a panic in Process:", panicInfo)
			panicInfoMsg := fmt.Sprintf("%s", panicInfo)
			if strings.Contains(panicInfoMsg, "nil") {
				// return err
				err = errors.New("FucA panic: nums == nil")
				errChan <- ErrorProcess{err, processNum}
			}
		}
	}()

	// 引发panic
	FuncA(nil)

	logger.Info("Exit function Process")

	return
}

func main() {
	errChan := make(chan ErrorProcess)
	for i := 0; i < 10; i++ {
		go Process(i, errChan)
	}

	for {
		errProcess, ok := <-errChan
		if ok {
			fmt.Printf("err is not nil: %v\n", errProcess)
		} else {
			fmt.Printf("All process is finished")
			break
		}
	}

}
