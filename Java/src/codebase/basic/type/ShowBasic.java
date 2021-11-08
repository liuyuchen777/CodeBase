package codebase.basic.type;

import java.lang.reflect.*;
import java.util.regex.*;
import static codebase.basic.tool.Print.*;

public class ShowBasic {
    private static String usage = """
        usage:
        ShowMethods qualified class name
        To show all methods in class or:
        ShowMethods qualified class name word
        To search for methods involving 'word'
    """;

    private static Pattern p = Pattern.compile("\\w+");

    public static void main(String[] args) {
        if (args.length < 1) {
            print(usage);
            System.exit(0);
        }
        int lines = 0;
        try {
            Class<?> c = Class.forName(args[0]);
            Method[] methods = c.getMethods();
            Constructor[] constructors = c.getConstructors();
            if (args.length == 1) {
                for (Method method : methods) {
                    print(p.matcher(method.toString()).replaceAll(""));
                }
                for (Constructor constructor : constructors) {
                    print(p.matcher(constructor.toString()).replaceAll(""));
                }
                lines = methods.length + constructors.length;
            } else {
                for (Method method : methods) {
                    if (method.toString().indexOf(args[1]) != -1) {
                        print(p.matcher(method.toString()).replaceAll(""));
                    }
                    lines++;
                }
                for (Constructor constructor : constructors) {
                    if (constructor.toString().indexOf(args[1]) != -1) {
                        print(p.matcher(constructor.toString()).replaceAll(""));
                    }
                    lines++;
                }
            }
        } catch (ClassNotFoundException e) {
            print("No such class: " + e);
        }
    }
}
