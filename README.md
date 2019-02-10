# 编程工具链

编程使用的工具链正变得越来越复杂。你过去构建的知识体系所基于的技术栈，今天可能就会被抽象掉。并不是我们过去懂的东西现在不存在了，而是做为一块基石，被更高阶的抽象给代替了。无论这些工具的名字怎么五花八门，他们所尝试解决的问题都是相似的。从解决的问题入手，可以帮助我理解和记得这个庞杂的工具体系。

这里会讨论的工具链的范围如下：

* Javascript in browser / node
* Java
* Go / Rust
* Linux / Mac OS

工具链解决的问题

| 问题 | 解决方案代号 |
| --- | --- |
| [如何驱动机器按照人的要求自动化执行重复的工作](executor.md) | executor |
| [如何复用其他人的工作，又不共享源代码](dynamic-library-linker.md) | dynamic library linker（executor 的一部分） |
| [如何标识并定位动态链接库](/dynamic-library-resolver.md) | dynamic library resolver（linker 的一部分） |
| [如何引用动态链接库指定的接口](symbol-binder.md) | symbol binder（linker 的一部分） |

解决方案案例表

| source file | source pkg | compiler |
| --- | --- | --- |
| [.ts](/dot-ts.md) | vue sfc | tsc |

