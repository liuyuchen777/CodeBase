/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-18 09:31:38
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-18 09:36:31
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/ThisTest.java
 * @GitHub: https://github.com/liuyuchen777
 */

class Flower {
    int pealCount = 0;
    String s = "initial value";

    Flower(int peals) {
        pealCount = peals;
        System.out.println("[Constructor] int only");
    }

    Flower(String ss) {
        s = ss;
        System.out.println("[Constructor] String only");
    }

    Flower(int peals, String s) {
        this(peals);
        // this(ss); // can't use this to call two constructor, and this need to at the first line
        this.s = s;
        System.out.println("[Constructor] int and String");
    }
}

public class ThisTest {
    public static void main(String[] args) {
        
    }
}