/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-29 21:21:17
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-29 21:21:42
 * @Description: 
 * @FilePath: /CodeBase/C++/GoogleTest/hello_test.cc
 * @GitHub: https://github.com/liuyuchen777
 */

#include <gtest/gtest.h>

// Demonstrate some basic assertions.
TEST(HelloTest, BasicAssertions) {
    // Expect two strings not to be equal.
    EXPECT_STRNE("hello", "world");
    // Expect equality.
    EXPECT_EQ(7 * 6, 42);
}