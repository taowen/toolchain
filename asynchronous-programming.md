[[toc]]


# 要解决的问题


* 如何降低异步编程的难度


# 解决方案


通过更高阶的编程模型，让程序员从底层的 calback/thread 等中脱离出来，从而专注于业务逻辑。


# 解决方案案例

## ReactiveX

ReactiveX 针对各平台提供了统一的响应式编程组件，它同时支持了声明式和函数式编程。

* 基于流的抽象将时间和对象视为整体，去除了传统模型中因为时间而导致的逻辑复杂度。
* 减少了因为引入 promise/future 等概念导致的非功能性代码。

### 构成

| 构成 | 解释 |
| --- | --- |
| 事件 | 导致多时间线的触发动作，可以是同步也可是异步 |
| Observable | 基于流的抽象概念，提供事件或数据的访问|
| Subscribe | 通过订阅 Observable，进行响应 |
| Operator | 类似函数式编程，对流进行变换 |

### sample.js

<<< @/asynchronous-programming/reactive-programming-sample.js
