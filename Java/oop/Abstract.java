/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-13 18:29:37
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-20 11:41:04
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/Abstract.java
 * @GitHub: https://github.com/liuyuchen777
 */

/**
 * 抽象类可以被视作一种纯接口
 * 如果一个类包含一个或者多个抽象方法，该类必须被限定为抽象的
 * 一个类如果从抽象类继承，他必须实现抽象基类中的所有函数
 */

abstract class Animals {
    // java don't support extend from multiple class
    protected String name;
    protected int id;
    private static int idNow = 1;

    public Animals(String myName) {
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

    /*
    abstract method must be overload
    and class contains abstract method must be abstract class
    */
    public abstract void dance();
}

class Cat extends Animals {
    public Cat(String myName) {
        super(myName);
    }

    public void dance() {
        System.out.println("cat can't dance");
    }
}

class People extends Animals {
    public People(String myName) {
        super(myName);
    }

    public void dance() {
        System.out.println("ID=" + super.id + ", " + super.name + " is dancing!");
    }
}

public class Abstract {
    public void basicOperate() {
        Cat ct = new Cat("Zhu Junyi");
        ct.dance();
        People pl = new People("Liu Yuchen");
        pl.dance();
    }

    public static void main(String[] args) {
        Abstract ab = new Abstract();
        ab.basicOperate();
    }
}