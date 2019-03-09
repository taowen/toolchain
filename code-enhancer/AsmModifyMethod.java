import org.objectweb.asm.*;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

import static org.objectweb.asm.Opcodes.*;

/**
 * @author liqingsong on 2019/3/9
 * @project javalang-common
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
            mv.visitLdcInsn("hello: method start");
            mv.visitMethodInsn(INVOKEVIRTUAL, "java/io/PrintStream", "println", "" +
                    "(Ljava/lang/String;)V", false);
            super.visitCode();
        }

        @Override
        public void visitInsn(int opcode) {
            if (opcode == ARETURN || opcode == RETURN) {
                mv.visitFieldInsn(Opcodes.GETSTATIC, "java/lang/System", "out",
                        "Ljava/io/PrintStream;");
                mv.visitLdcInsn("hello: method end");
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


    public static void main(String[] args) throws IOException {

        String clzName = Bean.class.getName().replace(",", "/");
        ClassReader classReader = new ClassReader(clzName);
        ClassWriter classWriter = new ClassWriter(0);
        LogVisitor logVisitor = new LogVisitor(ASM7, classWriter);
        classReader.accept(logVisitor, ClassReader.SKIP_DEBUG);
        byte[] classBytes = classWriter.toByteArray();

        File file = new File("./Bean.class");
        if (!file.exists()) {
            boolean result = file.createNewFile();
            System.out.println(result);
        }
        Files.write(Paths.get("./Bean.class"), classBytes);
    }
}
