# 要解决的问题

	如何理解golang多值返回的实现

# 解决方案
	查看go程序对应的汇编代码，分析如何实现多参数返回
	


## 解决方案示例
####代码示例

```
package main

import (
	"fmt"
)

func main() {
	a, b := swap(3, 4)
	fmt.Println(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}
```
### 生成汇编代码命令
    go tool compile -S test.go >test.s
### swap函数汇编代码
以下代码为在mac os上编译

```
 "".swap STEXT nosplit size=21 args=0x20 locals=0x0
 73     0x0000 00000 (test.go:12)   TEXT    "".swap(SB), NOSPLIT, $0-32
 74     0x0000 00000 (test.go:12)   FUNCDATA    $0, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
 75     0x0000 00000 (test.go:12)   FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
 76     0x0000 00000 (test.go:12)   MOVQ    "".b+16(SP), AX
 77     0x0005 00005 (test.go:13)   MOVQ    AX, "".~r2+24(SP)
 78     0x000a 00010 (test.go:12)   MOVQ    "".a+8(SP), AX
 79     0x000f 00015 (test.go:13)   MOVQ    AX, "".~r3+32(SP)
 80     0x0014 00020 (test.go:13)   RET
```
解读：

	1、这个"". 代表的是这个函数的命名空间。
	2、FUNCDATA是golang编译器自带的指令，它用来给gc收集进行提示。提示0和1是用于局部函数调用参数，需要进行回收。
	3、SP：伪寄存器，指向栈顶

swap(SB) 这里就有个SB的伪寄存器。全名未Static Base，代表swap这个函数地址，0−32中的0代表局部变量字节数总和，0表示不存在局部变量。-32代表在0的地址基础上空出32的长度作为传入和返回对象。这个也就是golang如何实现函数的多返回值的方法了。它在定义函数的时候，开辟了一定空间存储传入和传出对象。

76、77行，将b赋值到SP+24地址
78、79行，将a赋值到SP+32地址


###未完待续
plan9 函数调用过程


#参考：

[golang汇编](https://lrita.github.io/2017/12/12/golang-asm/)

[golang汇编解读](https://www.cnblogs.com/yjf512/p/6132868.html)
	








