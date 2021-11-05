package codebase.basic.type;

import java.lang.reflect.*;

class Apple {
    private int price = 0;
    public Apple() {
        System.out.println("I am in codebase.basic.type.Apple Constructor!");
    }
    public void setPrice(int price) {
        this.price = price;
    }
    public String toString() {
        return "The price is " + this.price;
    }
}

public class ReflectBasic {
    public static void main(String[] args) {
        System.out.println("-------------General Way to Use Class----------");
        Apple apple = new Apple();
        apple.setPrice(10);
        System.out.println(apple);
        System.out.println("------------Use Reflection---------");
        Object object = null;
        try {
            // 1. 获取反射中的Class对象
            // 1.1. 使用class.forName(...)静态方法
            Class clz = Class.forName("codebase.basic.container.AppleItem");
            // 1.2. 使用.class方法，只适合在编译器就知道的Class
            /*
            Class clz = String.cls;
            */
            // 1.3. 使用类对象的getClass()方法
            /*
            String str = new String("Hello!");
            Class clz = str.getClass()
            */

            // 2. 通过反射创建类对象
            // 2.1. 通过Constructor对象的newInstance()方法
            Constructor appleConstructor = clz.getConstructor();
            object = appleConstructor.newInstance();
            // 2.2. 通过Class对象的newInstance()方法，这个方法在Java9之后已经被淘汰了。通过Class对象
            // 只能使用默认的无参构造方法
            /*
            codebase.basic.type.Apple apple = (codebase.basic.type.Apple)clz.newInstance()
            */

            // 3. 通过反射获取类属性，方法，构造器
            // 3.1. 获取类方法
            Method method = clz.getMethod("setPrice", int.class);
            method.invoke(object, 4);
            // 3.2. 获取类属性
            /*
            // 1) Class.getFileds()无法获取私有属性
            Field[] fields = clz.getFields();
            for (Field field : fields) {
                System.out.println(field.getName());
            }
            // 2) Class.getDeclaredFields()可以获取包括私有属性在内的多有属性
            Class clz = codebase.basic.type.Apple.class;
            Field[] fields = clz.getDeclaredFields();
            for (Field field : fields) {
                System.out.println(field.getName());
            }
            */
        } catch(Exception e) {
            System.out.println("Reflect Error! " + e);
        }
        System.out.println(object);
    }
}
