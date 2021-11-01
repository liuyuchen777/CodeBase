/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-29 21:21:17
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-30 10:48:25
 * @Description: 
 * @FilePath: /CodeBase/C++/GoogleTest/hello_test.cc
 * @GitHub: https://github.com/liuyuchen777
 */

#include <gtest/gtest.h>
#include <iostream>
#include "fun.h"

using namespace std;

// Demonstrate some basic assertions.
TEST(HelloTest, BasicAssertions) {
    EXPECT_EQ(7 * 6, 42);
}

TEST(funTest, BasicAssertions) {
    EXPECT_EQ(square(10.0), 100.0) << "Square test is successful!";
}

// Tests factorial of 0.
TEST(FactorialTest, HandlesZeroInput) {
    EXPECT_EQ(Factorial(0), 1);
}

// Tests factorial of positive numbers.
TEST(FactorialTest, HandlesPositiveInput) {
    EXPECT_EQ(Factorial(1), 1);
    cout << "Test Output in Console!" << endl;
    EXPECT_EQ(Factorial(2), 2);
    EXPECT_EQ(Factorial(3), 6);
    EXPECT_EQ(Factorial(8), 40320);
}