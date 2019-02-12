// 引用的 uniq 包会被链接到最终的 executable 里
var unique = require('uniq');
var data = [1, 2, 2, 3, 4, 5, 5, 5, 6];
console.log(unique(data));