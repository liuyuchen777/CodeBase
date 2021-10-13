/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-10 20:10:56
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-10 20:30:17
 * @Description: 
 * @FilePath: \Basic\src\NumberMath.java
 * @GitHub: https://github.com/liuyuchen777
 */
public class NumberMath {
    /*
    Object
    ------------------
    |       |        |
    number  Boolean  Character
    |
    --------------------------------
    |    |     |       |     |     |
    Byte Short Integer Long  Float Double
    */
    public void javaMath() {
        System.out.println("------------basic math-----------");
        System.out.println("90 degree sin: " + Math.sin(Math.PI/2));  
        System.out.println("0 degree cos: " + Math.cos(0));  
        System.out.println("60 degree tan: " + Math.tan(Math.PI/3));  
        System.out.println("1 atan: " + Math.atan(1));  
        System.out.println("π/2 to degree: " + Math.toDegrees(Math.PI/2));  
        System.out.println(Math.PI); 
    }

    public void typeConvertor() {
        System.out.println("---------type convertor----------");
        /*
        xxxValue() -> convert number type to xxx number type
        byteValue()
        doubleValue()
        floatValue()
        intValue()
        longValue()
        shortValue()
        */
        Integer x = 5;

        System.out.println( x.byteValue() );
        System.out.println(x.doubleValue());
        System.out.println( x.longValue() );
        /*
        xxx.valueOf() -> convert other type to xxx number type
        */  
        Integer z = Integer.valueOf(9);
        Double c = Double.valueOf(5);
        Float a = Float.valueOf("80");              

        Integer b = Integer.valueOf("444",16); 

        System.out.println(z);
        System.out.println(c);
        System.out.println(a);
        System.out.println(b);
        /*
        toString() -> convert Integer to string
        */
        Integer intNum = 5;

        System.out.println(intNum.toString());  
        System.out.println(Integer.toString(12));
        /*
        parseInt() -> convert string to int
        */
        int xx =Integer.parseInt("9");
        double cc = Double.parseDouble("5");
        int bb = Integer.parseInt("444", 16);

        System.out.println(xx);
        System.out.println(cc);
        System.out.println(bb);
    }

    public void processNumber() {
        // round or ceil a number
    }

    public void compareNumber() {
        System.out.println("--------------compare numbers-------------");
        /*
        compare two numbers
        compareTo()
        1 -> larger
        0 -> equal
        -1 -> smaller
        */
        Integer m = 5;
        System.out.println(m.compareTo(3));
        System.out.println(m.compareTo(5));
        System.out.println(m.compareTo(8));
        /*
        public boolean equals(Object o)
        true/false whether equal
        */    
        Integer x = 5;
        Integer y = 10;
        Integer z =5;
        Short a = 5;

        System.out.println(x.equals(y));  
        System.out.println(x.equals(z));
        System.out.println(x.equals(a));
    }

    public static void main(String[] args) {
        NumberMath nm = new NumberMath();
        nm.javaMath();
        nm.typeConvertor();
        nm.compareNumber();
    }
}