/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 14:02:32
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 14:35:43
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/unordered_set_demo.cc
 * @GitHub: https://github.com/liuyuchen777
 */
#include "unordered_set_demo.h"

template<typename T>
ostream& operator<<(ostream &os, unordered_set<T> my_set) {
    for (auto it = my_set.begin(); it != my_set.end(); it++) {
        cout << *it << "  ";
    }

    return os;
}

void unordered_set_initilizer() {
    cout << "----------unordered set initializer-------" << endl;
    unordered_set<string> first;                                // empty
    unordered_set<string> second ( {"red","green","blue"} );    // init list
    unordered_set<string> third ( {"orange","pink","yellow"} ); // init list
    unordered_set<string> fourth ( second );                    // copy
    unordered_set<string> fifth ( fourth.begin(), fourth.end() ); // range
    cout << "1. " << first << endl;
    cout << "2. " << second << endl;
    cout << "3. " << third << endl;
    cout << "4. " << fourth << endl;
    cout << "5. " << fifth << endl;
}

void unordered_set_bucket() {
    cout << "--------unordered set bucket--------" << endl;
    
}

void unordered_set_policy() {
    
}

void unordered_set_observor() {

}