# 要解决的问题
  如何进行网站性能分析？
# 解决方案
  网站性能不仅关乎用户的体验，甚至关乎用户去留的问题。轻微性能问题可能只会导致微弱的延迟，给您的用户带来短暂的不便。 而严重的性能问题可能导致您的网站完全无法访问。认识到了性能对于web开发的重要同时为了解决性能当前性能测试的困难，W3C推出了一套性能API标准，各种浏览器对这套标准的支持如今也逐渐成熟起来
# API标准详解：
## window.performance 详解
  ```
    var performance = {  
    // memory 是非标准属性，只在 Chrome 有
    // 财富问题：我有多少内存
    memory: {
        usedJSHeapSize:  16100000, // JS 对象（包括V8引擎内部对象）占用的内存，一定小于 totalJSHeapSize
        totalJSHeapSize: 35100000, // 可使用的内存
        jsHeapSizeLimit: 793000000 // 内存大小限制
    },
 
    //  哲学问题：我从哪里来？
    navigation: {
        redirectCount: 0, // 如果有重定向的话，页面通过几次重定向跳转而来
        type: 0           // 0   即 TYPE_NAVIGATENEXT 正常进入的页面（非刷新、非重定向等）
                          // 1   即 TYPE_RELOAD       通过 window.location.reload() 刷新的页面
                          // 2   即 TYPE_BACK_FORWARD 通过浏览器的前进后退按钮进入的页面（历史记录）
                          // 255 即 TYPE_UNDEFINED    非以上方式进入的页面
    },
 
    timing: {
        // 在同一个浏览器上下文中，前一个网页（与当前页面不一定同域）unload 的时间戳，如果无前一个网页 unload ，则与 fetchStart 值相等
        navigationStart: 1441112691935,
 
        // 前一个网页（与当前页面同域）unload 的时间戳，如果无前一个网页 unload 或者前一个网页与当前页面不同域，则值为 0
        unloadEventStart: 0,
 
        // 和 unloadEventStart 相对应，返回前一个网页 unload 事件绑定的回调函数执行完毕的时间戳
        unloadEventEnd: 0,
 
        // 第一个 HTTP 重定向发生时的时间。有跳转且是同域名内的重定向才算，否则值为 0
        redirectStart: 0,
 
        // 最后一个 HTTP 重定向完成时的时间。有跳转且是同域名内部的重定向才算，否则值为 0
        redirectEnd: 0,
 
        // 浏览器准备好使用 HTTP 请求抓取文档的时间，这发生在检查本地缓存之前
        fetchStart: 1441112692155,
 
        // DNS 域名查询开始的时间，如果使用了本地缓存（即无 DNS 查询）或持久连接，则与 fetchStart 值相等
        domainLookupStart: 1441112692155,
 
        // DNS 域名查询完成的时间，如果使用了本地缓存（即无 DNS 查询）或持久连接，则与 fetchStart 值相等
        domainLookupEnd: 1441112692155,
 
        // HTTP（TCP） 开始建立连接的时间，如果是持久连接，则与 fetchStart 值相等
        // 注意如果在传输层发生了错误且重新建立连接，则这里显示的是新建立的连接开始的时间
        connectStart: 1441112692155,
 
        // HTTP（TCP） 完成建立连接的时间（完成握手），如果是持久连接，则与 fetchStart 值相等
        // 注意如果在传输层发生了错误且重新建立连接，则这里显示的是新建立的连接完成的时间
        // 注意这里握手结束，包括安全连接建立完成、SOCKS 授权通过
        connectEnd: 1441112692155,
 
        // HTTPS 连接开始的时间，如果不是安全连接，则值为 0
        secureConnectionStart: 0,
 
        // HTTP 请求读取真实文档开始的时间（完成建立连接），包括从本地读取缓存
        // 连接错误重连时，这里显示的也是新建立连接的时间
        requestStart: 1441112692158,
 
        // HTTP 开始接收响应的时间（获取到第一个字节），包括从本地读取缓存
        responseStart: 1441112692686,
 
        // HTTP 响应全部接收完成的时间（获取到最后一个字节），包括从本地读取缓存
        responseEnd: 1441112692687,
 
        // 开始解析渲染 DOM 树的时间，此时 Document.readyState 变为 loading，并将抛出 readystatechange 相关事件
        domLoading: 1441112692690,
 
        // 完成解析 DOM 树的时间，Document.readyState 变为 interactive，并将抛出 readystatechange 相关事件
        // 注意只是 DOM 树解析完成，这时候并没有开始加载网页内的资源
        domInteractive: 1441112693093,
 
        // DOM 解析完成后，网页内资源加载开始的时间
        // 在 DOMContentLoaded 事件抛出前发生
        domContentLoadedEventStart: 1441112693093,
 
        // DOM 解析完成后，网页内资源加载完成的时间（如 JS 脚本加载执行完毕）
        domContentLoadedEventEnd: 1441112693101,
 
        // DOM 树解析完成，且资源也准备就绪的时间，Document.readyState 变为 complete，并将抛出 readystatechange 相关事件
        domComplete: 1441112693214,
 
        // load 事件发送给文档，也即 load 回调函数开始执行的时间
        // 注意如果没有绑定 load 事件，值为 0
        loadEventStart: 1441112693214,
 
        // load 事件的回调函数执行完毕的时间
        loadEventEnd: 1441112693215
 
        // 字母顺序
        // connectEnd: 1441112692155,
        // connectStart: 1441112692155,
        // domComplete: 1441112693214,
        // domContentLoadedEventEnd: 1441112693101,
        // domContentLoadedEventStart: 1441112693093,
        // domInteractive: 1441112693093,
        // domLoading: 1441112692690,
        // domainLookupEnd: 1441112692155,
        // domainLookupStart: 1441112692155,
        // fetchStart: 1441112692155,
        // loadEventEnd: 1441112693215,
        // loadEventStart: 1441112693214,
        // navigationStart: 1441112691935,
        // redirectEnd: 0,
        // redirectStart: 0,
        // requestStart: 1441112692158,
        // responseEnd: 1441112692687,
        // responseStart: 1441112692686,
        // secureConnectionStart: 0,
        // unloadEventEnd: 0,
        // unloadEventStart: 0
    }
};
  ```
## 使用performance.timing简单计算出网页性能数据
```
// 计算加载时间
function getPerformanceTiming () {  
    var performance = window.performance;
 
    if (!performance) {
        // 当前浏览器不支持
        console.log('你的浏览器不支持 performance 接口');
        return;
    }
 
    var t = performance.timing;
    var times = {};
 
    //【重要】页面加载完成的时间
    //【原因】这几乎代表了用户等待页面可用的时间
    times.loadPage = t.loadEventEnd - t.navigationStart;
 
    //【重要】解析 DOM 树结构的时间
    //【原因】反省下你的 DOM 树嵌套是不是太多了！
    times.domReady = t.domComplete - t.responseEnd;
 
    //【重要】重定向的时间
    //【原因】拒绝重定向！比如，http://example.com/ 就不该写成 http://example.com
    times.redirect = t.redirectEnd - t.redirectStart;
 
    //【重要】DNS 查询时间
    //【原因】DNS 预加载做了么？页面内是不是使用了太多不同的域名导致域名查询的时间太长？
    // 可使用 HTML5 Prefetch 预查询 DNS ，见：[HTML5 prefetch](http://segmentfault.com/a/1190000000633364)            
    times.lookupDomain = t.domainLookupEnd - t.domainLookupStart;
 
    //【重要】读取页面第一个字节的时间
    //【原因】这可以理解为用户拿到你的资源占用的时间，加异地机房了么，加CDN 处理了么？加带宽了么？加 CPU 运算速度了么？
    // TTFB 即 Time To First Byte 的意思
    // 维基百科：https://en.wikipedia.org/wiki/Time_To_First_Byte
    times.ttfb = t.responseStart - t.navigationStart;
 
    //【重要】内容加载完成的时间
    //【原因】页面内容经过 gzip 压缩了么，静态资源 css/js 等压缩了么？
    times.request = t.responseEnd - t.requestStart;
 
    //【重要】执行 onload 回调函数的时间
    //【原因】是否太多不必要的操作都放到 onload 回调函数里执行了，考虑过延迟加载、按需加载的策略么？
    times.loadEvent = t.loadEventEnd - t.loadEventStart;
 
    // DNS 缓存时间
    times.appcache = t.domainLookupStart - t.fetchStart;
 
    // 卸载页面的时间
    times.unloadEvent = t.unloadEventEnd - t.unloadEventStart;
 
    // TCP 建立连接完成握手的时间
    times.connect = t.connectEnd - t.connectStart;
 
    return times;
}
```
## 使用performance.getEntries() 详解
```
    var entry = {  
    // 资源名称，也是资源的绝对路径
    name: "http://cdn.alloyteam.com/wp-content/themes/alloyteam/style.css",
    // 资源类型
    entryType: "resource",
    // 谁发起的请求
    initiatorType: "link", // link 即 <link> 标签
                           // script 即 <script>
                           // redirect 即重定向
    // 加载时间
    duration: 18.13399999809917,
 
    redirectStart: 0,
    redirectEnd: 0,
 
    fetchStart: 424.57699999795295,
 
    domainLookupStart: 0,
    domainLookupEnd: 0,
 
    connectStart: 0,
    connectEnd: 0,
 
    secureConnectionStart: 0,
 
    requestStart: 0,
 
    responseStart: 0,
    responseEnd: 442.7109999960521,
 
    startTime: 424.57699999795295
};

```
## 使用performance.getEntries() 获取某个资源的时间
```
    // 计算加载时间
function getEntryTiming (entry) {  
    var t = entry;
    var times = {};
 
    // 重定向的时间
    times.redirect = t.redirectEnd - t.redirectStart;
 
    // DNS 查询时间
    times.lookupDomain = t.domainLookupEnd - t.domainLookupStart;
 
    // 内容加载完成的时间
    times.request = t.responseEnd - t.requestStart;
 
    // TCP 建立连接完成握手的时间
    times.connect = t.connectEnd - t.connectStart;
 
    // 挂载 entry 返回
    times.name = entry.name;
    times.entryType = entry.entryType;
    times.initiatorType = entry.initiatorType;
    times.duration = entry.duration;
 
    return times;
}
 
// test
// var entries = window.performance.getEntries();
// entries.forEach(function (entry) {
//     var times = getEntryTiming(entry);
//     console.log(times);
// });
```
## performance.now()详解
performance.now() 与 Date.now() 不同的是，返回了以微秒（百万分之一秒）为单位的时间，更加精准。

并且与 Date.now() 会受系统程序执行阻塞的影响不同，performance.now() 的时间是以恒定速率递增的，不受系统时间的影响（系统时间可被人为或软件调整）。

注意 Date.now() 输出的是 UNIX 时间，即距离 1970 的时间，而 performance.now() 输出的是相对于 performance.timing.navigationStart(页面初始化) 的时间。

使用 Date.now() 的差值并非绝对精确，因为计算时间时受系统限制（可能阻塞）。但使用 performance.now() 的差值，并不影响我们计算程序执行的精确时间。
```
    function getFunctionTimeWithPerformance (func) {  
    var timeStart = window.performance.now();
 
    // 执行开始
    func();
    // 执行结束
    var timeEnd = window.performance.now();
 
    // 返回执行时间
    return (timeEnd - timeStart);
}

```
## window.performance.measure 详解
```
var t = performance.timing  
var domReadyTime = t.domComplete - t.responseEnd;  
console.log(domReadyTime)

or: 
window.performance.measure('domReady','responseEnd' , 'domComplete');  
var domReadyMeasure = window.performance.getEntriesByName('domReady');  
console.log(domReadyMeasure);

看下保存起来的测量 measure
window.performance.getEntriesByType('measure')
// 清除指定测量
window.performance.clearMeasures('domReady');  
// 清除所有测量
window.performance.clearMeasures();
看下保存起来的测量 measure
window.performance.getEntriesByType('measure') 


```

## window.performance.mark 详解
```
function randomFunc (n) {  
    if (!n) {
        // 生成一个随机数
        n = ~~(Math.random() * 10000);
    }
    var nameStart = 'markStart' + n;
    var nameEnd   = 'markEnd' + n;
    // 函数执行前做个标记
    window.performance.mark(nameStart);
 
    for (var i = 0; i < n; i++) {
        // do nothing
    }
 
    // 函数执行后再做个标记
    window.performance.mark(nameEnd);
 
    // 然后测量这个两个标记间的时间距离，并保存起来
    var name = 'measureRandomFunc' + n;
    window.performance.measure(name, nameStart, nameEnd);
}
 
randomFunc();

// 看下保存起来的标记 mark
var marks = window.performance.getEntriesByType('mark');  
console.log(marks);
// 清除指定标记
window.performance.clearMarks('markStart888');  
// 清除所有标记
window.performance.clearMarks();
```