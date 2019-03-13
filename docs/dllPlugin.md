# 要解决的问题
如何提升webpack的构建效率，尤其在面对后台管理应用依赖很多库的情况下，构建效率变得尤为重要，这将直接影响开发人员的工作效率。
# 解决方案
提升构建效率的办法有很多，这是一个系列探索问题。当前方案是通过DllPlugin的方式，即提前构建依赖库（只会在业务开发中引用而不会被修改的库）并产出一个配置文件manifest.json，通过这个配置文件manifest.json来引用提前打包好的库，这样每次构建的成本只有业务逻辑代码部分，效率提升是明显的。
# 解决方案案例
拿react + antd的常规项目来分析，在没有使用DllPlugin的项目，我们的构建结果如下：

![常规构建](https://user-images.githubusercontent.com/11829615/53709305-892afb00-3e72-11e9-9822-4ce3cfc28864.jpg)

接下来我们在项目里引用DllPlugin，看看会发生什么？先直接看执行效果。

![dll构建](https://user-images.githubusercontent.com/11829615/53709328-a9f35080-3e72-11e9-8e96-506e7ddf3738.jpg)

构建流程使用DllPlugin效果还是十分明显，构建时间从38秒减少到了10秒，这大大提升我们的开发构建效率。那么DllPlugin是如何应用到我们的项目里的呢？以及如何做到效率的提升？

## DllPlugin & DllReferencePlugin
> The DllPlugin and DllReferencePlugin provide means to split bundles in a way that can drastically improve build time performance.

简单解释这个插件的功能：DllPlugin它能够把第三方库分离开，每次项目代码更改的时候，它只会打包业务代码本身，因此节省了三方库的打包时间，打包速度更快，也就提升了打包效率。

首先我们看下这个插件如何应用在项目里：

DllPlugin插件的使用有两个部分组成，webpack.dll.conf.js中定义DllPlugin打包规则和webpack.conf.js中定义DllReferencePlugin的引用规则。
首先在webpack.conf.js同级目录下建立一个webpack.dll.conf.js文件，配置内容如下：
```
const path = require('path')
const webpack = require('webpack')

module.exports = {
  entry: {
    vendor: ['react-dom', 'antd', 'react', 'react-router', 'react-redux']
  },
  output: {
    path: path.resolve(__dirname, '../app/public'),
    filename: '[name].dll.js',
    library: '[name]_library'
  },
  plugins: [
    new webpack.DllPlugin({
      path: path.join(__dirname, '.', '[name]-manifest.json'),
      name: '[name]_library',
      context: __dirname
    })
  ]
}
```
这个文件里跟其他配置结构一致，逻辑也很清晰：将一些三方库打包成vendor.dll.js并放在特定的目录下，重点看下插件部分，这部分会产出一个manifest.json文件，这个文件的作用十分关键，是业务调用三方库的“使用说明书”，这个说明书功能的体现在DllReferencePlugin中得以体现。

接着在webapck.conf.js中配置DllReferencePlugin的调用：
```
...
module.exports = {
  ...
  plugins: [
    ...
    // 核心是引用manifest.json
    new webpack.DllReferencePlugin({
      context: __dirname,
      manifest: require('./vendor-manifest.json')
    })
  ]
  ...
}
...
```
不出意外，DllReferencePlugin在构建主流程中只是引用了一下manifest.json文件。

此时先构建webpack.dll.conf.js，执行完毕产出vendor.dll.js和vendor.manifest.js两个文件。

然后接着构建主流程webpack.conf.js文件，也就是上述的图二的结果，构建效果明显。此时你是不是觉得可以很愉快的跑页面了？

年轻人不要着急，有个重要的配置需要设置，手动引用vendor.dll.js在html页面里，这里很容易理解：因为”说明书“只是三方库的一个映射，因此需要提前加载这样一个真实库文件。引用如下：
```
...
<body>
  <div id="app"></div>
  <script src="/public/vendor.dll.js"></script>
</body>
...
```
此时，页面可以正常打开，资源也正常加载。

## 3.2 那么DllPlugin & DllReferencePlugin究竟如何做到的呢？
看下DllPlugin源码部分
```
...
apply(compiler) {
  compiler.hooks.entryOption.tap("DllPlugin", (context, entry) => {
    const itemToPlugin = (item, name) => {
      if (Array.isArray(item)) {
  return new DllEntryPlugin(context, item, name);
      }
      throw new Error("DllPlugin: supply an Array as entry");
    };
    if (typeof entry === "object" && !Array.isArray(entry)) {
      Object.keys(entry).forEach(name => {
  itemToPlugin(entry[name], name).apply(compiler);
      });
    } else {
      itemToPlugin(entry, "main").apply(compiler);
    }
    return true;
  });
  new LibManifestPlugin(this.options).apply(compiler);
  if (!this.options.entryOnly) {
    new FlagInitialModulesAsUsedPlugin("DllPlugin").apply(compiler);
  }
}
...
```
三个函数分别做了三件事：
- itemToPlugin：通过addEntry记录Dll的关联关系，进而可以通过manifest通过对应关系来调用
- LibManifestPlugin：毫无疑问产出manifest.json关系配置文件
- FlagInitialModulesAsUsedPlugin：treeShake的用途（暂不介绍）

再来看看DllReferencePlugin的部分核心源码
```
// DelegatedModuleFactoryPlugin.js
// 该方法里重要的核心就是通过mainfest将业务代码里的引用和dll里的真实对应的方法关联起来
// 下面的方法通过解析manifest，这个模块方法进行代理 
const request = module.libIdent(this.options);
if (request && request in this.options.content) {
  const resolved = this.options.content[request];
  return new DelegatedModule(
    this.options.source,
    resolved,
    this.options.type,
    request,
    module
  );
}

// DelegatedModule.js
// 代理的核心就是绑定id，也就是下面代码的this.request，这个id会在dll.js找到对应的源码，从而执行源码部分
source(depTemplates, runtime) {
  const dep = /** @type {DelegatedSourceDependency} */ (this.dependencies[0]);
  const sourceModule = dep.module;
  let str;

  if (!sourceModule) {
    str = WebpackMissingModule.moduleCode(this.sourceRequest);
  } else {
    str = `module.exports = (${runtime.moduleExports({
      module: sourceModule,
      request: dep.request
    })})`;

    switch (this.type) {
      case "require":
        str += `(${JSON.stringify(this.request)})`;
        break;
      case "object":
        str += `[${JSON.stringify(this.request)}]`;
        break;
    }

    str += ";";
  }

  if (this.useSourceMap) {
    return new OriginalSource(str, this.identifier());
  } else {
    return new RawSource(str);
  }
}

// 例如在app.js里产生 n(120)这样的方式，120是dll.js源码方法序列里相应源码的序列号。
```
简单从打包结果看流程，dll.js会生成一个vendor_library的方法，方法的入参就是打包的所有的一个一个三方库方法形成的数组，且调用的逻辑按照方法在数组里的序号来算。我们看下打包后的代码样式
```
// dll.js因篇幅问题，只展示核心部分
var vendor_library = function(e) {
  ...
  // 核心方法，这个方法就是调用入参数组里的某个对应方法
  function n(r) {
    if (t[r]) return t[r].exports;
    var o = t[r] = {
      i: r,
      l: !1,
      exports: {}
    };
    // r代表着在入参数组里的顺序，将方法绑定实例和参数
    return e[r].call(o.exports, o, o.exports, n),
    o.l = !0,
    o.exports
  }
  ...
  return ...,n(941)
}([fun1, fun2, fun3, ..., funN])
```
上述代码里方法数组则是三方库的所有方法，那么在app.js里会对外暴露vendor_library这个调用方法，接着app里需要引用到哪个三方库的方法，直接调用vendor_library里的序列即可。看下在app.js里是如何引用这个方法的：
```
// app.js
!function(){
  ...
}({
  ...
  AOTY: function(t, e) {
    t.exports = vendor_library
  },
  ANjH: function(t, e, n) {
    t.exports = n("AOTY")(936)
  },
  ...
  FTny: function(t, e, n) {
    var r, o = n("ANjH"), i = (r = n("IK7v")) && r.__esModule ? r : {
      default: r
    };
    e.default = (0,
    o.combineReducers)({
      home: i.default
    })
  },
  ...
})
```
上述代码完整阐述了暴露vendor_library，然后到调用dll.js里对应的某个函数，同时被调用的方法在另一个方法里引用。再看看manifest.json的格式，从这里对比我们可以发现原来这里是调用了redux.js这个三方库。
```
// manifest.json
{
  "name": "vendor_library",
  "content": {
    ...
    "../node_modules/redux/es/redux.js":{
      "id":936,
      "buildMeta":{
        "exportsType":"namespace",
        "providedExports":[
          "createStore",
          "combineReducers",
          "bindActionCreators",
          "applyMiddleware",
          "compose",
          "__DO_NOT_USE__ActionTypes"
        ]
      }
    },
    ...
  }
}
```
# 总结
DllPlugin方案的核心是：提前打包指定三方库，同时产出manifest.json配置问题，同时在构建核心主流程里通过DllReferencePlugin引用manifest.json来解析代理Dll.js里打包的源码，从而在我们的业务代码里通过调用对应代理需要从而直接获取到提前预加载的三方包Dll.js对外暴露的方法，完成源码的引用加载。

该方法极大程度的降低了我们的构建速度，让每次的构建只关系业务代码部分。就目前使用程度上来说可能存在的缺点：
- 需要提前构建
- 需要明确指定三方包，而且三方包会很大
- 需要手动引用dll包在html头部
- 每增减一个三方包都需要手动构建，重新引用

后续的其他方案探索看看有没有更灵活的方案。)