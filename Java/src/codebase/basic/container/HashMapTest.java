package codebase.basic.container;

import java.util.HashMap;

public class HashMapTest {
    public void basicOperate() {
        System.out.println("-----------codebase.basic.oop.Basic Operation--------");
        // <Integer, String>
        HashMap<Integer, String> sites = new HashMap<>();
        sites.put(1, "Google");
        sites.put(3, "Microsoft");
        sites.put(5, "Zhihu");
        System.out.println(sites);
        HashMap<String, String> sites2 = new HashMap<String, String>();
        // <String, String>
        sites2.put("one", "Google");
        sites2.put("two", "Runoob");
        sites2.put("three", "Taobao");
        sites2.put("four", "Zhihu");
        System.out.println(sites2);
        // get element
        System.out.println(sites2.get("one"));
        // remove element
        System.out.println(sites2.remove("two"));
        System.out.println(sites2);
    }

    public void iterateOperate() {
        System.out.println("-------------Iterate Operation------------");
        HashMap<String, String> sites = new HashMap<String, String>();
        // <String, String>
        sites.put("one", "Google");
        sites.put("two", "Runoob");
        sites.put("three", "Taobao");
        sites.put("four", "Zhihu");
        // loop base on key value
        for (String i : sites.keySet()) {
            System.out.println("key: " + i + ", value: " + sites.get(i));
        }
        for (String i : sites.values()) {
            System.out.println(i);
        }
    }

    public static void main(String[] args) {
        HashMapTest hm = new HashMapTest();
        hm.basicOperate();
        hm.iterateOperate();
    }
}
