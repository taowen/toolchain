# 要解决的问题

如何用更高阶的编程语言来驱动机器减少工作量

# 解决方案

用高阶编程语言编写代码，然后转换成机器直接支持的代码，从而允许利用机器不直接支持的抽象模式来减少工作量

## 构成

* source 源代码
* compiler 编译器
* executable

# 解决方案案例

## typescript

* source：.ts文件
* executable：.js文件
* compiler：tsc

typescript 提供了两类编程上的便利

* 给 javascript 增加了类型
* 把高版本的 javascript 编译为低版本的 javascript 以兼容更多的 executor

<<< @/compiler/tsc/source.ts

```typescript
// /opt/source.ts
let messages: string[] = ['hello', 'world'] 
for (let elem of messages) {
  console.log(elem)
}
```

```
// --target ES3 把高版本的 emcascript 编译为最初版的 javascript
tsc --target ES3 --outFile executable.js source.ts
```

```js
// /opt/executable.js
var messages = ['hello', 'world']; 
for (var _i = 0, messages_1 = messages; _i < messages_1.length; _i++) {
    var elem = messages_1[_i];
    console.log(elem);
}
```

可以看到类型信息没有了， for of 的语法被编译成了等价的普通的 for 循环

