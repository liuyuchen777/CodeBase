/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-19 09:09:23
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-19 09:22:11
 * @Description: standard extend example
 * @FilePath: /CodeBase/Java/OOP/src/Extend.java
 * @GitHub: https://github.com/liuyuchen777
 */

class Cleanser {
    private String s = "Cleanser ";
    public void append(String a) { s += a; }
    public void dilute() { append(" dilute() "); }
    public void apply() { append(" apply() "); }
    public void scrub() { append(" scrub() "); }
    public String toString() { return s; }
}

class Detergent extends Cleanser {
    public void scrub() {
        // cover scrub() in Cleanser
        append(" Detergent.srub() ");
        super.scrub();
    }
    public void foam() { append(" foam() "); }
}

public class Extend {
    public static void main(String[] args) {
        Detergent dg = new Detergent();
        dg.dilute();
        dg.foam();
        dg.scrub();
        dg.apply();
        System.out.println(dg);
    }
}
