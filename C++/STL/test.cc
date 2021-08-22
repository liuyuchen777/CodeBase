/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 00:05:13
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 14:19:24
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/test.cc
 * @GitHub: https://github.com/liuyuchen777
 */

#include <gtest/gtest.h>
#include "vector_demo.h"
#include "set_demo.h"
#include "map_demo.h"
#include "unordered_map_demo.h"
#include "unordered_set_demo.h"

TEST(basic, vector) {
    vector_initialization();
    vector_iterator();
    vector_capacity();
    vector_element_access();
    vector_modifier();
}

TEST(basic, set) {
    set_modifier();
    set_operations();
    set_observations();
}

TEST(basic, map) {
    map_basic_operations();
}

TEST(basic, unorder_set) {
    unordered_set_initilizer();
    unordered_set_bucket();
    unordered_set_policy();
    unordered_set_observor();
}