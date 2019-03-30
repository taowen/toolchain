# 要解决的问题
在探索前端构建速度的时候，构建工具的诉求大致包含这样几个部分：构建配置、构建速度、产出文件大小（加载速度）以及产出文件的可读性。构建工具的痛点在于很难平衡这些特性，追求极致的构建速度，往往需要因业务场景进行复杂的配置，产出文件过大或者代码可读性很低，或者追求文件大小或者代码可读性，往往会增加更加构建的速度和文件产出。webpack没有做到面面俱到，需要探索别的构建工具的特性，好让我们在前端开发中因地制宜地选择工具，做到更好的开发体验。

# 解决方案
本文探索的是新工具（其实也不新了）——rollup.js，该工具在构建方面有着自己的优势：文件产出较小，构件速度够快（加载执行很快）且易读。但是也有缺陷：插件生态的薄弱导致很多业务场景的构建不能实现，循环依赖不好处理，因此出现问题不方便排查。相较于webpack ，rollup更加适合类库的打包，部分的业务开发场景也可以使用rollup。

# 解决方案案例
此次案例里通过最基本的简单例子来看比较rollup.js和webpack的打包区别，案例代码如下：
```
// ./module.js
export function f1() {
  console.log('we will rollup this function')
}

export function f2() {
  console.log('we will ignore this function')
}

// ./entry.js
import * as modules from './module.js'

modules.f1()
```
构建工具相应的配置如下：
```
// webpack.pro.js
module.exports = {
  entry: path.resolve(__dirname, '../src/rollupvswebpack/entry.js'),
  output: {
    path: path.resolve(__dirname, '../'),
    filename: 'webpack_bundle.js'
  }
}

// rollup.config.js
export default {
  input: path.resolve(__dirname, '../src/rollupvswebpack/entry.js'),
  output: {
    file: path.resolve(__dirname, '../bundle_rollup_file.js'),
    format: 'iife'
  }
}
```
重点来了，看下构建效率和构架结果，先看webpack的结果

![webpack构建](https://user-images.githubusercontent.com/11829615/54578553-22453e80-4a3b-11e9-8443-8e9d061513cf.jpg)

webpack的构建时间为52ms，构建结果为3k，看看核心结果
```
(function(){
  ...
  __webpack_require__() {
    ...
  }
  ...
})([
  (function() { ... }),
  (function(module, __webpack_exports__, __webpack_require__) {

  "use strict";
  /* harmony export (immutable) */ __webpack_exports__["a"] = f1;
  /* unused harmony export f2 */
  function f1() {
    console.log('we will rollup this function')
  }

  function f2() {
    console.log('we will ignore this function')
  }

  /***/ })
])
```

下面是rollup的构建结果

![rollup构建](https://user-images.githubusercontent.com/11829615/54578585-3e48e000-4a3b-11e9-8790-67cf00dcb372.jpg)

rollup的构建时间为20ms，构建结果为100B左右，rollup的构建结果内容很少
```
// iife
(function () {
  'use strict';

  function f1() {
    console.log('we will rollup this function');
  }

  f1();

}());

// cjs
'use strict';

function f1() {
  console.log('we will rollup this function');
}

f1();
```
看到这里有没有很惊喜，在这里例子里rollup在构建速度、构建结果和代码可读性方面全面超越了webpack。

仔细分析一下rollup的结果，会发现rollup会进行无用代码的删除工作，只针对需要的代码进行打包，也就是常常听说到的tree-shaking。关于tree-shaking又是另一个值得探讨话题，tree-shaking在大型应用里效果不是很好，在这里不详细探讨，当然在小型应用里的效果还是很好的。

# rollup
js模块打包，小模块打包成大模块。更多的是做library或者应用程序。充分利用es6特性，只打包你需要代码块中的特定部分，不打包不使用的代码。定位在构建js库，可以构架大部分的应用程序（拆分代码和动态导入）

rollup的配置规则跟webpack相似但相较于webpack更加简洁，规则如下：
```
export default {
  input: 'entry.js',
  output: {
    format: 'umd',  // umd, cjs, iife, es
    file: 'output.js'
  },
  plugin: [
    resolve(),
    commonjs(),
    eslint({
      include: [],
      exclude: []
    }),
    babel(),
    (ENV === 'production' && uglify())
  ]
}
```
format决定了文件产出的格式，目前有四种格式umd/cjs/iife/es。

plugins在rollup构建当中扮演着重要的功能，这可以从这几个重要的插件功能中窥探一二。

## resolve()
如果我们在使用一个三方库js库，我们还需要手动下载这个库是个很糟糕的体验，因此在rollup里为了解决这个问题，提供了resolve插件，将三方库和我们手写的源码进行合并。

在使用方面也就是简单的引用resolve()这样一个方法即可。该方法会自动对三方库进行集成。
```
// 安装指令
npm i -D rollup-plugin-node-resolve
```
另外，与resolve()配合使用的有另一个external属性，改方法就是制定哪些库不要进行集成。更加方便我们定制化。
```
import resolve from 'rollup-plugin-node-resolve'

export default {
  ...
  plugins: [
    resolve()
  ],
  external: ['some-plugin']
}
```
## commonjs()
这个原因很简单，rollup不支持commonJs模块，当代码里引用了commonJs规范的一个包的时候，在构建的时候会报错，无法识别commonJs模块。此时需要插件的支持
```
// 安装指令
npm i -D rollup-plugin-commonjs
```
但是，commonJs模块不支持tree-shaking特性，有兴趣的可以自己试一下，commonJs模块的无用参数和方法不会在打包的时候删除掉。因此，推荐大家在使用rollup.js的时候尽量使用ES模块，享受更精简的代码体验。

关于tree-shaking这部分还有个小技巧，就是在此类模块的package.json的配置里面配置module属性，rollup.js默认情况下会优先寻找并加载module属性指向的模块。

```
{
  "name": "some-plugin",
  "version": "0.0.1",
  "description": "test",
  "main": "dist/some-plugin.js",
  "module": "dist/some-plugin-es.js"
}
```
如上述配置，rollup首先会执行es模块，也就可以享受到tree-shaking的特性。关于这个优先规则rollup.js官方的说明如下：
> 在 package.json 文件的 main 属性中指向当前编译的版本。如果你的 package.json 也具有 module 字段，像 Rollup 和 webpack 2 这样的 ES6 感知工具(ES6-aware tools)将会直接导入 ES6 模块版本。

## babel()
关于babel这部分就更好理解了，当我们代码是es6规范，如果最后产出的还是es6规范的代码，那么在不支持es6的环境下执行代码将会无法执行。这个有点像webpack里的loader，这个效果大家应该都很清楚，这里就不详细说明了。
```
// 安装指令
npm i -D rollup-plugin-babel
```

## uglify()
这个插件也没有可更多解释的部分，只是有个提示：uglify不支持ES模块和ES6语法，也就是说我们在打包前需要进行babel转化，然后压缩打包。

从rollup配置来看，插件在业务需求开发中扮演者重要的角色，rollup生态里插件的丰富程度决定了这个构建工具的适用场景。目前看来一些业务需求场景基本都能实现。

rollup有watch文件修改重新打包能力，但是目前rollup本身还没有聚合浏览器重新热加载的能力，需要引用另外的包来实现该功能。从这样一个行为也可以看出rollup工具的定位：js模块打包器，更多的是对库的打包。

# 总结
各种构建工具的执行流程几近相似：文件目录获取（入口文件获取）、（文件分析）、文件编译、文件打包、文件产出。但每个库的定位不同，因此在配置和执行方面各不相同

RollUp.js相较于webpack在配置上更加简洁，构建速度也更快，构建的结果也相对更加易读。但是RollUp上述的特性也是受限于本身生态的规模或者说本身工具的定位。作为一个库打包器，配置方面相较于大型应用应该是要简单一些。

从RollUp的使用体验来看，该工具的定位十分明确。面对工具库的开发或者一些小的业务需求开发，推荐使用RollUp来构建，除了享用构建的各种优势还能够体验tree-shaking的优势。在大型应用面前webpack优势巨大，可以满足构建需求就是优势本身，而且强大的生态可以满足任何场景和问题追溯，使用放心。

# 思考
关于tree-shaking这部分在大型应用里的效率值得探讨，tree-shaking的结果是好的，但是tree-shaking的构建过程相较于没有tree-shaking的构建是不是一个更加耗时的过程？这个消耗有多大影响？