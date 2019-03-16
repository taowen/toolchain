import org.objectweb.asm.*;

import java.io.File;
import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.nio.file.Files;
import java.nio.file.Paths;

import static org.objectweb.asm.Opcodes.*;

/**
 * @author liqingsong on 2019/3/9
 */
public class AsmModifyMethod {

    public static class LogVisitor extends ClassVisitor {

        public LogVisitor(int api, ClassVisitor classVisitor) {
            super(api, classVisitor);
        }

        @Override
        public MethodVisitor visitMethod(int access, String name, String descriptor, String
                signature, String[] exceptions) {
            if (!(name.equalsIgnoreCase("<init>") || name.equalsIgnoreCase("main"))) {
                MethodVisitor methodVisitor = cv.visitMethod(access, name, descriptor, signature,
                        exceptions);
                return new LogMethodVisitor(this.api, methodVisitor);
            }
            return super.visitMethod(access, name, descriptor, signature, exceptions);
        }
    }

    public static class LogMethodVisitor extends MethodVisitor {

        public LogMethodVisitor(int api, MethodVisitor methodVisitor) {
            super(api, methodVisitor);
        }

        @Override
        public void visitCode() {
            mv.visitFieldInsn(Opcodes.GETSTATIC, "java/lang/System", "out",
                    "Ljava/io/PrintStream;");

            mv.visitLdcInsn("method start");
            mv.visitMethodInsn(INVOKEVIRTUAL, "java/io/PrintStream", "println", "" +
                    "(Ljava/lang/String;)V", false);
            super.visitCode();
        }

        @Override
        public void visitInsn(int opcode) {
            if (opcode == ARETURN || opcode == RETURN) {
                mv.visitFieldInsn(Opcodes.GETSTATIC, "java/lang/System", "out",
                        "Ljava/io/PrintStream;");
                mv.visitLdcInsn("method end");
                mv.visitMethodInsn(INVOKEVIRTUAL, "java/io/PrintStream", "println", "" +
                        "(Ljava/lang/String;)V", false);
            }
            super.visitInsn(opcode);
        }

        @Override
        public void visitEnd() {
            mv.visitMaxs(6, 6);
            super.visitEnd();
        }

    }

    public static class CustomizeClassLoader extends ClassLoader {
        public Class<?> defineClass(String name, byte[] b) {
            return defineClass(name, b, 0, b.length);
        }
    }

    public static void main(String[] args) throws IOException, IllegalAccessException,
            InstantiationException, NoSuchMethodException, InvocationTargetException {

        /**
         * Use asm operate bytecode
         */
        String clzName = Bean.class.getName().replace(",", "/");
        ClassReader classReader = new ClassReader(clzName);
        ClassWriter classWriter = new ClassWriter(0);
        LogVisitor logVisitor = new LogVisitor(ASM7, classWriter);
        classReader.accept(logVisitor, ClassReader.SKIP_DEBUG);
        byte[] classBytes = classWriter.toByteArray();

        /**
         * Dynamically load modified class
         */
        CustomizeClassLoader customizeClassLoader = new CustomizeClassLoader();

        Class<?> clazz = customizeClassLoader.defineClass("Bean", classBytes);
        /**
         * The new class has the same name as the old one.
         */
        System.out.println(clazz.getName());

        /**
         * Invoke static method without instance. Modified method will be called.
         */
        Method staticHelloMethod = clazz.getMethod("hello");
        staticHelloMethod.invoke(null);

        /**
         * Bean can't cast to Bean for different classloader
         */
        System.out.println(Bean.class.getClassLoader());
        System.out.println(clazz.getClassLoader());

        /**
         * This Bean is still old one, and hello still is the old method.
         */
        Bean.hello();

        /**
         * Interface helps resolve classloader issue.
         */
        BeanService beanService = (BeanService) clazz.newInstance();
        beanService.service1();


        /**
         * Another way: create new static class file
         */

        File file = new File("./Bean.class");
        if (!file.exists()) {
            boolean result = file.createNewFile();
            System.out.println(result);
        }
        Files.write(Paths.get("./Bean.class"), classBytes);
    }
}
