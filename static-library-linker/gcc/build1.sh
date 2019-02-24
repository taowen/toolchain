gcc -c file1.c # 创建 file1.o 提供 hello
gcc -c file2.c # 创建 file2.o 提供 world
ar -r lib12.a *.o # 把两个object文件合并成一个静态链接库
