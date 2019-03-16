# 编程工具链

编程使用的工具链正变得越来越复杂。你过去构建的知识体系所基于的技术栈，今天可能就会被抽象掉。
并不是我们过去懂的东西现在不存在了，而是做为一块基石，被更高阶的抽象给代替了。
无论这些工具的名字怎么五花八门，他们所尝试解决的问题都是相似的。
以他们所解决的问题来主轴，可以建立一个网络形式的知识索引，方便记忆。
同时把解决相似问题的解决方案罗列在一起，可以达到触类旁通的效果。

解决的问题

| 问题 | 解决方案代号 |
| --- | --- |
| [如何用更高阶的编程语言来驱动机器减少工作量](compiler.md) | compiler |
| [如果复用其他人的工作又不依赖二进制接口](source-library-linker.md) | source library linker |
| [如果复用其他人的工作又不需要运行时去额外加载](static-library-linker.md) | static library linker |
| [如何驱动机器按照人的要求自动化执行重复的工作](executor.md) | executor |
| [如何复用其他人的工作又减少executable的文件尺寸](dynamic-library-linker.md) | dynamic library linker（一般是 executor 的一部分） |
| [如何标识并定位动态链接库](dynamic-library-resolver.md) | dynamic library resolver（linker 的一部分） |
| [如何引用动态链接库指定的接口](symbol-binder.md) | symbol binder（linker 的一部分） 
| [如何理解高级语言的垃圾回收](garbage-collection.md) | garbage collection |
| [如何减少内存碎片](reduce-memory-fragmentation.md) | reduce memory fragmentation  |
| [如何降低异步编程的难度](asynchronous-programming.md) | asynchronous programming |
| [如何深入准确地理解一门高阶语言的特性](functional-programming.md) | functional programming |
| [如何理解GO GC原理和优化思路](golang-garbage-collection.md)|golang garbage collection|
| [如何动态修改已有代码](code-enhancer.md)|code enhancer |
| [如何降低外部服务使用的复杂度](service-management.md)|service-management |
# 本地预览
* 安装 https://yarnpkg.com/
* 安装 https://github.com/tj/n

```
yarn install
yarn dev
```

访问 http://localhost:8080/

# 作者邀请表

#作者邀请表
* [taowen](https://github.com/taowen)
* [qinguanri](https://github.com/qinguanri)
  * zhangchengshuai
  * nizeyang
* zhangkun
* zhanghuanming
* yuhongyu
  * [chenyi](https://github.com/mrgeneralgoo)
* dongqingming
* [liqingsong](https://github.com/matoujun)
