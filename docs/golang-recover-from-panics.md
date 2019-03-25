[[toc]]

# 要解决的问题

* 如何实现golang程序panic时的自救

# 解决方案

* 对Go语言panic相关原理进行剖析，以尽可能优雅地在程序panic发生时进行自救。

# Go语言panic相关原理
## panic发生时程序控制权如何变更

* golang使用调用栈维护函数间调用关系
* 每当发生函数调用，即封装当前函数运行信息并压入调用栈，并转移程序控制权给被调用函数
* 每当代码产生panic，即从调用栈弹出其上一级调用函数，并转移程序控制权给调用函数
* 程序控制权沿调用栈的反方向，逐级传播至最外层调用函数——go函数或main函数

## Go语言内建的recover()函数
* 调用时如果程序未处于panic，则返回nil，对程序控制权无影响
* 调用时如果程序处于panic，则返回panic信息，并重新获得程序控制权
* 使用方式：与defer语句联用

## defer语句
* 作用：延迟与defer语句联用的函数代码的执行
* 延迟时机：defer语句所属函数即将结束执行的那一刻，无论是否产生panic
* 前置条件：在函数产生panic之前，defer语句已经执行

# 解决方案案例
## 单goroutine场景
* recover()恢复 + 日志打印
```
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
}
```

源代码路径：../../golang-recover-from-panics/recover&log/demo.go

* recover() + log + 异常封装成错误
```
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
```

源代码路径：../../golang-recover-from-panics/recover&log&error/demo.go

## 多goroutine场景(异常监听)
* recover() + log + 异常封装成错误 + channel

```
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
```
源代码路径：../../golang-recover-from-panics/recover&log&channel/demo.go
