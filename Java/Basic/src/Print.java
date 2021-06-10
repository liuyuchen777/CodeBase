/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-09 23:30:59
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-10 19:51:03
 * @Description: print function example in Java
 * @FilePath: \Basic\src\Print.java
 * @GitHub: https://github.com/liuyuchen777
 */
import java.util.Date;

public class Print {    
    public void simplePrint() {
        // print a line
        System.out.println("Hello, World!");
        // print a string in line
        String str1 = "Hello, World!";
        System.out.println(str1);
        System.out.print("Hello World!");
        System.out.print("Hello World Again!");
    }

    /*
     * printf方法中,格式为"%s"表示以字符串的形式输出第二个可变长参数的第一个参数值;
     * 格式为"%n"表示换行;格式为"%S"表示将字符串以大写形式输出;在"%s"之间用"n$"表示
     * 输出可变长参数的第n个参数值.格式为"%b"表示以布尔值的形式输出第二个可变长参数
     * 的第一个参数值.
     */
    /*
     * 格式为"%d"表示以十进制整数形式输出;"%o"表示以八进制形式输出;"%x"表示以十六进制
     * 输出;"%X"表示以十六进制输出,并且将字母(A、B、C、D、E、F)换成大写.格式为"%e"表
     * 以科学计数法输出浮点数;格式为"%E"表示以科学计数法输出浮点数,而且将e大写;格式为
     * "%f"表示以十进制浮点数输出,在"%f"之间加上".n"表示输出时保留小数点后面n位.
     */
    /*
     * 格式为"%t"表示输出时间日期类型."%t"之后用y表示输出日期的二位数的年份(如99)、用m
     * 表示输出日期的月份,用d表示输出日期的日号;"%t"之后用Y表示输出日期的四位数的年份
     * (如1999)、用B表示输出日期的月份的完整名,用b表示输出日期的月份的简称."%t"之后用D
     * 表示以"%tm/%td/%ty"的格式输出日期、用F表示以"%tY-%tm-%td"的格式输出日期.
     */
    /*
     * "%t"之后用H表示输出时间的时(24进制),用I表示输出时间的时(12进制),用M表示输出时间
     * 分,用S表示输出时间的秒,用L表示输出时间的秒中的毫秒数、用 p 表示输出时间的是上午还是
     * 下午."%t"之后用R表示以"%tH:%tM"的格式输出时间、用T表示以"%tH:%tM:%tS"的格式输出
     * 时间、用r表示以"%tI:%tM:%tS %Tp"的格式输出时间.
     */
    /*
     * "%t"之后用A表示输出日期的全称,用a表示输出日期的星期简称.
     */

    public void formatPrint() {
        /*
        Format Print:
        1. String
        %s -> string
        %S -> string in upper letters
        %n -> new line
        $1 -> in multiple variables, use mark to distinguish
        */
        System.out.printf("%n%s%n", "A new line!");
        System.out.printf("%S%n", "a new line!");
        System.out.printf("%1$s, %2$s, %3$s", "var1", "var2", "var3");
        /*
        2. %b -> boolean
        */
        System.out.printf("%b%n", true);
        /*
        3. Integer:
        %d -> 10
        %o -> 8
        %x -> 16
        %X -> 16 in upper letter
        */
        System.out.printf("%d, %d, %d %n", -500, 2343L, 342);
        System.out.printf("%o, %o, %o %n", -500, 2343L, 342);
        System.out.printf("%x, %x, %x %n", -500, 2343L, 342);
        System.out.printf("%X, %X, %X %n", -500, 2343L, 342);
        /*
        4. float
        %e -> scentific float
        %E -> upper letter scentific float
        %f -> 10 float
        %.numberf -> Limit the number of digits after the decimal point
        */
        Double dObj = 22.22222222;
        System.out.printf("%e, %e, %e %n", -756.403f, 7464.232641d, dObj);
        System.out.printf("%E, %E, %E %n", -756.403f, 7464.232641d, dObj);
        System.out.printf("%f, %f, %f %n ", -756.403f, 7464.232641d, dObj);
        System.out.printf("%.1f, %.2f, %.3f %n", dObj, dObj, dObj);
        /*
        5. Date
        %t -> date/time
        %T -> date/time in upper letter
        ex. % + $number + ...t
        y -> year
        m -> month
        d -> day
        Y -> complete year
        B/b -> other month form
        D -> year/month/day
        F -> year-month-day
        */
        Date date = new Date();
        long dataL = date.getTime();
        System.out.printf("%1$ty-%1$tm-%1$td; %2$ty-%2$tm-%2$td%n", date, dataL);
        System.out.printf("%1$tY-%1$tB-%1$td; %2$tY-%2$tb-%2$td%n", date, dataL);
        System.out.printf("%1$tD%n", date);
        System.out.printf("%1$tF%n", date);
        /*
        6. Time
        ex. % + number + ... + t
        H -> 24 hour
        I -> 12 hour
        M -> minute
        S -> second
        L -> ms
        p -> am/pm
        R -> hour:minute
        T -> hour:minute:second
        r -> hour:minute:second am/pm
        A -> full day of week
        a -> short day of week
        c -> full date information
        */
        System.out.printf("%1$tH:%1$tM:%1$tS, %2$tI:%2$tM:%2$tS%n", date, dataL);
        System.out.printf("%1$tH:%1$tM:%1$tS, %1$tL%n", date);
        System.out.printf("%1$tH:%1$tM:%1$tS, %1$tL, %1$tp%n", date);
        System.out.printf("%1$tR%n", date);
        System.out.printf("%1$tT%n", date);
        System.out.printf("%1$tr%n", date);
        System.out.printf("%1$tF, %1$tA %n", date);
        System.out.printf("%1$tF, %1$ta %n", date);
        System.out.printf("%1$tc%n", date);
    }

    public static void main(String[] args) {
        Print print = new Print();
        print.simplePrint();
        print.formatPrint();
    }
}
