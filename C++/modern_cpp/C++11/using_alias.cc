//
// Created by Yuchen Liu on 2021/07/10.
//

#include "c++11_test.h"

// typedef could be to define a alias
typedef unsigned int uint_t;

void func(unsigned int a);
void func(uint_t a);

template <typename Val>
using str_map_t = std::map<std::string, Val>;

// using is same as typedef
typedef unsigned int uint32_t_1;
using uint32_t_2 = unsigned int;


int main() {
    cout << "--------------------using alias!----------------" <<endl;
    str_map_t<int> map1;


}