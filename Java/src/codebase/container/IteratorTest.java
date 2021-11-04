package codebase.container;

import java.util.*;

class PeopleItem {
    final private String name;
    final private int id;
    private static int count = 0;
    // constructor
    public PeopleItem() {
        this.id = count;
        count++;
        this.name = usingUUID();
    }
    // static methods
    public static ArrayList<PeopleItem> arrayList(int number) {
        ArrayList<PeopleItem> tmpList = new ArrayList<>(number);
        for (int i = 0; i < number; i++) {
            tmpList.add(new PeopleItem());
        }
        return tmpList;
    }

    public static String usingUUID() {
        UUID randomUUID = UUID.randomUUID();
        return randomUUID.toString().replaceAll("-", "");
    }

    @Override
    public String toString() {
        return "codebase.container.PeopleItem{" +
                "name='" + name + '\'' +
                ", id=" + id +
                '}';
    }
}

class ReversibleArrayList<T> extends ArrayList<T> {
    public ReversibleArrayList(Collection<T> c) { super(c); }

    public Iterable<T> reversed() {
        return new Iterable<T>() {
            @Override
            public Iterator<T> iterator() {
                return new Iterator<T>() {
                    private int current = size() - 1;
                    @Override
                    public boolean hasNext() {
                        return current > -1;
                    }

                    @Override
                    public T next() {
                        return get(current--);
                    }

                    @Override
                    public void remove() {
                        throw  new UnsupportedOperationException();
                    }
                };
            }
        };
    }
}

public class IteratorTest {
    public static void display(Iterator<PeopleItem> it) {
        while (it.hasNext()) {
            PeopleItem item = it.next();
            System.out.println(item);
        }
    }

    public static void main(String[] args) {
        List<PeopleItem> peoples = PeopleItem.arrayList(12);
        System.out.println(peoples);

        // iterator
        System.out.println("-------------Iterator Approach------------");
        Iterator<PeopleItem> it = peoples.iterator();
        while (it.hasNext()) {
            PeopleItem item = it.next();
            System.out.println(item);
        }
        System.out.println("-------------Simpler Approach------------");
        for (PeopleItem item : peoples) {
            System.out.println(item);
        }
        // remove item
        System.out.println("-------------Remove Item------------");
        it = peoples.iterator();
        for (int i = 0; i < 6; i++) {
            it.next();
            it.remove();
        }
        System.out.println(peoples);
        System.out.println("-------------With API------------");
        display(peoples.iterator());

        // list Iterator
        System.out.println("-------------List Iterator------------");
        ListIterator<PeopleItem> lit = peoples.listIterator();
        while (lit.hasNext()) {
            System.out.println(lit.next() + ", " + lit.nextIndex() + ", " + lit.previousIndex());
        }
        System.out.println("-------------Inverse Traversal------------");
        while (lit.hasPrevious()) {
            System.out.println(lit.previous());
        }
        System.out.println("-------------Index------------");
        lit = peoples.listIterator(3);  // ready for return third element, start from 0
        System.out.println(lit.next());

        System.out.println("-------------Reverse Iterator------------");
        ReversibleArrayList<String> ral = new ReversibleArrayList<>(
                Arrays.asList("To be or not to be".split(" "))
        );
        for (String s : ral) {
            System.out.print(s + " ");
        }
        System.out.println();
        for (String s : ral.reversed()) {
            System.out.print(s + " ");
        }
    }
}
