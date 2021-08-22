/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-24 21:46:33
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-24 21:50:20
 * @Description: 
 * @FilePath: /CodeBase/Java/Generics/src/GenericBasic.java
 * @GitHub: https://github.com/liuyuchen777
 */
public class GenericBasic {
    // generic method
    public static <E> void printArray(E[] inputArray) {
        for (E element : inputArray) {
            System.out.print(element + " ");
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
    }
}
