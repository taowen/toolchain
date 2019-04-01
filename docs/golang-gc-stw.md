[TOC]

# 要解决的问题

* 如何理解GO中特殊情况下GC引起的卡死

### 问题背景

在GO中，goroutine经常被使用以满足并发的需求，但在某些特殊情况下可能会有问题

```$xslt
func infiniteLoop() {
	fmt.Println("infiniteLoop goroutine started")
	for {
	}
}

func main() {
	fmt.Println("main started")
	go infiniteLoop()
	fmt.Println("prepare to gc")
	runtime.Gosched() // 让出goroutine执行权
	runtime.GC()      // 手动gc
	fmt.Println("main finished")
}
```

<<< @/golang-gc-stw/gc-stw.go

程序会在打印“prepare to gc”后卡住，无法执行至“main finished”。而如果注释掉runtime.Gosched()和runtime.GC()，程序则可顺利打印“main finished”，并结束整个进程。

# 解决方案

### goroutine调度方式

GO中目前goroutine的调度方式是“合作式抢占”（cooperative preemption）。对于一个正在执行的goroutine，只有当此goroutine中出现fucntion call的时候，才会让出执行权。这里的fucntion call包括普通的函数调用，系统函数调用，channel阻塞等。

从调度模型而言，GO使用的是**GPM模型**：

G（goroutine）：GO中的协程。

P（processor）：逻辑执行单元，一次跑一个goroutine，默认为核心数。每个P执行的goroutine以队列的形式存在。

M（machine）：对内核级线程的封装，CPU数。

特别的，如果跑goroutine的执行单元P只有一个，那当一个无限循环的goroutine执行之后，会有程序卡死的情况。但当P的数量大于1时，基于Work-Stealing策略，空闲的执行单元P会“偷取”其他P中的goroutine，所以P>1时，不会出现被无限循环卡死的情况。

```$xslt

func main() {
	runtime.GOMAXPROCS(1) // 设置P的数量为1
	fmt.Println("main started")
	go infiniteLoop()
	fmt.Println("prepare to gc")
	runtime.Gosched() // 让出goroutine执行权
	fmt.Println("main finished")
}
```
设置P的数量为1后，此时不触发GC也会卡死，因为goroutine执行权已被infiniteLoop获取。

### GC STW

以上的调度方式实现较为简单，在绝大多数场景下是没有问题的。但当触发GC的时候，仍会有问题。

GO中GC目前使用的是**三色标记清除**法，在标记阶段仍会有短暂的STW阶段。由于“合作式抢占”的策略，STW会被动地等待所有goroutine运行到可以被调度的状态，然后暂停此goroutine。在程序执行至runtime.GC()时，main goroutine由于本身的函数调用（runtime.GC()），因而暂停等待STW。但infiniteLoop goroutine中没有任何函数调用，于是导致GC进程一直等待infiniteLoop goroutine。最终的结果是main goroutine等待另一个无限循环的goroutine，导致程序卡死。

### 解决方案

上述情况出现的概率极小。但在官方未解决此问题之前，避免写“无函数调用的无限循环“即可

```$xslt

func infiniteLoopCall() {
	fmt.Println("infiniteLoopCall")
}

func infiniteLoop() {
	fmt.Println("infiniteLoop goroutine started")
	for {
		infiniteLoopCall() // 加入此函数调用可避免卡死
	}
}
```

在18年，已有一个“Non-cooperative goroutine preemption”的方案被提出（2019-01-18有更新），通过信号的方式强制抢占goroutine的执行权。

### 参考资料

[Golang并发原理及GPM调度策略（一）](https://www.cnblogs.com/mokafamily/p/9975980.html)

[Golang 里一个有趣的小细节](https://zhuanlan.zhihu.com/p/44851211)

[Proposal: Non-cooperative goroutine preemption](https://github.com/golang/proposal/blob/master/design/24543-non-cooperative-preemption.md)