/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 14:02:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 14:23:20
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/unordered_set_demo.h
 * @GitHub: https://github.com/liuyuchen777
 */
#include <iostream>
#include <unordered_set>
#include <string>

using namespace std;

template<typename T>
ostream& operator<<(ostream &os, unordered_set<T> my_set);

void unordered_set_initilizer();

void unordered_set_bucket();

void unordered_set_policy();

void unordered_set_observor();