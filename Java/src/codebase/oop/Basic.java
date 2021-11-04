package codebase.oop;

class Animal {
    // java don't support extend from multiple class
    private String name;
    private int id;
    private static int idNow = 1;

    public Animal(String myName) {
        name = myName;
        id = idNow;
        idNow++;
    }

    public void eat() {
        System.out.println("ID=" + id + ", " + name + " is eating");
    }

    public void sleep() {
        System.out.println("ID=" + id + ", " + name + " is sleeping");
    }
}

class Penguin extends Animal {
    /*
    子类是不继承父类的构造器（构造方法或者构造函数）的，它只是调用（隐式或显式）。
    如果父类的构造器带有参数，则必须在子类的构造器中显式地通过 super 关键字调用父类的构造器并配以适当的参数列表。
    如果父类构造器没有参数，则在子类的构造器中不需要使用 super 关键字调用父类构造器，系统会自动调用父类的无参构造器。
    */
    public Penguin(String myName) {
        super(myName);
    }
}

class Mouse extends Animal {
    public Mouse(String myName) {
        super(myName);
    }
}

public class Basic {
    public void basicExtend() {
        System.out.println("------------codebase.oop.Extend----------");
        Penguin pg = new Penguin("Geroge");
        pg.eat();
        pg.sleep();
        Mouse ms = new Mouse("Kawaii");
        ms.eat();
        ms.sleep();
    }

    public static void main(String[] args) {
        Basic bt = new Basic();
        bt.basicExtend();
    }
}