import io.reactivex.Flowable;
import io.reactivex.schedulers.Schedulers;

public class ConcurrentRpc {

    public static <T> T latency(T v) {
        try {
            Thread.sleep(200);
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        return v;
    }

    public static Flowable<Integer> rpcA() {
        return Flowable.range(1, 10)
                .flatMap(v -> Flowable.just(v)
                        .subscribeOn(Schedulers.computation())
                        .map(ConcurrentRpc::latency));
    }

    public static Flowable<String> rpcB() {
        String[] arr = {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"};
        return Flowable.fromArray(arr)
                .flatMap(v -> Flowable.just(v)
                        .subscribeOn(Schedulers.computation())
                        .map(ConcurrentRpc::latency));
    }

    public static Flowable<String> rpcC(Flowable<Integer> aStrm, Flowable<String> bStrm) {
        return aStrm.zipWith(bStrm, (a,b) -> String.format("%s_%d", b, a));
    }

    public static void main(String[] args) throws Exception{
        long start = System.currentTimeMillis();
        rpcC(rpcA(), rpcB()).blockingSubscribe(a -> System.out.println(a));
        System.out.println(System.currentTimeMillis() - start);

        Thread.sleep(10000);
    }
}

