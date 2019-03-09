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
| executor | node.js 或者浏览器 |

## rustc

rust 提供了更灵活和安全的类型系统来辅助描述驱动机器的逻辑

### hello.rs
<<< @/compiler/rustc/hello.rs

### build.sh
<<< @/compiler/rustc/build.sh

编译出来的是二进制的可执行文件，运行在CPU上

| 构成 | 对应 |
| --- | --- |
| source | hello.rs |
| executable | ./hello |
| compiler | rustc |
| executor | CPU |

## vue-loader

webpack/vue-loader/vue-template-compiler 组合成了一个完整的编译器。
它可以把一个 .vue 单文件组件编译成 javascript 写成的 render 函数。
.vue 单文件组件，可以使用 vue 的模板语法，比 javascript 渲染 dom 的写法更可读。

### hello.vue
<<< @/compiler/vue-loader/src/hello.vue

### webpack.config.js
<<< @/compiler/vue-loader/webpack.config.js

### build.sh
<<< @/compiler/vue-loader/build.sh

编译出来的 hello.js 文件是这个样子的

### hello.js
<<< @/compiler/vue-loader/dist/hello.js

其中 eval 的关键源代码，就是把 vue 模板编译出来的 javascript 代码

```js
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "render", function() { return render; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "staticRenderFns", function() { return staticRenderFns; });
var render = function() {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c("div", [
    _c(
      "ul",
      _vm._l(_vm.items, function(item) {
        return _c("li", [_vm._v(_vm._s(item))])
      }),
      0
    )
  ])
}
var staticRenderFns = []
render._withStripped = true
```

使用这个编译好的 hello.js 文件

### index.html
<<< @/compiler/vue-loader/index.html
