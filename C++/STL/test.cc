/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 00:05:13
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 10:27:55
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/test.cc
 * @GitHub: https://github.com/liuyuchen777
 */

#include <gtest/gtest.h>
#include "vector_demo.h"
#include "set_demo.h"

TEST(basic, vector) {
    vector_initialization();
    vector_iterator();
    vector_capacity();
    vector_element_access();
    vector_modifier();
}

TEST(basic, set) {
    set_modifier();
}