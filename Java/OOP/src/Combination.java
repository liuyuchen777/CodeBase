/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-19 08:56:05
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-19 09:09:12
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/Combination.java
 * @GitHub: https://github.com/liuyuchen777
 */
class WaterSource {
    private String s;
    WaterSource() {
        System.out.println("WaterSource()");
        s = "Constructed";
    }
    public String toString() {
        /*
            System.out.print()系列函数会直接调用String.valueOf()，String.valueOf()又会调用类中的toString()
        */
        return s;
    }
}

class SprinklerSystem {
    private String valve1, valve2;
    private WaterSource source = new WaterSource(); // 组合：在一个类中使用另一个类

    public SprinklerSystem(String valve1, String valve2) {
        this.valve1 = valve1;
        this.valve2 = valve2;
    }

    public String toString() {
        // 延迟初始化，在需要该成员的时候才进行初始化，减少程序运行时的负担
        if (valve1 == null) {
            valve1 = "lyc";
        }
        if (valve2 == null) {
            valve2 = "zzh";
        }

        return 
            "valve1 = " + valve1 + " " +
            "valve2 = " + valve2 + "\n" + 
            "source = " + source;
    }
}

public class Combination {
    public static void main(String[] args) {
        WaterSource ws = new WaterSource();
        System.out.println(ws);
    }
}
