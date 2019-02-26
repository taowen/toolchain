import io.reactivex.*;
import io.reactivex.schedulers.*;

public class ParallelComputation {
    public static int task(int v) {
        try {
            Thread.sleep(200);
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        return v * v;
    }
    public static void main(String[] args) {
        long start = System.currentTimeMillis();
        Flowable.range(1, 10)
                .map(ParallelComputation::task)
                .blockingSubscribe(v -> System.out.printf("%d ", v));
        System.out.printf("sequenial mode time: %d\n", System.currentTimeMillis() - start);

        start = System.currentTimeMillis();
        Flowable.range(1, 10)
                .flatMap(v ->
                        Flowable.just(v)
                                .subscribeOn(Schedulers.computation())
                                .map(ParallelComputation::task))
                .blockingSubscribe(v -> System.out.printf("%d ", v));
        System.out.printf("parallel mode time: %d\n", System.currentTimeMillis() - start);
    }
}

