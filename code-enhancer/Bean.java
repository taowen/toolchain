/**
 * @author liqingsong on 2019/3/6
 */
public class Bean implements BeanService{
    private int  f;

    public int getF() {
        return f;
    }

    public void setF(int f) {
        this.f = f;
    }

    public static void hello() {
        System.out.println("Hello, world");
    }

    public static void main(String[] args) {
        Bean.hello();
    }

    @Override
    public void service1() {
        System.out.println("echo from service1");
    }
}
