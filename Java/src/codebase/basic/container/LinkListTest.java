package codebase.basic.container;

import java.util.LinkedList;

public class LinkListTest {
    /*
    codebase.basic.oop.Basic operation of LinkList is same as ArrayList
    But data storage in LL is not linear
    */
    public void generalOperation() {
        System.out.println("------------codebase.basic.oop.Basic Operation-----------");
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
