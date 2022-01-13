package codebase.basic.oop;

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
    public final void f() { System.out.println(" codebase.basic.oop.WithFinals.f() "); }
    private void g() { System.out.println(" codebase.basic.oop.WithFinals.g() "); }
}

class OverridingPrivate extends WithFinals {
    // public final void f() { System.out.println(" codebase.basic.oop.OverridingPrivate.f() ");} // error
    private void g() { System.out.println(" codebase.basic.oop.OverridingPrivate.f() "); }

}

final class FinalClass {
    int i = 8;
    int j = 8;
}

// error
// class TryToExtend extends codebase.basic.oop.FinalClass {
    
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