package codebase.basic.stream;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Random;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class BasicStream {
    public static void main(String[] args) {
        System.out.println("------------codebase.basic.oop.Basic Stream-----------");
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

        System.out.println("-------------Use Reduce---------");
        /**
         * Stream.reduce()则是Stream的一个聚合方法，它可以把一个Stream的所有元素按照聚合函数聚合成一个结果。
         * The first element of reduce is the initial value
         *
         * Identity – an element that is the initial value of the reduction operation and the
         *            default result if the codebase.basic.stream is empty
         * Accumulator – a function that takes two parameters: a partial result of the reduction
         *               operation and the next element of the codebase.basic.stream
         * Combiner – a function used to combine the partial result of the reduction operation when
         *            the reduction is parallelized or when there's a mismatch between the types of
         *            the accumulator arguments and the types of the accumulator implementation
         */
        int sum = Stream.of(1, 2, 3, 4, 5, 6, 7, 8, 9).reduce(0, (acc, n) -> {
            acc += n;
            return acc;
        });
        // 0 => Identity, (acc, n) -> acc + n => Accumulator
        System.out.println("sum: " + sum);

        List<String> props = List.of("profile=native", "debug=true", "logging=warn", "interval=500");
        Map<String, String> map = props.stream()
                .map(kv -> {
                    String[] ss = kv.split("\\=", 2);
                    return Map.of(ss[0], ss[1]);
                })
                .reduce(new HashMap<String, String>(), (m, kv) -> {
                    m.putAll(kv);
                    return m;
                });
        map.forEach((k, v) -> {
            System.out.println(k + " = " + v);
        });


    }
}
