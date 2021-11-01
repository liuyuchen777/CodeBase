
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class Loop {
    public void whileLoop() {
        int x = 10;

        System.out.println("--------while loop---------");

        while (x < 20) {
            System.out.print("value of x: " + x);
            x++;
            System.out.println();
        }
    }

    public void doWhileLoop() {
        int x = 10;

        System.out.println("--------do..while loop------");
    
        do {
            System.out.print("value of x : " + x );
            x++;
            System.out.println();
        } while( x < 20 );
    }

    public void forLoop() {
        System.out.println("---------simple for loop----------");

        for (int x = 10; x < 20; x++) {
            System.out.println("value of x: " + x);
        }

        System.out.println("--------enhanced for loop----------");

        int[] numbers = {10, 20, 30, 40, 50};
        for (int x : numbers) {
            System.out.println("value of number: " + x);
        }

        String[] names = {"Liu", "Zhao", "Wang", "Jiang"};
        for (String x : names) {
            System.out.println("value of name: " + x);
        }

        /*
        Interator
        */
        List<String> list = new ArrayList<String>();
        
        list.add("Google");
        list.add("Runoob");
        list.add("Taobao");
        System.out.println("----------iterator loop-----------");
        for(Iterator<String> iter = list.iterator(); iter.hasNext();)
        {
            System.out.println(iter.next());
        }
        System.out.println("---------list for each loop----------");
        for(String str: list)
        {
            System.out.println(str);
        }

        /**
         * any array could use forEach
         * String can use toCharArray() to convert a char array
         */
        for (char c : "I am king of the world".toCharArray()) {
            System.out.print(c + " ");
        }
        System.out.println("");
    }

    public static void main(String[] args) {
        Loop loop = new Loop();
        loop.whileLoop();
        loop.doWhileLoop();
        loop.forLoop();
    }
}
