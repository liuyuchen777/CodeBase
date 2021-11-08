package codebase.basic.tool;

public class Print {

    private static boolean flag = true;

    public static void print(Object obj) {
        if (flag) {
            System.out.println(obj);
        }
    }
}
