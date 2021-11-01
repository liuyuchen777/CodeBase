
public class Condition {
    public void basicSwitch() {
        /*
        basic usagae of Java switch
        */
        char grade = 'C';

        switch (grade) {
            case 'A' -> System.out.println("You are good!");
            case 'B' -> System.out.println("You are not bad!");
            case 'C' -> System.out.println("You are bad!");
            default -> System.out.println("Don't know!");
        }
    }

    public void basicIf() {
        /*
        basic usage of if, else, else if
        */
        int x = 30;
    
        if( x == 10 ){
            System.out.println("Value of X is 10");
        } else if( x == 20 ) {
            System.out.println("Value of X is 20");
        } else if( x == 30 ) {
            System.out.println("Value of X is 30");
        } else {
            System.out.println("This is else");
        }
    }
    
    public static void main(String[] args) {
        Condition sw = new Condition();
        sw.basicSwitch();
        sw.basicIf();
    }
}
