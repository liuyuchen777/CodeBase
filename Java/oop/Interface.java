/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-20 11:33:50
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-20 12:25:10
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/Interface.java
 * @GitHub: https://github.com/liuyuchen777
 */

/**
 * 一个接口中的所有方法都是抽象的
 * 成员变量是不可变且静态的编译阶的常量
 */

interface Instrument {
    // compile-time constant
    int VALUE = 5; // static & final
    // cammot have method definitions
    void play(String n);
    void adjust();
    // toString() is part of Object, no need to write in Interface
}

class Wind implements Instrument {
    public void play(String n) {
        System.out.println(this + ".play() " + n);
    }
    public String toString() { return "wind"; }
    public void adjust() { System.out.println(this + ".adjust()"); }
}

class Brass extends Wind {
    public String toString() { return "Brass"; }
}

class WoodWind extends Wind {
    public String toString() { return "WoodWind"; }
}

class Music {
    public static void tune (Instrument i) {
        i.play("Down to the Hello");
    }

    public static void tuneAll (Instrument[] e) {
        for (Instrument i : e) {
            tune(i);
        }
    }
}

/** 
 * multiple extend in Java:
 * Java don't directly support multiple extend
 * use Interface to realize
 */
interface CanFight {
    void fight();
}

interface CanFly {
    void fly();
}

interface CanSwim {
    void swim();
}

class ActionCharacter {
    public void fight() { System.out.println("ActionCharacter.fight()"); }
}

class Hero extends ActionCharacter
    implements CanFight, CanSwim, CanFly {
    public void swim() { System.out.println("Hero.swim()"); }
    public void fly() { System.out.println("Heo.fly()"); }
}

/**
 * interface extend:
 * 借口可以多重继承，见interface Vampire
 */

interface Monster {
    void menace();
}

interface DangerousMonster extends Monster {
    void destroy();
}

class DragonZilla implements DangerousMonster {
    public void menace() {}
    public void destroy() {}
}

interface Lethal {
    void kill();
}

interface Vampire extends DangerousMonster, Lethal {
    void drinkBoold();
}

class VeryBadVampire implements Vampire {
    public void menace() {}
    public void destroy() {}
    public void kill() {}
    public void drinkBoold() {}
}

// 组合借口时命名冲突，不能有相同输出类型但不同返回值的两个方法

public class Interface {
    public static void t(CanFight x) { x.fight(); }
    public static void u(CanSwim x) { x.swim();}
    public static void v(CanFly x) { x.fly(); }
    public static void w(ActionCharacter x) { x.fight(); }
    public static void main(String[] args) {
        System.out.println("----------Implement Interface--------");
        Instrument[] orchestra = {
            new Wind(),
            new Brass(),
            new WoodWind(),
        };
        Music.tuneAll(orchestra);

        System.out.println("-----------Multiple Extend------------");
        Hero h = new Hero();
        t(h);
        u(h);
        v(h);
        w(h);

        System.out.println("---------Extend Interface-------");
    }
}