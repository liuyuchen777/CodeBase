/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-20 12:29:05
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-23 20:35:40
 * @Description: anything about inner class
 * @FilePath: /CodeBase/Java/OOP/src/InnerClass.java
 * @GitHub: https://github.com/liuyuchen777
 */

/**
 * Inner class could be used to realize interator
 */

interface Selector {
    boolean end();
    Object current();
    void next();
}

class Sequence {
    private Object[] items;
    private int next = 0;
    public Sequence(int size) { items = new Object[size]; }
    public void add(Object x) {
        if (next < items.length)
            items[next++] = x;
    }
    private class SequenceSelector implements Selector {
        private int i = 0;
        public boolean end() { return i == items.length; }
        public Object current() { return items[i]; }
        public void next() { if (i < items.length) i++; }
    }
    public Selector selector() {
        return new SequenceSelector();
    }
}

interface Destination {
    String readLabel();
}

interface Contents {
    int value();
}

// 局部内部类 
class Parcel5 {
    public Destination destination(String s) {
        class PDestination implements Destination {
            private String label;
            private PDestination(String whereTo) {
                label = whereTo;
            }
            public String readLabel() {
                return label;
            }
        }

        return new PDestination(s);
    }
}

// 作用域中嵌入内部类
class Parcel6 {
    private void internalTracking(boolean b) {
        if (b) {
            class TrackingSlip {
                private String id;
                TrackingSlip(String s) {
                    id = s;
                }
                String getSlip() { return id; }
            }
            TrackingSlip ts = new TrackingSlip("slip");
            String s = ts.getSlip();
        }
    }
    public void track() { internalTracking(true); }
}

// 匿名内部类
class Parcel7 {
    public Contents contents() {
        return new Contents() {
            private int i = 11;
            public int value() { return i; }
        };
    }
}

// 匿名内部类的构造器是带参的
class Wrapping {
    private int i;
    public Wrapping(int x) { i = x; }
    public int value() { return i; }
}

class Parcel8 {
    public Wrapping wrapping(int x) {
        // Base constructor call
        return new Wrapping(x) {
            public int value() {
                return super.value() * 47;
            }
        };
    }
}

public class InnerClass {
    public static void main(String[] args) {
        Sequence sequence = new Sequence(10);
        for (int  i = 0; i < 10; i++)
            sequence.add(Integer.toString(i));
        Selector selector = sequence.selector();
        while (!selector.end()) {
            System.out.print(selector.current() + " ");
            selector.next();
        }
        System.out.println();
    }
}
