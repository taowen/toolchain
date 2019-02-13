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




