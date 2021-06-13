/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-13 18:29:37
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-13 18:44:01
 * @Description: 
 * @FilePath: \OOP\src\Abstract.java
 * @GitHub: https://github.com/liuyuchen777
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