/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-23 23:29:39
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-23 23:39:26
 * @Description: 
 * @FilePath: /CodeBase/Java/Type/src/LoadClass.java
 * @GitHub: https://github.com/liuyuchen777
 */

/**
 * 每一个类都有一个Class对象，存储Class相关信息
 * 保存在.class文件中
 */

class Candy {
    // static line will be execute when class is load
    static { System.out.println("Loading Candy!"); }
}

class Gum {
    static { System.out.println("Loading Gum!"); }
}

class Cookie {
    static { System.out.println("Loading Cookie!"); }
}

public class LoadClass {
    public static void main(String[] args) {
        System.out.println("------------------Java Load Mechansim---------------");
        System.out.println("Inide main!");
        new Candy();
        System.out.println("After creating Candy!");
        try {
            Class.forName("Gum");
        } catch(ClassNotFoundException e) {
            System.out.println("Couldn't find Gum!");
        }
        System.out.println("After Class.forName(\"Gum\")!");
        new Cookie();
        System.out.println("After creating Cookie!");
    }
}
