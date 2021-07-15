//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"

class MoveDemo {
public:
    MoveDemo() : num(new int(0)) {
        std::cout << "!!!construct!!!" << std::endl;
    }
    MoveDemo(const MoveDemo &d) : num(new int(*d.num)) {
        std::cout << "!!!copy construct!!!" << std::endl;
    }
    // use for std::move
    MoveDemo(MoveDemo &&d)  noexcept : num(d.num) {
        std::cout << "!!!move construct!!!" << std::endl;
    }
public:
    int *num;
};

int main() {
    /*
     * left value is exp at left of = (it could also be able to be at right side)
     * right value is exp at right side and can't be at left side
     */
    {
        int foo = 42;
        int bar(43);

        // foo, bar is left value
        foo = bar;
        bar = foo;
        foo = foo * bar;

        // foo * bar is right value
        int baz;
        baz = foo * bar;
        // foo * bar = 42;   // error
    }

    /*
     * left reference vs. right reference
     * left reference:
     * bind address of a obj to another variable
     * we can't bind a right value to left reference
     * const left reference assure us can't change according memory space
     * right reference:
     * rr can only be bind on right value
     * right reference object is temporary
     */
    {
        int foo(42);
        int &bar = foo;
        // int &baz = 42; // error
        const int &qux = 42; // 42 is right value, but compiler give a space for 42 because it;s const
    }
    {
        int foo(42);
        int &bar = foo;
        // int &&baz = foo; // error, foo is left value
        int &&qux = 42;
        int &&quux = foo * 1;
        // int &garoly = foo++; // error, ++return left value
        int &&waldo = foo--;
    }
    /*
     * move (base on right value reference)
     * C++11 标准中借助右值引用可以为指定类添加移动构造函数，这样当使用该类的右值对象（可以理解为临时对象）
     * 初始化同类对象时，编译器会优先选择移动构造函数。
     */
    {
        MoveDemo demo;
        std::cout << "demo2: " << std::endl;
        MoveDemo demo2 = demo;
        std::cout << "demo3: " << std::endl;
        MoveDemo demo3 = std::move(demo);
        // not a good practice to write like following, demo.num is moved
        std::cout << *demo.num << endl;
    }
}