
import java.util.Arrays;
import java.util.List;

public class ParallelStream {
    public static void main(String[] args) {
        System.out.println("----------Parallel Stream---------");
        List<String> strings = Arrays.asList("abc", "", "bc", "efg", "abcd","", "jkl");
        // 获取空字符串的数量
        long count = strings.parallelStream()
                        .filter(String::isEmpty)
                        .count();
        System.out.println("count: " + count);

        System.out.println("--------Parallel Reduce--------");
        
    }
}
