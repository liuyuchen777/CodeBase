/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-12 09:20:17
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-13 17:46:55
 * @Description: example for Java container: ArrayList
 * @FilePath: \Container\src\ArrayListTest.java
 * @GitHub: https://github.com/liuyuchen777
 */
import java.util.ArrayList;
import java.util.Iterator;
import java.util.Collections;

public class ArrayListTest {
    public void basicOperate() {
        System.out.println("-----------------Basic Operate---------------");
        ArrayList<String> sites = new ArrayList<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Weibo");
        System.out.println(sites);
        // get specific elements
        System.out.println(sites.get(1));
        // change element
        sites.set(1, "wiki");
        System.out.println(sites);
        // remove element
        sites.remove(3);
        System.out.println(sites);
        // size of list
        System.out.println(sites.size());
        // iterate list
        for (int i = 0; i < sites.size(); i++) {
            System.out.println("***general loop***");
            System.out.println(sites.get(i));
        }
        for (String item : sites) {
            System.out.println("***advance loop***");
            System.out.println(item);
        }
        for (Iterator<String> it = sites.iterator(); it.hasNext(); ) {
            System.out.println("***iterator loop***");
            System.out.println(it.next());
        }
        /* 
        use other type need to use wrapper class:
        基本类型  | 引用类型
        --------------------
        boolean	    Boolean
        byte	    Byte
        short	    Short
        int	        Integer
        long	    Long
        float	    Float
        double	    Double
        char	    Character
        */
        ArrayList<Integer> myNumbers = new ArrayList<Integer>();
        myNumbers.add(10);
        myNumbers.add(15);
        myNumbers.add(20);
        myNumbers.add(25);
        for (int i : myNumbers) {
            System.out.println("***ArrayList Integer***");
            System.out.println(i);
        }
    }
    
    public void sortOperate() {
        System.out.println("-------------ArrayList sort-------------");
        ArrayList<String> sites = new ArrayList<String>();
        sites.add("Taobao");
        sites.add("Wiki");
        sites.add("Runoob");
        sites.add("Weibo");
        sites.add("Google");
        Collections.sort(sites);  // 字母排序
        for (String i : sites) {
            System.out.println(i);
        }
        
        ArrayList<Integer> myNumbers = new ArrayList<Integer>();
        myNumbers.add(33);
        myNumbers.add(15);
        myNumbers.add(20);
        myNumbers.add(34);
        myNumbers.add(8);
        myNumbers.add(12);
        Collections.sort(myNumbers);  // 数字排序
        for (int i : myNumbers) {
            System.out.println(i);
        }
    }

    public void advanceOperate() {
        System.out.println("-------------Other Operations-------------");
        // add all element
        ArrayList<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(3);
        System.out.println("ArrayList: " + list);
        list.forEach((e) -> {
            e = e * 10;
            System.out.println(e);
        });
    }

    public static void main(String[] args) {
        ArrayListTest list = new ArrayListTest();
        list.basicOperate();
        list.sortOperate();
        list.advanceOperate();
    }
}