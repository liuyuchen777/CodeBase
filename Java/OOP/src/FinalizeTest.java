/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-18 09:44:31
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-18 09:50:21
 * @Description: 
 * @FilePath: /CodeBase/Java/OOP/src/FinalizeTest.java
 * @GitHub: https://github.com/liuyuchen777
 */

class Book {
    boolean checkOut = false;
    
    Book(boolean checkOut) {
        this.checkOut = checkOut;
    }

    void checkIn() {
        checkOut = false;
    }

    protected void finalize() {
        // finalize is used to label object to be cleaned, but not necessary be cleaned
        if (checkOut)
            System.out.println("Error: checked out");
        // call base-class finalize function
        // super.finalize();
    }
}

public class FinalizeTest {
    public static void main(String[] args) {
        Book novel = new Book(true);
        // proper cleanup
        novel.checkIn();
        new Book(true);
        System.gc(); // call gc to clean garbage
    }
}