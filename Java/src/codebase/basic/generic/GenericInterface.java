package codebase.basic.generic;

import java.util.*;

interface Generator<T> { T next(); }

class Coffee {
    private static long counter = 0;
    private final long id = counter++;
    public String toString() {
        return getClass().getName() + " " + id;
    }
}

class Latte extends Coffee {}

class Mocha extends Coffee {}

class Cappuccino extends Coffee {}

class Americano extends Coffee {}

class Breve extends  Coffee {}

class CoffeeGenerator implements Generator<Coffee>, Iterable<Coffee> {

    private final Class[] types = { Latte.class, Mocha.class, Cappuccino.class, Americano.class, Breve.class };

    private static final Random rand = new Random(47);

    public CoffeeGenerator() {}

    private int size = 0;

    public CoffeeGenerator(int sz) { size = sz; }

    @Override
    public Coffee next() {
        try {
            return (Coffee) types[rand.nextInt(types.length)].newInstance();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    class CoffeeIterator implements Iterator<Coffee> {
        int count = size;
        public boolean hasNext() { return count > 0; }
        public Coffee next() {
            count--;
            return CoffeeGenerator.this.next();
        }
        public void remove() {
            throw new UnsupportedOperationException();
        }
    }

    public Iterator<Coffee> iterator() {
        return new CoffeeIterator();
    }
}

public class GenericInterface {

    public static void main(String[] args) {
        CoffeeGenerator gen = new CoffeeGenerator();
        for (int i = 0; i < 5; i++) {
            System.out.println(gen.next());
        }
        for (Coffee c : new CoffeeGenerator(5)) {
            System.out.println(c);
        }
    }
}
