import java.util.Arrays;
/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-18 10:21:29
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-20 11:30:27
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/InitializeConstructor.java
 * @GitHub: https://github.com/liuyuchen777
 */

/**
 * An example for Java constructor use order
 * 构造函数的调用顺序:
 * 1. 基类构造器，反复递归，直到最底层的导出类（基类）
 * 2. 按声明顺序调用成员的初始化方法
 * 3. 导出类构造器的主体
*/
class Bread {
    Bread() { System.out.println("Bread()"); }
}

class Cheese {
    Cheese() { System.out.println("Cheese()"); }
}

class Meal {
    Meal() { System.out.println("Meal()"); }
}

class Lunch extends Meal {
    Lunch() { System.out.println("Lunch()"); }
}

class PortalLunch extends Lunch {
    PortalLunch() { System.out.println("PortalLunch()");}
}

class SandWich extends PortalLunch {
    private Bread b = new Bread();
    private Cheese c = new Cheese();
    public SandWich() { System.out.println("Sandwich()"); }

    public static void main(String[] args) {
        SandWich s = new SandWich();
    }
}

public class InitializeConstructor {
    /*
        you can direct initialize member value in Java
        which is not allowed in C++
    */
    boolean bool = true;
    char ch = 'x';
    byte b = 47;
    short s = 0xff;
    int i = 999;
    long lng = 1;
    float f = 3.14f;
    double d = 3.1415926;
    // other ways to initialize
    int intValue = f();

    int f() { 
        return 11; 
    }

    // constructor initialzie
    int intValue1;
    int intValue2;
    int intValue3;

    InitializeConstructor() {
        /*
            when use constructor to initialize, member variable is first setted default value of type
            and then get value in constructor
        */
        System.out.println("intValue1 = " + intValue1);
        intValue1 = 1;
        intValue2 = 2;
        intValue3 = 3;
        System.out.println("intValue1 = " + intValue1);
    }

    // array initialize
    public int[] aInt = new int[10];
    public float[] fInt = new float[10];
    public boolean[] bInt = new boolean[10];
    /*
        Java中数组有基本类型的数组和引用数组之分
        基本类型数组可以直接使用，引用类型的则需要先初始化然后再使用
        初始化之前值为null
        Java数组全部存储在堆上
    */
    public Integer[] aInteger = new Integer[10];
    public Integer[] aInteger2 = new Integer[] {
        Integer.valueOf(1), // new Integer(value) is deprecated
        2,
        3,
    };

    // static initialize
    static int staticValue;
    static {
        staticValue = 47;
    }

    // Java可变参数
    public static void printArray(Object[] args) {
        for (Object obj : args) {
            System.out.print(obj + " ");
        }
        System.out.println();
    }
    // new approach
    public static void printArrayNew(Object... args) {
        for (Object obj : args) {
            System.out.print(obj + " ");
        }
        System.out.println();
    }
    
    public static void main(String[] args) {
        InitializeConstructor it = new InitializeConstructor();
        System.out.println(Arrays.toString(it.aInt));
        System.out.println(Arrays.toString(it.fInt));
        System.out.println(Arrays.toString(it.bInt));
        System.out.println(it.aInteger[0]);
        System.out.println(Arrays.toString(it.aInteger2));

        InitializeConstructor.printArray(new Object[]{1, "lyc", 10.2});
        InitializeConstructor.printArrayNew(1, 10.2, "lyc");
    }
}
