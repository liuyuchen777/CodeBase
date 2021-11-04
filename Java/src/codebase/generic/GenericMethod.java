package codebase.generic;

import java.util.*;

class New {
    public static <K, V> Map<K, V> map() {
        return new HashMap<>();
    }

    public static <T> List<T> list() {
        return new ArrayList<>();
    }

    // public static <T> List<> list() {    // compile error: can't auto deduce codebase.type
    //     return new ArrayList<>();
    // }

    public static <T> LinkedList<T> lList() {
        return new LinkedList<>();
    }

    public static <T> Set<T> set() {
        return new HashSet<>();
    }

    public static <T> Queue<T> queue() {
        return new LinkedList<>();
    }
}

public class GenericMethod {
    public static <E> void printArray(E[] inputArray) {
        for (E element : inputArray) {
            System.out.print(element + " ");
        }
        System.out.println();
    }

    public static  <T> void display(T... arrays) {
        for (T item : arrays) {
            System.out.print(item + " ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        System.out.println("--------------Generic Method------------");
        Integer[] intArray = {1, 2, 3, 4, 5};
        Double[] doubleArray = {1., 2., 3., 4., 5.};
        Character[] charArray = {'H', 'E', 'L', 'L', 'O'};

        printArray(intArray);
        printArray(doubleArray);
        printArray(charArray);

        System.out.println("--------------Type Deduce------------");
        Map<String, List<String>> sls = New.map();
        List<String> ls = New.list();
        LinkedList<String> lls = New.lList();
        Set<String> ss = New.set();
        Queue<String> qs = New.queue();

        System.out.println("--------------Generic and Variable Element------------");
        display("Hello", "World", "You", "Little", "Guy");
    }
}
