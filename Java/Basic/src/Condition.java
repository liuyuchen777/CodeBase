/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-10 20:04:09
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-10 20:10:30
 * @Description: 
 * @FilePath: \Basic\src\Condition.java
 * @GitHub: https://github.com/liuyuchen777
 */
public class Condition {
    public void basicSwitch() {
        /*
        basic usagae of Java switch
        */
        char grade = 'C';
        
        switch(grade) {
            case 'A': System.out.println("You are good!");break;
            case 'B': System.out.println("You are not bad!");break;
            case 'C': System.out.println("You are bad!");break;
            default:
                System.out.println("Don't know!");
        }
    }

    public void basicIf() {
        /*
        basic usage of if, else, else if
        */
        int x = 30;
    
        if( x == 10 ){
            System.out.print("Value of X is 10");
        }else if( x == 20 ){
            System.out.print("Value of X is 20");
        }else if( x == 30 ){
            System.out.print("Value of X is 30");
        }else{
            System.out.print("This is else");
        }
    }
    
    public static void main(String[] args) {
        Condition sw = new Condition();
        sw.basicSwitch();
        sw.basicIf();
    }
}
