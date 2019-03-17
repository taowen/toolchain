[[toc]]

# 要解决的问题

如何动态修改已有代码

一个系统通常划分成多个纵向的功能模块，各个模块之间除了包括自身模块的业务逻辑实现以外，还有些共同的代码实现，例如权限，日志，监控等，这部分代码可以从各个模块中抽取出来，作为横向的切面服务于其他模块。面向切面编程可使用代码动态增强技术来实现。
编译期替换原有代码；对构建打包过程透明；不影响debug；对框架透明。


# 解决方案
Java中可使用asm实现类、方法的字节码的CRUD。


# 解决方案案例

## 方法调用前后添加日志

通过asm生成新的class文件，在Bean类的hello方法的body前后添加日志输出的代码。自定义类加载器，动态加载修改后的代码。为避免实例的ClassCastException，为Bean定义了一个接口。

### 添加asm依赖

pom依赖： 

 ~~~xml
 <dependency>
     <groupId>org.ow2.asm</groupId>
     <artifactId>asm</artifactId>
     <version>7.1</version>
 </dependency>
 ~~~


### 修改后的class文件

javap -c Bean.class：

~~~java
 public static void hello();
    Code:
       0: getstatic     #18                 // Field java/lang/System.out:Ljava/io/PrintStream;
       3: ldc           #20                 // String hello: method start
       5: invokevirtual #26                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
       8: getstatic     #18                 // Field java/lang/System.out:Ljava/io/PrintStream;
      11: ldc           #35                 // String Hello, world
      13: invokevirtual #26                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
      16: getstatic     #18                 // Field java/lang/System.out:Ljava/io/PrintStream;
      19: ldc           #32                 // String hello: method end
      21: invokevirtual #26                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
      24: return
~~~
在正常的hello方法body前后添加了start，end输出。

### 代码路径
<<< @/code-enhancer/BeanService.java  
<<< @/code-enhancer/Bean.java   
<<< @/code-enhancer/AsmModifyMethod.java
