gcc -c main.c # 创建 main.o，引用 hello / world 这两个符号
gcc -o main main.o lib12.a # 把静态链接库链接成 executable

./main # 返回值应该是 1 + 2 = 3
echo $?
