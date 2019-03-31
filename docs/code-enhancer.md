[[toc]]

# 要解决的问题

如何动态修改已有代码

一个系统通常划分成多个纵向的功能模块，各个模块之间除了包括自身模块的业务逻辑实现以外，还有些共同的代码实现，例如权限，日志，监控等，这部分代码可以从各个模块中抽取出来，作为横向的切面服务于其他模块。使用动态代理技术会在系统运行期动态生成Proxy类，Caller从调用Callee改为Callee的Proxy，一方面生成大量的代理类，另一方面也会影响系统性能。面向切面编程可使用代码动态增强技术来实现。
##目标
编译期替换原有代码；对构建打包过程透明；不影响debug；对框架透明。


# 解决方案
## 1、Java中可使用asm实现类、方法的字节码的CRUD。
## 2、Java Compiler Annotation Process
JSR269提供一种基于Annotation的编译器插件开发API，允许在编译阶段生成源码，字节码和资源文件。"Don't Repeat Yourself, Generate Your Code."。Javac的编译过程可分为以下几个步骤：  
 
* 1）解析：对java源代码进行词法和语法分析。词法分析把java源码转化为Token流，语法分析把Token流转化为抽象语法树（Abstract syntax Tree， AST），分别对应 ``` com.sun.tools.javac.parser.Scanner```类和```com.sun.tools.javac.parser.Parser```类，该阶段生成的AST由```com.sun.tools.javac.tree.JCTree```类表示，后续步骤都建立在AST基础之上。
* 2）填充符号表：遍历AST生成作用空间中的变量类型，位置，源码行数等信息，即符号表。该阶段生成一个AST TODO列表，该列表需要后续步骤处理，并生成class文件。
* 3）注解处理：通过注解处理器修改被标注的类的AST。
* 4）语义分析和class文件生成：检查AST上的元素是否满足规则，例如变量使用前是否已经声明，变量类型与赋值是否匹配；程序控制流是否满足规则，例如异常是否捕获，final类型变量是否多次赋值；语法糖解码；最终生成class文件。  

JSR269与第3）中的注解处理过程对应，但其有一个局限，只能操作到方法层面，不能修改方法体。需要借助javac编译器本身提供的api操作方法体。


## 3、Lombok
## 4、AutoValue




# 解决方案案例

## 1、方法调用前后添加日志

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

## 2、Assert语句转化为throw exception

运行样例时需要enableassertions：java -ea或java -enableassertiions。

~~~java
public class AssertExample {
    public static void main(String[] args) {
        /**
         * args is not null for ever, so assert output assertion error: args must be null
         */
        assert args == null : "args must be null";
    }
}
~~~

~~~java
import java.util.Set;

import javax.annotation.processing.AbstractProcessor;
import javax.annotation.processing.ProcessingEnvironment;
import javax.annotation.processing.RoundEnvironment;
import javax.annotation.processing.SupportedAnnotationTypes;
import javax.annotation.processing.SupportedSourceVersion;
import javax.lang.model.SourceVersion;
import javax.lang.model.element.Element;
import javax.lang.model.element.ElementKind;
import javax.lang.model.element.TypeElement;
import javax.tools.Diagnostic;

import com.sun.source.util.Trees;
import com.sun.tools.javac.processing.JavacProcessingEnvironment;
import com.sun.tools.javac.tree.JCTree;
import com.sun.tools.javac.tree.TreeMaker;
import com.sun.tools.javac.tree.TreeTranslator;
import com.sun.tools.javac.tree.JCTree.JCAssert;
import com.sun.tools.javac.tree.JCTree.JCStatement;
import com.sun.tools.javac.util.Context;
import com.sun.tools.javac.util.List;
import com.sun.tools.javac.util.Names;


@SupportedSourceVersion(SourceVersion.RELEASE_8)
@SupportedAnnotationTypes("*")
public class ChangeAssertProcessor extends AbstractProcessor {

    private int tally;
    private Trees trees;
    private TreeMaker make;
    private Names names;

    /**
     * Initial java processiong env
     *
     * @param env
     */
    @Override
    public synchronized void init(ProcessingEnvironment env) {
        super.init(env);
        trees = Trees.instance(env);
        Context context = ((JavacProcessingEnvironment)
                env).getContext();
        make = TreeMaker.instance(context);
        names = Names.instance(context);
        tally = 0;
    }

    @Override
    public boolean process(Set<? extends TypeElement> annotations,
                           RoundEnvironment roundEnv) {
        if (!roundEnv.processingOver()) {
            Set<? extends Element> elements =
                    roundEnv.getRootElements();
            for (Element each : elements) {
                if (each.getKind() == ElementKind.CLASS) {
                    JCTree tree = (JCTree) trees.getTree(each);
                    TreeTranslator visitor = new Inliner();
                    tree.accept(visitor);
                }
            }
        } else {
            processingEnv.getMessager().printMessage(
                    Diagnostic.Kind.NOTE, tally + " assertions inlined.");
        }
        return false;
    }
    /**
     * Inliner only override visitAssert method, so only assert statement is modified.
     */
    private class Inliner extends TreeTranslator {

        @Override
        public void visitAssert(JCAssert tree) {
            super.visitAssert(tree);
            JCStatement newNode = makeIfThrowException(tree);
            result = newNode;
            tally++;
        }

        private JCTree.JCStatement makeIfThrowException(JCTree.JCAssert node) {
            // make: if (!(condition) throw new AssertionError(detail);
            List<JCTree.JCExpression> args = node.getDetail() == null
                    ? List.<JCTree.JCExpression>nil()
                    : List.of(node.detail);
            JCTree.JCExpression expr = make.NewClass(
                    null,
                    null,
                    make.Ident(names.fromString("AssertionError")),
                    args,
                    null);
            return make.If(
                    make.Unary(JCTree.Tag.NOT, node.cond),
                    make.Throw(expr),
                    null);
        }

    }
}

~~~
####1）编译processor：
javac -cp /Library/Java/JavaVirtualMachines/jdk1.8.0_162.jdk/Contents/Home/lib/tools.jar ChangeAssertProcessor.java 
####2）编译example：
javac -processor ChangeAssertProcessor AssertExample.java
####3）运行example：
java -disableassertions AssertExample

输出：

~~~java
Exception in thread "main" java.lang.AssertionError: args must be null
        at AssertExample.main(AssertExample.java:5)
~~~
####4）查看修改后的example：
```javac -processor ChangeAssertProcessor -printsource  -s ./org   AssertExample.java ```
生成的AssertExample.java放到org目录下，防止覆盖旧的代码。
~~~java
public class AssertExample {
    
    public AssertExample() {
        super();
    }
    
    public static void main(String[] args) {
        if (!(args == null)) throw new AssertionError("args must be null");
    }
    }

~~~

### 代码路径
<<< @/code-enhancer/AssertExample.java  
<<< @/code-enhancer/ChangeAssertProcessor.java