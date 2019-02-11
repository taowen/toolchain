# 要解决的问题

* 复用其他人的工作
* 不需要运行时支持动态链接库
* 不共享源代码的前提下复用

# 解决方案

静态链接库 linker 在构建的时期，把 object 文件和 static library 预先合并成一个完整的 executable。object 与 static library 彼此知道自己的链接关系。

