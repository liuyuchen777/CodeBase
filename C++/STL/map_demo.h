/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 10:52:04
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 12:58:23
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/map_demo.h
 * @GitHub: https://github.com/liuyuchen777
 */

#include <iostream>
#include <map>

using namespace std;

template<typename key, typename value>
ostream& operator<<(ostream& os, map<key, value> &m);

void map_basic_operations();