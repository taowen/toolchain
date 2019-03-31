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

/**
 * @author liqingsong on 2019/3/31
 * @project javalang-common
 */
@SupportedSourceVersion(SourceVersion.RELEASE_8)
@SupportedAnnotationTypes("*")
public class ChangeAssertProcessor extends AbstractProcessor {

    private int tally;
    private Trees trees;
    private TreeMaker make;
//    private Name.Table names;

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

        /**
         * Convert assertion to exception
         *
         * @param node
         * @return
         */
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
