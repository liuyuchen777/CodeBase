import java.util.Arrays;

/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-18 10:21:29
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-18 14:21:28
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/Initialize.java
 * @GitHub: https://github.com/liuyuchen777
 */

public class Initialize {
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

    Initialize() {
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
        Initialize it = new Initialize();
        System.out.println(Arrays.toString(it.aInt));
        System.out.println(Arrays.toString(it.fInt));
        System.out.println(Arrays.toString(it.bInt));
        System.out.println(it.aInteger[0]);
        System.out.println(Arrays.toString(it.aInteger2));

        Initialize.printArray(new Object[]{1, "lyc", 10.2});
        Initialize.printArrayNew(1, 10.2, "lyc");
    }
}
