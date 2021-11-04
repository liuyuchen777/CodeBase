package codebase.type;

/**
 * 每一个类都有一个Class对象，存储Class相关信息
 * 保存在.class文件中
 */

class Candy {
    // static line will be execute when class is load
    static { System.out.println("Loading codebase.type.Candy!"); }
}

class Gum {
    static { System.out.println("Loading codebase.type.Gum!"); }
}

class Cookie {
    static { System.out.println("Loading codebase.type.Cookie!"); }
}

public class LoadClass {
    public static void main(String[] args) {
        System.out.println("------------------Java Load Mechansim---------------");
        System.out.println("Inide main!");
        new Candy();
        System.out.println("After creating codebase.type.Candy!");
        try {
            Class.forName("codebase.type.Gum");
        } catch(ClassNotFoundException e) {
            System.out.println("Couldn't find codebase.type.Gum!");
        }
        System.out.println("After Class.forName(\"codebase.type.Gum\")!");
        new Cookie();
        System.out.println("After creating codebase.type.Cookie!");
    }
}
