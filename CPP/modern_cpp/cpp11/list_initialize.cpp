//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"

// POD: Plain Old Data, which is able to use memcpy()

// list initialize could also be used in function return value
struct Foo {
  Foo(int, double) {}
};

Foo func() {
    return {123, 321.0};
}

int main() {
    // recap several initialize methods in C++98/03
    {
        int i_arr[3] = {1, 2, 3}; // normal array
        struct A {
            int x;
            struct {
                int i;
                int j;
            } b;
        } a = {1, {2, 3}};

        // copy initialize
        int i = 0;
        class Foo {
        public:
            Foo(int) {}
        } foo = 123; // need copy constructor

        // direct initialize
        int j(0);
        Foo bar(123);
    }
    // c++11, more powerful list initialize
    {
        class Foo {
        public:
            Foo(int) {}
        private:
            Foo(const Foo &);
        };

        Foo a1(123);
        // Foo a2 = 123;  //error: 'Foo::Foo(const Foo &)' is private
        Foo a3 = { 123 };
        Foo a4 { 123 };
        int a5 = { 3 };
        int a6 { 3 };
    }
}