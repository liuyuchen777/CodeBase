
public class StaticTest {
    /** 
     * Two Reasons for static:
     * 1. you want to have only one copy of data in the whole program -> gloabl variable
     * 2. you want to use a method without create a object -> global function
     */

    public static int i = 1;

    public void increment() {
        i++;
    }

    public static void sayHello() {
        System.out.println("i is " + i);
    }

    public void showProperties() {
        /**
         * show system properties
         */
        System.out.println("-----------Show System Properties--------");
        System.getProperties().list(System.out);
        System.out.println("-----------Show Some Properties---------");
        System.out.println("user.name: " + System.getProperty("user.name"));
        System.out.println("java.library.path: " + System.getProperty("java.library.path"));
    }
    
    public static void main(String[] args) {
        StaticTest t1 = new StaticTest();
        StaticTest t2 = new StaticTest();

        t1.increment();
        t2.increment();
        
        StaticTest.sayHello();

        t1.showProperties();
    }
}