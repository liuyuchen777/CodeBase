//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"

template <typename R = int, typename T>  // error in C++98/03: default template arguments
R func(T val)
{
    return val;
}

int main() {
    std::cout << "--------------------Default Type of Template---------------" << std::endl;
    func(97); // R=int, U=int
    func<char>(97); // R=char, U=int
    func<double, int>(97); // R=double, U=int

    return 0;
}