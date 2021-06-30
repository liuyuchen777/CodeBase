import java.util.Arrays;
import java.util.List;

/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-30 11:39:16
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-30 11:46:40
 * @Description: 
 * @FilePath: /CodeBase/Java/Basic/src/MethodReference.java
 * @GitHub: https://github.com/liuyuchen777
 */
@FunctionalInterface
interface Supplier<T> {
    T get();
}

class Car {
    //Supplier是jdk1.8的接口，这里和lamda一起使用了
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

    // public String toString() {
    //     return "This is a Car!";
    // }
}

public class MethodReference {
    public static void main(String[] args) {
        final Car car = Car.create(Car::new);
        System.out.println(car);
        final List<Car> cars = Arrays.asList(car);
        cars.forEach(Car::collide);
        cars.forEach(Car::repair);
        // 特定对象的方法引用
        final Car police = Car.create( Car::new );
        cars.forEach(police::follow);
    }
}
