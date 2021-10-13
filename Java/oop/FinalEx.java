/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-19 09:35:02
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-19 10:04:01
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/FinalEx.java
 * @GitHub: https://github.com/liuyuchen777
 */

/*
    final could be used at variable, method and class
    final + varible: constant (static/global, or unchanged class variable)
    final + method/function: method can't be override
    final + class: class can't be extended
*/

import java.util.*;

class Value {
    int i;
    public Value(int i) { this.i = i; }
}

class WithFinals {
    /*
        1. final variable
    */
    // compile-time constants
    // 在编译时就确定了值
    private final int valueOne = 9;
    private static final int VALUE_TWO = 99;
    // typical public constant
    public static final int VALUE_THREE = 39;
    // cannot be compile-time constants
    // 因为涉及到初始化和函数调用，在运行时才能确定值
    private Value v1 = new Value(11);
    private final Value v2 = new Value(22);
    private static final Value VAL_3 = new Value(33);
    // array
    private final int[] a = {1, 2, 3, 4, 5, 6};
    /*
        2. final method
        1) lock method, any other extend class can't refine it, keep it invariant
        2) efficienncy (abandoned)
        * final and private
        private关键字隐含final的含义，因为其类外成员无法调用
    */
    public final void f() { System.out.println(" WithFinals.f() "); }
    private void g() { System.out.println(" WithFinals.g() "); }
}

class OverridingPrivate extends WithFinals {
    // public final void f() { System.out.println(" OverridingPrivate.f() ");} // error
    private void g() { System.out.println(" OverridingPrivate.f() "); }

}

final class FinalClass {
    int i = 8;
    int j = 8;
}

// error
// class TryToExtend extends FinalClass {
    
// }

class Book {
    boolean checkOut = false;
    
    Book(boolean checkOut) {
        this.checkOut = checkOut;
    }

    void checkIn() {
        checkOut = false;
    }

    protected void finalize() {
        // finalize is used to label object to be cleaned, but not necessary be cleaned
        if (checkOut)
            System.out.println("Error: checked out");
        // call base-class finalize function
        // super.finalize();
    }
}

public class FinalEx {
    public static void main(String[] args) {
        System.out.println("-------------final variable-----------");

        Book novel = new Book(true);
        // proper cleanup
        novel.checkIn();
        new Book(true);
        System.gc(); // call gc to clean garbage
    }
}
