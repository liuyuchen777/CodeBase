/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-12 22:00:02
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-13 17:57:39
 * @Description: example of Java hash set
 * @FilePath: \Container\src\HashSetTest.java
 * @GitHub: https://github.com/liuyuchen777
 */

import java.util.HashSet;

public class HashSetTest {
    public void basicOperation() {
        System.out.println("--------------Basic Operation-------------");
        // add 
        HashSet<String> sites = new HashSet<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Zhihu");
        sites.add("Runoob");  // the repeat element won't be added
        System.out.println(sites);
        // judge element whether exist
        System.out.println(sites.contains("Taobao"));
        // delete element
        sites.remove("Taobao");
        System.out.println(sites);
        // clear all element
        sites.clear();
        System.out.println(sites);
    }
    
    public static void main(String[] args) {
        HashSetTest hs = new HashSetTest();
        hs.basicOperation();
        
    }
}
