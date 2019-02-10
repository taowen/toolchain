# 要解决的问题

如何引用动态链接库指定的符号

# 解决方案

动态链接库提供了很多个函数，类或者常量之类的东西，我们统称为symbol。引用了具体的符号才能开始使用。

## 构成

* exported symbol: 动态链接库导出的符号
* imported symbol：导入的符号
* binder：把导入的符号和导出的符号绑定到一起

# 解决方案案例

## 传统浏览器

传统浏览器没有 binder。全局共享一个namespace（window对象），动态链接库导出和导入通过读写全局变量实现。

## ES6 Module

* exported symbol
  * `export function my_func() {}` 定义和export合一
  * `function my_func() {}; export {my_func as my_exported_func}` 定义和export分离
  * `export default function my_func() {}` export成为特殊的符号`default`
* imported symbol
  * `import './library.mjs'` 只加载动态链接库，但是并不导入符号
  * `import {my_func} from './library.mjs'` 导入动态链接库中的指定符号 my_func
  * `import my_func from 
