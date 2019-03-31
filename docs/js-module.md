# 要解决的问题
  js的原始的开发方式，无法满足复杂度高、代码量多的项目。 无法解决js中命名全局污染冲突、js加载顺序、多文件的网络加载次数等问题
# 解决方案
 理想情况下，开发者只需要实现核心的业务逻辑，无需关心命名冲突问题、使用什么模块引入即可。今天探讨Javascript模块化编程

# js Module 详解：
## commonJS
```js
  // moduleA.js
  var name = 'weiqinl'
  function foo() {}

  module.exports = exports = {
    name,
    foo
  }

  // moduleB.js
  var ma = require('./moduleA')
  exports.bar = function() {
    ma.name === 'weiqinl'
    ma.foo()
  }

  // moduleC.js
  var mb = require('./moduleB')
  mb.bar()

```
* 特点
  1. 适用于服务端，
* 优点：
  1. 非常简单，很容易使用

* 缺点：
  1. 阻塞式的调用不适用于网络
  2. 不能并发请求多个模块

## AMD(Require.js)
```js
  // moduleA.js
  // 定义模块 define(id?, dependencies?, factory);
  define(['jQuery','lodash'], function($, _) {
    var name = 'weiqinl',
    function foo() {}
    return { // 输出API
      name,
      foo
    }
  })

  // index.js
  // 加载模块 require([module]?, callback); 加载模块
  require(['moduleA'], function(a) {
    a.name === 'weiqinl'
    a.foo()
  })

  // index.html
  <script src="js/require.js" data-main="js/index"></script>
```
* 特点：
  1. 推崇依赖前置，提前加载，提前执行
* 优点：
  1. 实现了异步请求方式；
  2. 并行加载多个模块

* 缺点：
  1. 代码读写起来比较吃力
## CMD (Sea.js)
```js
  // moduleA.js
  // 定义模块
  define(function(require, exports, module) {
    var func = function() {
      var a = require('./a') // 到此才会加载a模块
      a.func()
      if(false) {
        var b = require('./b') // 到此才会加载b模块
        b.func()
      }
    }
    exports.func = func;
  })

  // index.js
  // 加载使用模块
  seajs.use('moduleA.js', function(ma) {
    var ma = math.func()
  })

  // HTML，需要在页面中引入sea.js文件。
  <script src="./js/sea.js"></script>
```
* 特点：
  1. 推崇依赖就近，延迟执行
* 优点：
  1. 实现了commonJS的exports、require语法，
  2. 按需加载对是web页面性能提高，也贡献不少

## UMD
```js
  // 使用Node, AMD 或 browser globals 模式创建模块
  (function (root, factory) {
    if (typeof define === 'function' && define.amd) {
      // AMD模式. 注册为一个匿名函数
      define(['b'], factory);
    } else if (typeof module === 'object' && module.exports) {
      // Node等类CommonJS的环境
      module.exports = factory(require('b'));
    } else {
      // 浏览器全局变量 (root is window)
      root.returnExports = factory(root.b);
      }
    }(typeof self !== 'undefined' ? self : this, function (b) {
      return {};
      }
    )
  );
```
* 特点
  1. 判断define为函数，并且是否存在define.amd，来判断是否为AMD规范,
  2. 判断module是否为一个对象，并且是否存在module.exports来判断是否为CommonJS规范
  3. 如果以上两种都没有，设定为原始的代码规范
* 优点
  1. 该模式主要用来解决CommonJS模式和AMD模式代码不能通用的问题，并同时还支持老式的全局变量规范

## ES6
Export
```js
  // 输出变量
  export var name = 'weiqinl'
  export var year = '2018'

  // 输出一个对象（推荐）
  var name = 'weiqinl'
  var year = '2018'
  export { name, year}


  // 输出函数或类
  export function add(a, b) {
    return a + b;
  }

  // export default 命令
  export default function() {
    console.log('foo')
  }
```
Import
```js 
  // 正常命令
  import { name, year } from './module.js' //后缀.js不能省略

  // 如果遇到export default命令导出的模块
  import ed from './export-default.js'

```
* 特点
  1. 它因为是标准，所以未来很多浏览器会支持，可以很方便的在浏览器中使用。
  2. 它同时兼容在node环境下运行。
  3. 模块的导入导出，通过import和export来确定。
  4. 可以和Commonjs模块混合使用。
  5. CommonJS输出的是一个值的拷贝。ES6模块输出的是值的引用,加载的时候会做静态优化。
  6. CommonJS模块是运行时加载确定输出接口，ES6模块是编译时确定输出接口。

# 总结
以上是js模块的发展历程，从最原始js-> commonjs -> amd -> cmd - es6，中间每一个里程碑的发展，都解决了用户当时的痛点。期待后续模块化的发展历程