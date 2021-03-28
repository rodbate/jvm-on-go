/**
 * @author rodbate
 * @since 2020/12/20
 */
public class MethodInvocationSample {

    public static void main(String[] args) {
        final String line = "line1\nline2\nline3";
        MethodInvocationSample sample = new MethodInvocationSample();
        for (String s : line.split("\n")) {
            sample.print(s);
        }
    }

    private void print(String s) {
        System.out.println("print: " + s);
    }
}