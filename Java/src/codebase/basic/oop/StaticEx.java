package codebase.basic.oop;/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-19 10:05:30
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-19 10:11:49
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/codebase.basic.oop.StaticEx.java
 * @GitHub: https://github.com/liuyuchen777
 */

/*
    由于Java中所有变量都在类中，且类仅在使用时才加载，因此不会出现C++中需要static相互依赖对方先初始化而产生错误
    Java的static变量存在JVM的数据区（Data Segment）
*/

/*
    Java加载顺序:
    1. find static main
    2. if initial a class, load its .class
    3. if it extends from other base class, load its base class's .class
    4. util to the most base class, start initial static element
    5. then initial class element as default value (through set memory 0 value)
*/

class GlobalVar {
    public static int i = 10;
    public static String str = "lyc";
}

public class StaticEx {
    public static void main(String[] args) {
        
    }
}
