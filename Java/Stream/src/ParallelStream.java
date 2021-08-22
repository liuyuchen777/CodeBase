/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-29 21:01:40
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-30 11:29:15
 * @Description: 
 * @FilePath: /CodeBase/Java/Stream/src/ParallelStream.java
 * @GitHub: https://github.com/liuyuchen777
 */
import java.util.Arrays;
import java.util.List;

public class ParallelStream {
    public static void main(String[] args) {
        System.out.println("----------Parallel Stream---------");
        List<String> strings = Arrays.asList("abc", "", "bc", "efg", "abcd","", "jkl");
        // 获取空字符串的数量
        long count = strings.parallelStream()
                        .filter(string -> string.isEmpty())
                        .count();
        System.out.println("count: " + count);

        System.out.println("--------Parallel Reduce--------");
        
    }
}
