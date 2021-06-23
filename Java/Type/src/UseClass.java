/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-23 23:42:05
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-23 23:51:18
 * @Description: 
 * @FilePath: /CodeBase/Java/Type/src/UseClass.java
 * @GitHub: https://github.com/liuyuchen777
 */

interface HasBatteries {}
interface WaterProof {}
interface Shoots {}

class Toy {
    Toy() {}
    Toy(int i) {}
}

class FancyToy extends Toy 
implements HasBatteries, WaterProof, Shoots {
    FancyToy() { super(1); }
}

public class UseClass {
    static void printInfo(Class cc) {
        System.out.println("Class name: " + cc.getName() + 
                    " is interface? [" + cc.isInterface() + "]");
        System.out.println("Simple name: " + cc.getSimpleName());
        System.out.println("Canonical name: " + cc.getCanonicalName());
    }
    public static void main(String[] args) {
        Class c = null;
        try {
            c = Class.forName("FancyToy");
        } catch (ClassNotFoundException e) {
            System.out.println("Can't find FancyToy!");
            System.exit(1);
        }
        printInfo(c);
        // print all interface
        for (Class face : c.getInterfaces()) 
            printInfo(face);
        // get base class
        Class up = c.getSuperclass();
        Object obj = null;
        try {
            // requires default constructor
            obj = up.newInstance();
        } catch(InstantiationException e) {
            System.out.println("Cannot instantiate!");
            System.exit(1);
        } catch(IllegalAccessException e) {
            System.out.println("Cannot access!");
            System.exit(1);
        }
        printInfo(obj.getClass());
    }
}
