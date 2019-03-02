import java.math.BigDecimal;
import java.util.function.*;

public class TestFunction {
    public static void main(String[] args) {
        /**
         * test function and functions composition.
         */
        Function<BigDecimal, Long> multInt1 = TestFunction::multIntMethod;
        Function<Long, Long> multInt2 = returnInnerClassWithState(classStateLong);
        Function<Integer, Long> multInt = returnLambda();
        Function<Long, BigDecimal> multLong = (input) -> BigDecimal.valueOf(input * Integer
                .MAX_VALUE);
        System.out.println(multLong.compose(multInt).andThen(multInt1).andThen(multInt2).apply
                (100));
        multLong.compose(multInt).apply(1);
        Predicate<? super Number> ps = (num) -> num.intValue() >= 2;
        System.out.println("Predicate<? super Number>: " + ps.test(1));
        System.out.println("Predicate<? super Number>: " + ps.test(2L));
        Predicate<? extends String> pe = (str) -> str.startsWith("a");
	//System.out.println("Predicate<? extends Number>: " + pe.test(null));
	//compile failed
	//System.out.println("Predicate<? extends Number>: " + pe.test("abc");

        Function<String, Long> innerFun = (intVar) -> new Long(intVar + 1);
        Function<Number, String> outterFun = (longVar) -> Long.toString(longVar.longValue() + 1);
        String resStr = outterFun.compose(innerFun).apply("1");
        System.out.println("resStr: " + resStr);

        TestFunction mainTest = new TestFunction();
        Long res = multInt2.compose(mainTest.objectFun()).andThen(mainTest::zeroForEver).apply
                ("123");
        System.out.println("res:" + res);

        BiFunction<Long, Long, Long> sumLong = (f1, f2) -> f1 + f2;
        System.out.println(sumLong.apply(4L, 5L));

        Consumer<Integer> integerConsumer = (i) -> System.out.println("integerConsumer:" + i);
        Consumer<Number> numberConsumer = (num) -> System.out.println("numberConsumer:" + num
                .intValue());
        Consumer<Object> objConsumer = (obj) -> System.out.println("objConsumer:" + obj);
        integerConsumer.andThen(numberConsumer).andThen(objConsumer).accept(222);

        Consumer<Integer> integerConsumer2 = (i) -> System.out.println("integerConsumer2:" + i);
        integerConsumer.andThen(integerConsumer2).accept(333);

        Supplier<? extends Number> supplier = () -> new Integer(1);
        System.out.println("Supplier<? extends Number>: " + supplier.get() + ", class: " +
                supplier.get().getClass().getName());
        Supplier<? super Number> supplier1 = () -> new Long(2);
        System.out.println("Supplier<? super Number>: " + supplier1.get() + ", class: " +
                supplier1.get().getClass().getName());
        supplier = () -> new Long(1);
        System.out.println("Supplier<? extends Number>: " + supplier.get() + ", class: " +
                supplier.get().getClass().getName());

        UnaryOperator<String> unaryOperator = (str) -> str + "tttt";
        System.out.println(unaryOperator.apply("abc"));

        Function<String, String> biOperator = Function.identity();
        System.out.println(biOperator.apply("abc"));
        UnaryOperator.identity().apply("abc");
        Function.identity().apply("abc");

    }

    private static Long classStateLong = 2L;

    public static Long multIntMethod(BigDecimal input) {
        final Integer factor = -1;
        Long res = Long.valueOf(input.longValue() * factor * classStateLong);
        System.out.println("multIntMethod, input:" + input.longValue() + ";res:" + res);
        return res;
    }

    private static Function<Integer, Long> returnInnerClass() {
        return new Function<Integer, Long>() {
            @Override
            public Long apply(Integer integer) {
                return Long.valueOf(integer * -1);
            }
        };
    }

    private static Function<Long, Long> returnInnerClassWithState(Long state) {
        return new Function<Long, Long>() {
            @Override
            public Long apply(Long integer) {
                return Long.valueOf(integer * -1 * state);
            }
        };
    }

    private static Function<Integer, Long> returnLambda() {
        return (i) -> {
            Long res = Long.valueOf(i * -1);
            System.out.println("returnLambda, i:" + i + "; res:" + res);
            return res;
        };
    }

    public Function<String, Long> objectFun() {
        return new Function<String, Long>() {
            @Override
            public Long apply(String s) {
                return Long.parseLong(s);
            }
        };
    }

    public Long zeroForEver(Long input) {
        return input * 0;
    }
}
