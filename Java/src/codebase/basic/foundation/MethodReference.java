package codebase.basic.foundation;

import java.util.List;

@FunctionalInterface
interface Supplier<T> {
    T get();
}

class Car {
    //Supplier是jdk1.8的接口，这里和lambda一起使用了
    public static Car create(final Supplier<Car> supplier) {
        return supplier.get();
    }

    public static void collide(final Car car) {
        System.out.println("Collided: " + car.toString());
    }

    public void follow(final Car car) {
        System.out.println("Following the: " + car.toString());
    }

    public void repair() {
        System.out.println("Repaired: " + this.toString());
    }

}

public class MethodReference {
    public static void main(String[] args) {
        final Car car = Car.create(Car::new);
        System.out.println(car);
        final List<Car> cars = List.of(car);
        cars.forEach(Car::collide);
        cars.forEach(Car::repair);
        // 特定对象的方法引用
        final Car police = Car.create( Car::new );
        cars.forEach(police::follow);
    }
}
