import java.util.Arrays;
import java.util.List;
import java.util.Random;
import java.util.stream.Collectors;

/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-29 20:47:36
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-29 21:05:00
 * @Description: 
 * @FilePath: /CodeBase/Java/Stream/src/BasicStream.java
 * @GitHub: https://github.com/liuyuchen777
 */

public class BasicStream {
    public static void main(String[] args) {
        System.out.println("------------Basic Stream-----------");
        List<String> strings = Arrays.asList("abc", "", "efg", 
            "abcd", "", "jkl");
        List<String> filtered = strings.stream()
                                    .filter(string -> !string.isEmpty())
                                    .collect(Collectors.toList());
        System.out.println("origin: " + strings);
        System.out.println("filtered: " + filtered);

        System.out.println("-----------Example 2----------");
        List<Integer> numbers = Arrays.asList(3, 2, 2, 3, 7, 3, 5);
        List<Integer> squaresLists = numbers.stream()
                                    .map(i -> i * i)
                                    .distinct()
                                    .collect(Collectors.toList());
        System.out.println("origin: " + numbers);
        System.out.println("squares: " + squaresLists);

        System.out.println("----------forEach in Stream--------");
        Random random = new Random();
        random.ints()
            .limit(10)
            .map(num -> num % 10)
            .forEach(System.out::println);

        System.out.println("---------Stream sort--------");
        random.ints()
            .limit(10)
            .map(num -> num % 10)
            .sorted()
            .forEach(System.out::println);
    }
}
