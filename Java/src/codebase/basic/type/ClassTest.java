package codebase.basic.type;

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

public class ClassTest {
    static void printInfo(Class cc) {
        System.out.println("Class name: " + cc.getName() + " is interface? [" + cc.isInterface() + "]");
        System.out.println("Simple name: " + cc.getSimpleName());
        System.out.println("Canonical name: " + cc.getCanonicalName());
    }

    public static void main(String[] args) {
        Class c = null;
        try {
            c = Class.forName("codebase.basic.type.FancyToy");
        } catch (ClassNotFoundException e) {
            System.out.println("Can't find codebase.basic.type.FancyToy!");
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