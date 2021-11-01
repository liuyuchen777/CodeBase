/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-30 10:36:00
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-30 10:47:23
 * @Description: 
 * @FilePath: /CodeBase/C++/GoogleTest/fun.cc
 * @GitHub: https://github.com/liuyuchen777
 */
#include "fun.h"

double square(double x) {
    return x * x;
}

int Factorial(int n) {
    if (n == 0)
        return 1;
    return n * Factorial(n - 1);
}

