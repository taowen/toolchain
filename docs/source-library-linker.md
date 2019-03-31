[[toc]]

# 要解决的问题

* 复用其他人的工作
* 不依赖二进制接口
* 完整的类型信息保留

# 解决方案

编程语言基本上都有把源代码拆分成多个source library来编写的能力。
编译的时候把多个source library定义的逻辑编译成一个或者多个object文件。
把 compiler 在这个过程里扮演的角色称之为 source library linker

## 构成

| 构成 | 解释 |
| --- | --- |
| source library | 源代码库 |
| object | 多个 source library 一步到位合并成一个 object，或者变成多个有链接关系的 object |
| source library linker | 一般就是编译器自身 |

# 解决方案案例

## tsc

### src/main.ts
<<< @/source-library-linker/tsc/src/main.ts

### src/lib.ts
<<< @/source-library-linker/tsc/src/lib.ts

### build.sh
<<< @/source-library-linker/tsc/build.sh

### dist/object.js
<<< @/source-library-linker/tsc/dist/object.js

| 构成 | 对应 |
| --- | --- |
| source library | .ts 文件 |
| static library | CJS / AMD / ES6 Module |
| object | 合并成 .js 文件 |
| source library linker | tsc |

## go

### go.mod
<<< @/source-library-linker/go/go.mod

### main.go
<<< @/source-library-linker/go/main.go

### lib/lib.go
<<< @/source-library-linker/go/lib/lib.go

### build.sh
<<< @/source-library-linker/go/build.sh

| 构成 | 对应 |
| --- | --- |
| source library | 包含 .go 文件的目录 |
| static library | 其他包含 .a 文件的目录 |
| object | 产生的二进制可执行文件 |
| source library linker | go |
