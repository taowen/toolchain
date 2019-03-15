# 要解决的问题

# rollup
js模块打包，小模块打包成大模块。更多的是做library或者应用程序。充分利用es6特性，只打包你需要代码块中的特定部分，不携带不使用的代码。定位在构建js库，可以构架大部分的应用程序（拆分代码和动态导入）

例子
```
// 入口文件是main，产出的文件是bundle。后续是文件格式以及文件的重命名
$ rollup main.js --file bundle.js --format cjs --name "myBundle"
```

tree-shaking
```
let utils = require('utils')
utils.ajax()

import { ajax } from 'utils'
ajax()
```