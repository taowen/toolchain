# 要解决的问题
webpack构建过程中的有两个部分是直接影响构建效率的，一个是文件的编译，另一个则是文件的分类打包。相较之下文件的编译更为耗时，而且在Node环境下文件只能一个一个去处理，因此这块的优化需要解决。

# 解决方案
这里引入的HappyPack这样一个插件，在webpack构建过程中，我们需要使用Loader对js，css，图片，字体等文件做转换操作，并且转换的文件数据量也是非常大的，且这些转换操作不能并发处理文件，而是需要一个个文件进行处理，HappyPack的基本原理是将这部分任务分解到多个子进程中去并行处理，子进程处理完成后把结果发送到主进程中，从而减少总的构建时间。

# 解决方案案例
同样拿react + antd的常规项目来分析，在没有使用HappyPack的项目，我们的构建结果如下：

![non-happypack](https://user-images.githubusercontent.com/11829615/54013269-75a1cc00-41b3-11e9-8fa0-384972897ace.jpg)

接下来我们在项目里引用HappyPack，看看会发生什么？先直接看执行效果。

![happypack](https://user-images.githubusercontent.com/11829615/54013314-9bc76c00-41b3-11e9-9fa2-e634e3717e4a.jpg)

HappyPack只是开启了三个线程，不过还是可以看到构建速度上面的提升。

# HappyPack
> HappyPack是让webpack对loader的执行过程，从单一进程形式扩展为多进程模式，也就是将任务分解给多个子进程去并发的执行，子进程处理完后再把结果发送给主进程。从而加速代码构建 与 DLL动态链接库结合来使用更佳。

惯例还是先看下使用规则：
```
const HappyPack = require('happypack')
const os = require('os')
const happyThreadPool = HappyPack.ThreadPool({size: os.cpus().length})

module.exports = {
  module: {
    rules: [
      {
        test: /\.js$/,
        // 将.js文件交给id为happyBabel的happypack实例来执行
        loader: 'happypack/loader?id=happyBabel',
        exclude: /node_modules/
      }
    ]
  },
  plugins: [
    new HappyPack({
      // id标识happypack处理那一类文件
      id: 'happyBabel',
      // 配置loader
      loaders: [{
        loader: 'babel-loader?cacheDirectory=true'
      }],
      // 共享进程池
      threadPool: happyThreadPool,
      // 日志输出
      verbose: true
    })
  ]
}
```
使用规则很好理解，针对webpack常规的loader部分直接引用happypack的loader，同时参数id是关联HappyPack插件的唯一标识。在插件里示例化一个HappyPack实例，这里引用真实的编译loader，同时制定线程池等参数。这部分流程上的原理下面会做介绍。

# How It Works?
![HappyPack](https://raw.githubusercontent.com/amireh/happypack/master/doc/HappyPack_Workflow.png)

上图是HappyPack在webpack里作用的一个整体流向。

因篇幅问题，源码部分通过伪代码的形式来解释逻辑关系
```
// happypack插件
function HappyPlugin() {
  // 基础配置

  // 线程池初始化，创建线程
    
  // 编译缓存初始化
}
```
基础配置
```
// 基础配置
...
this.id = String(userConfig.id || ++uid)
this.name = 'HappyPack'
this.state = {
  ...
  // 输出日志
  verbose: true,
  // 实际运行的Loader
  loaders: [],
  // 开启缓存
  cache: true
  ...
}
...
```
- id：关联loader和plugin；
- name：与id一同使用限制HappyPack的获取，因为可能存在其他plugin的id同名；
- 其他参数：跟插件运行过程中的一些参数进行配置

线程初始化
```
// 线程初始化
// HappyThreadPool.js
  // HappyRPCHandler.js 绑定当前运行的loader与compiler，同时在文件中，针对Loader和compiler定义调用接口
  // HappyThread.js  createThread的核心，通过实际的size来创建真实的thread的个数，这里有控制进程的相关控制hook
    // HappyWorkerChannel.js  子进程执行文件，stream.on('message')来订阅消息，传递给主线程
      // HappyWorker.js 子进程编译逻辑文件 fs.writeFileSync到compilePath这个编译路径
```
编译缓存初始化
```
// 编译缓存，在编译开始和woker编译完成的时候进行缓存加载、设置等操作
// HappyFSCache.js
  // getCompiledSourceCodePath() 获取缓存内容的物理位置
  // updateMTimeFor 更新新编译文件缓存的设置
```
最后有个遗漏问题，在webpack配置里面，默认Loader里配置的happy/loader做了呢？
```
// HappyLoader 首先拿到配置 id ,然后对所有的 webpack的plugins进行遍历
// 找到Id匹配的happypackPlugin
```
# 总结
在基于webpack的配置构建中，构建流程上会先走到插件部分，HappyPack插件会先因为webpack的run执行一系列的初始化，为后续的多线程执行做准备，这里的初始化包括：基础配置、线程初始化和编译缓存初始化。接下来走到webpack流程上的文件编译，此时会调用基础配置里的happy/loader，此loader会通过参数的id遍历真实的插件数组，找到对应的happyPlugin，通过happyPlugin的配置获取真实的Loader并通过之前初始化完成的多线程进行编译，将编译结果传递给主线程。编译完成后，插件还会针对编译的结果缓存，以及新编译的文件进行缓存的设置。

# 几个问题
1、第一个问题摘自官方文档
> Is it necessary for Webpack 4?
> 
> Short answer: maybe not.
> 
> Long answer: there's now a competing add-on in the form of a loader for processing files in multiple threads, exactly what HappyPack does. The fact that it's a loader and not a plugin (or both, in case of H.P.) makes it much simpler to configure. Look at thread-loader and if it works for you - that's great, otherwise you can try HappyPack and see which fares better for you.

从上面的解释看来，webpack4已经融合了多线程机制，因此happypack的作用不是很明显。如果你使用的版本是<4，那么还是可以继续使用HappyPack。

2、线程的使用

线程在使用过程中不仅可以指定多个实例分别各自执行指定的文件类型，还可以通过共享的形式共享线程池，如下：
```
// @file: webpack.config.js
var HappyPack = require('happypack');
var happyThreadPool = HappyPack.ThreadPool({ size: 5 });

module.exports = {
  // ...
  plugins: [
    new HappyPack({
      id: 'js',
      threadPool: happyThreadPool,
      loaders: [ 'babel-loader' ]
    }),

    new HappyPack({
      id: 'styles',
      threadPool: happyThreadPool,
      loaders: [ 'style-loader', 'css-loader', 'less-loader' ]
    })
  ]
};

```
上述的js和样式公用size为5的线程池，最大化利用线程的空闲时间。