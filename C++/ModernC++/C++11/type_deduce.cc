/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-08 22:18:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-08 22:19:06
 * @Description: auto, decltype example in C++11
 * @FilePath: /CodeBase/C++/ModernC++/C++11/type_dudce.cc
 * @GitHub: https://github.com/liuyuchen777
 */

#include "c++11_test.h"

class A{
public:
    static int get(void){
        return 100;
    }
};

class B{
public:
    static const char* get(void){
        return "https://liuyuchen777.github.io";
    }
};

template <typename T>
void func(void){
    auto val = T::get();
    cout << val << endl;
}

void auto_deduce() {
    {
        // basic
        auto n = 10; // int
        auto f = 12.8; // double
        auto x = &n; // int *
        auto url = "https://liuyuchen777.github.io"; // const char *
        auto *p = &n, m = 99;
    }
    {
        // advance
        int x = 0;
        auto *p1 = &x; // p1 is int *, auto is int
        auto p2 = &x; // p2 is int *, auto is int *
        auto &r1 = x; // r1 is int &, auto is int
        auto r2 = x; // r2 is int
    }
    {
        // auto and const
        int x = 0;
        const auto n = x; // n is const int, auto is int
        auto f = n; // f is int, auto is int (const is abandoned)
        const auto &r1 = x; // r1 is const int &, auto is int
        auto &r2 = r1; // r1 is const int &, auto is const int

        // recap:
        // when type is not reference, auto deduction will not keep const
        // when type is reference, auto deduction will keep const
    }
    {
        // application and notice of auto usage
        // 1. in STL iterator

        // 2. in generic programming
        func<A>();
        func<B>();
    }
}

void decltype_deduce() {
    {
        // basic
        // usage: decltype(exp) varname
        int a = 0;
        decltype(a) b = 1;

    }
}

int main() {
    cout << "--------------type deduce!----------------" << endl;
    auto_deduce();
}