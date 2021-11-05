package codebase.basic.generic;

// tuple
class TwoTuple<A, B> {
    public final A first;
    public final B second;
    public TwoTuple(A a, B b) {
        first = a;
        second = b;
    }
    public String toString() {
        return "(" + first + ", " + second + ")";
    }
}

// stack
class LinkedStack<T> {
    private static class Node<U> {
        U item;
        Node<U> next;
        Node() { item = null; next = null; }
        Node(U item, Node<U> next) {
            this.item = item;
            this.next = next;
        }
        boolean end() { return item == null && next == null; }
    }

    private Node<T> top = new Node<>();

    public void push(T item) {
        top = new Node<T>(item, top);
    }

    public T pop() {
        T result = top.item;
        if (!top.end())
            top = top.next;
        return result;
    }
}

public class GenericBasic {
    static TwoTuple<String, Integer> f() {
        return new TwoTuple<>("hi", 47);
    }

    public static void main(String[] args) {
        System.out.println("--------------Tuple------------");
        TwoTuple<String, Integer> tmp = f();
        System.out.println(tmp);
        // tmp.first = "World"; // compile error

        System.out.println("--------------Stack------------");
        LinkedStack<String> lss = new LinkedStack<>();
        for (String s : "To be or not to be".split(" ")) {
            lss.push(s);
        }
        String s;
        while ((s = lss.pop()) != null) {
            System.out.println(s);
        }
    }
}
