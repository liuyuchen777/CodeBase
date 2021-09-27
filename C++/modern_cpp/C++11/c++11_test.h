/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-08 22:15:41
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-08 22:18:02
 * @Description: 
 * @FilePath: /CodeBase/C++/ModernC++/C++11/c++11_test.h
 * @GitHub: https://github.com/liuyuchen777
 */

#ifndef _CPP11_H_
#define _CPP11_H_


#include <iostream>
#include <map>

using namespace std;

class Student {
public:
    string name;
    int age;
    float scores;
    static int total;
};

// for decltype
int& func_int_r(int, char);  //返回值为 int&
int&& func_int_rr(void);  //返回值为 int&&
int func_int(double);  //返回值为 int
const int& fun_cint_r(int, int, int);  //返回值为 const int&
const int&& func_cint_rr(void);  //返回值为 const int&&

class Base {
public:
    int x;
};

/**
 * brief: auto in c++11
 */
void auto_deduce();

/**
 * brief: decltype in c++11
 */
void decltype_deduce();

/*
 * back set of return type
 */
template <typename T, typename U>
auto add(T t, U u) -> decltype(t + u) {
    return t + u;
}

#endif