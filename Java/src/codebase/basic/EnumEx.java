package codebase.basic;

// in Java, enum is kind of class
enum Spiciness {
    NOT, MILD, MEDIUM, HOT, FLAMING
}

public class EnumEx {
    Spiciness degree;

    public EnumEx(Spiciness degree) {
        this.degree = degree;
    }

    public void describe() {
        System.out.print("The spiciness of this is: ");
        switch (degree) {
            case NOT -> System.out.println("You are week!");
            case MILD -> System.out.println("You are not so week!");
            case MEDIUM -> System.out.println("You are not bad!");
            case HOT -> System.out.println("You are pretty strong!");
            case FLAMING -> System.out.println("You are super string!");
        }
    }

    public static void main(String[] args) {
        for (Spiciness s : Spiciness.values()) {
            System.out.println(s + ", " + s.ordinal());
        }

        // use enum in switch
        EnumEx em = new EnumEx(Spiciness.HOT);
        em.describe();
    }
}
