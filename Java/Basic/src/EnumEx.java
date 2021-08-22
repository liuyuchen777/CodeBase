/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-18 14:23:42
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-18 14:34:51
 * @Description: 
 * @FilePath: /CodeBase/Java/Basic/src/EnumEx.java
 * @GitHub: https://github.com/liuyuchen777
 */

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
        switch(degree) {
            case NOT: System.out.println("You are week!"); break;
            case MILD: System.out.println("You are not so week!"); break;
            case MEDIUM: System.out.println("You are not bad!"); break;
            case HOT: System.out.println("You are pretty strong!"); break;
            case FLAMING: System.out.println("You are super string!"); break;
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
