/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 10:14:57
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 10:27:37
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/set_demo.h
 * @GitHub: https://github.com/liuyuchen777
 */
#include <iostream>
#include <set>

using namespace std;

template<typename T>
ostream& operator<<(ostream& os, set<T> s);

void set_modifier();