package codebase.type;

import java.util.*;

abstract class Shape {

    void draw() {
        System.out.println(this + ".draw()");
    }

    abstract public String toString();
}

class Cricle extends Shape {

    @Override
    public String toString() {
        return "Circle";
    }
}

class Square extends Shape {

    @Override
    public String toString() {
        return "codebase.type.Square";
    }
}

class Triangle extends Shape {

    @Override
    public String toString() {
        return "codebase.type.Triangle";
    }
}

class CandyItem {
    static { System.out.println("Loading codebase.type.Candy!"); }
}

class GumItem {
    static { System.out.println("Loading codebase.type.Gum!"); }
}

class CookieItem {
    static { System.out.println("Loading codebase.type.Cookie!"); }
}

public class RTTIBasic {
    public static void main(String[] args) {
        System.out.println("------------codebase.type.Shape List-----------");
        List<Shape> shapeList = Arrays.asList(
                new Cricle(), new Square(), new Triangle()
        );
        for (Shape shape : shapeList) {
            shape.draw();
        }

        System.out.println("-----------Load Class-----------");
        System.out.println("Inside main");
        new Candy();
        System.out.println("After creating candy");
        try {
            Class.forName("codebase.type.GumItem");
        } catch (ClassNotFoundException e) {
            System.out.println("Couldn't find codebase.type.Gum");
        }
        System.out.println("After Class.forName()");
        new CookieItem();
        System.out.println("After creating codebase.type.Cookie");
    }
}
