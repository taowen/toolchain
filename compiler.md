[[toc]]

# 要解决的问题

如何用更高阶的编程语言来驱动机器减少工作量

# 解决方案

用高阶编程语言编写代码，然后转换成机器直接支持的代码，从而允许利用机器不直接支持的抽象模式来减少工作量

## 构成

| 构成 | 解释 |
| --- | --- |
| source | 源代码 |
| object | 目标文件，需要静态链接才能变成executable |
| executable | 可执行文件 |
| compiler | 编译器，把 source 编译成 object，或者一步到位变成 executable |

# 解决方案案例

## tsc

typescript 提供了两类编程上的便利

* 给 javascript 增加了类型
* 把高版本的 javascript 编译为低版本的 javascript 以兼容更多的 executor

### source.ts
<<< @/compiler/tsc/source.ts

### build.sh
<<< @/compiler/tsc/build.sh

### executable.js
<<< @/compiler/tsc/executable.js

可以看到类型信息没有了， for of 的语法被编译成了等价的普通的 for 循环

| 构成 | 对应 |
| --- | --- |
| source | source.ts |
| executable | executable.js |
| compiler | tsc |



