[[toc]]

# 要解决的问题

* 复用其他人的工作
* 不需要运行时支持动态链接库
* 不共享源代码的前提下复用

# 解决方案

静态链接库 linker 在构建的时期，把 object 文件和 static library 预先合并成一个完整的 executable。
object 与 static library 彼此知道自己的链接关系。

## 构成

| 构成 | 解释 |
| --- | --- |
| object | 第一方开发的可执行代码 |
| static library | 第三方开发的可执行代码 |
| static library linker | 把 object 和 static library 链接成 executable |
| executable | 给 executor 执行用的可执行文件 |

## 衍生的问题

* [如何标识并定位静态链接库](static-library-resolver.md)
* [如何兼容不同的静态链接库的格式](static-library-adapter.md)
* [如何减少静态链接后的文件大小](symbol-stripper.md)

# 解决方案案例

## browserify

browserify 把多个 CJS 格式的静态链接库链接成一个完整自包含的executable。

### main.js
<<< @/static-library-linker/browserify/main.js

### build.sh
<<< @/static-library-linker/browserify/build.sh

| 构成 | 解释 |
| --- | --- |
| object | main.js |
| static library | uniq |
| static library linker | browserify |
| executable | bundle.js |

## webpack

webpack 是一个更全面的静态链接库的链接器，是 browserify 的改进。

### src/main.js
<<< @/static-library-linker/webpack4/src/main.js

### build.sh
<<< @/static-library-linker/webpack4/build.sh

| 构成 | 解释 |
| --- | --- |
| object | src/main.js |
| static library | uniq |
| static library linker | webpack |
| executable | dist/bundle.js |

## rollup

rollup 的出发点是改进静态链接库的格式。
采用 ES6 Module 的导入导出语法，从而实现 tree-shaking 来减少静态链接后的文件尺寸。

### src/main.js
<<< @/static-library-linker/rollup/src/main.js

### rollup.config.js
<<< @/static-library-linker/rollup/rollup.config.js

### build.sh
<<< @/static-library-linker/rollup/build.sh

| 构成 | 解释 |
| --- | --- |
| object | src/main.js |
| static library | uniq |
| static library linker | rollup |
| executable | dist/bundle.js |

## gcc

ar 把多个 linux 下用 c 编译出来的 .o 文件，合并成 .a 静态链接文件

gcc 把 .o 文件和 .a 文件静态链接成 executable

### file1.c
<<< @/static-library-linker/gcc/file1.c

### file2.c
<<< @/static-library-linker/gcc/file2.c

### build1.sh
<<< @/static-library-linker/gcc/build1.sh

演示静态链接库如何产生出来

### main.c
<<< @/static-library-linker/gcc/main.c

### build2.sh
<<< @/static-library-linker/gcc/build2.sh

演示静态链接库如何被使用


| 构成 | 解释 |
| --- | --- |
| object | file1.o file2.o main.o |
| static library | lib12.a |
| static library linker | gcc |
| executable | ./main |