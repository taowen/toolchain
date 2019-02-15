# 用 AMD 的格式输出 object 文件
# main.ts  和 lib.ts 被链接成一个 object.js
tsc --outFile dist/object.js --module amd src/main.ts
