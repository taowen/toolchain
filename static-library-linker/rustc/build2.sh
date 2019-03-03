rustc main.rs -L . # 链接当前目录下的 libfile1.a libfile2.a

./main # 返回值应该是 1 + 2 = 3
echo $?