package codebase.basic.generic;

import java.lang.reflect.InvocationTargetException;
import java.util.*;

class Frob {}

class Fnorkle {}

class Quark<Q> {}

class Particle<POSITION, MOMENTUM> {}

class HasF {
    public void f() {
        System.out.println("codebase.basic.generic.HasF.f()");
    }
}

class Manipulator<T> {
    private T obj;
    public Manipulator(T x) {
        obj = x;
    }
    // public void manipulate() { obj.f(); } // compile error: Object doesn't have f()
}

class ManipulatorRevise<T extends  HasF> {
    private T obj;
    public ManipulatorRevise(T x) {
        obj = x;
    }
    public void manipulate() { obj.f(); }
}

class GenericBase<T> {
    private T element;
    public void set(T arg) { arg = element; }
    public T get() { return element; }
}

class Derived1<T> extends GenericBase<T> {}

class Derived2 extends GenericBase {} // No Warning

public class GenericErasure {
    public static void main(String[] args)
            throws NoSuchMethodException, InvocationTargetException, IllegalAccessException {
        System.out.println("---------------Type Erasure--------------");
        Class c1 = new ArrayList<String>().getClass();
        Class c2 = new ArrayList<Integer>().getClass();
        System.out.println(c1 == c2);

        ArrayList<Integer> list = new ArrayList<Integer>();
        list.add(1);  //这样调用 add 方法只能存储整形，因为泛型类型的实例为 Integer
        list.getClass().getMethod("add", Object.class).invoke(list, "A String");
        for (int i = 0; i < list.size(); i++) {
            System.out.println(list.get(i));
        }

        System.out.println("---------------Get Generic Type Information--------------");
        List<Frob> alist = new ArrayList<>();
        Map<Frob, Fnorkle> map = new HashMap<>();
        Quark<Fnorkle> quark = new Quark<>();
        Particle<Long, Double> p = new Particle<>();

        System.out.println(Arrays.toString(alist.getClass().getTypeParameters()));
        System.out.println(Arrays.toString(map.getClass().getTypeParameters()));
        System.out.println(Arrays.toString(quark.getClass().getTypeParameters()));
        System.out.println(Arrays.toString(p.getClass().getTypeParameters()));

        System.out.println("---------------Generaic Bound--------------");
        Derived2 d2 = new Derived2();
        Object obj = d2.get();
        d2.set(obj);
    }
}
