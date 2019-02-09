# 要解决的问题

* 驱动机器

# 解决方案

机器是一个抽象的概念，提供可被驱动机制的东西。不仅仅我们常用的x86芯片是机器，Java虚拟机也是机器。

## 构成

* executor：机器本身，提供 builtin api/conditio/loop
* executable：承载算法的介质，本身的格式说明
* bootstrap：让executor加载executable的过程
* builtin api：机器自身提供的编程接口，executable是对builtin api的组装
  * input/output api：和executor外部提供输入输出的交互，是builtin api的一个子集
  * condition api：表达分支逻辑
  * loop api：表达循环逻辑

## 衍生的问题

* [如何复用其他人的工作，又不共享源代码](/dynamic-library-linker.md)

# 解决方案案例

## 浏览器中的 JavaScript

* executor：浏览器的JavaScript引擎，例如 Chrome 的 V8引擎
* executable：直接嵌入到 html 的 `<script>` 标签的 JavaScript 代码
* bootstrap：url获取html文件，html文件引用js文件的url，加载js文件执行

```html
// http://localhost/index.html
<html>
<head>
<script>
for (let i = 0; i < 3; i++) {
  console.log('i am the executable')
}
</script>
</head>
<body>
</body>
</html>
```

用浏览器执行，访问 `http://localhost/index.html` 在浏览器的控制台输出

```
i am the executable
i am the executable
i am the executable
```

## 服务端的 JavaScript

* executor：nodejs
* executable：.js文件
* bootstrap：命令行执行 node xxx.js

```js
// executable.js
for (let i = 0; i < 3; i++) {
  console.log('i am the executable')
}
```

用 nodejs 执行 executable.js

```
node executable.js
// Output:
// i am the executable
// i am the executable
// i am the executable
```