package codebase.basic.oop;

class Cleanser {
    private String s = "codebase.basic.oop.Cleanser ";
    public void append(String a) { s += a; }
    public void dilute() { append(" dilute() "); }
    public void apply() { append(" apply() "); }
    public void scrub() { append(" scrub() "); }
    public String toString() { return s; }
}

class Detergent extends Cleanser {
    public void scrub() {
        // cover scrub() in codebase.basic.oop.Cleanser
        append(" codebase.basic.oop.Detergent.srub() ");
        super.scrub();
    }
    public void foam() { append(" foam() "); }
}

public class Extend {
    public static void main(String[] args) {
        Detergent dg = new Detergent();
        dg.dilute();
        dg.foam();
        dg.scrub();
        dg.apply();
        System.out.println(dg);
    }
}