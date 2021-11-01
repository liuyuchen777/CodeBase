//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"
#include <tuple>

using std::tuple;

int main() {
    std::cout << "---------------Tuple-------------" << std::endl;

    /*
     * tuple is able to store any number, type of data
     * when function need to return multiple value, can store this data in a tuple and return it
     */
    {
        std::tuple<int, char> first;                             // 1. first{}
        std::tuple<int, char> second(first);                     // 2. second{}
        std::tuple<int, char> third(std::make_tuple(20, 'b'));   // 3. third{20,'b'}
        std::tuple<long, char> fourth(third);                    // 4. fourth{20,'b'}
        std::tuple<int, char> fifth(10, 'a');                    // 5. fifth{10.'a'}
        std::tuple<int, char> sixth(std::make_pair(30, 'c'));    // 6. sixth{30,''c}
    }

    // make tuple
    {
        auto first = std::make_tuple(10, 'a'); // tuple<int, char>
        const int a = 0;
        int b[3];
        auto second = std::make_tuple(a, b); // tuple<int, int *>
    }

    return 0;
}