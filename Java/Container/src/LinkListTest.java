/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-12 22:00:11
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-13 17:52:31
 * @Description: 
 * @FilePath: \Container\src\LinkListTest.java
 * @GitHub: https://github.com/liuyuchen777
 */

import java.util.LinkedList;

public class LinkListTest {
    /*
    Basic operation of LinkList is same as ArrayList
    But data storage in LL is not linear
    */
    public void generalOperation() {
        System.out.println("------------Basic Operation-----------");
        LinkedList<String> sites = new LinkedList<String>();
        // add
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Weibo");
        System.out.println(sites);
        // add at head
        sites.addFirst("Recruit");
        System.out.println(sites);
        // add at tail
        sites.addLast("Facebook");
        System.out.println(sites);
        // remove
        // remove head element
        sites.removeFirst();
        System.out.println(sites);
        // remove tail element
        sites.removeLast();
        System.out.println(sites);
        // get first
        System.out.println(sites.getFirst());
        // get last
        System.out.println(sites.getLast());
    }

    public static void main(String[] args) {
        LinkListTest list = new LinkListTest();
        list.generalOperation();
    }
}
