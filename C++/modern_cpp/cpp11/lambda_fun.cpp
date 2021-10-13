//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"
#include <algorithm>

/*
 * lambda grammar:
 * [outside variable] (parameter) mutable noexcept/throw() -> return type {
 *      function body;
 * };
 */

/*
    外部变量格式	        功能
    []	                空方括号表示当前 lambda 匿名函数中不导入任何外部变量。
    [=]	                只有一个 = 等号，表示以值传递的方式导入所有外部变量；
    [&]	                只有一个 & 符号，表示以引用传递的方式导入所有外部变量；
    [val1,val2,...]	    表示以值传递的方式导入 val1、val2 等指定的外部变量，同时多个变量之间没有先后次序；
    [&val1,&val2,...]	表示以引用传递的方式导入 val1、val2等指定的外部变量，多个变量之间没有前后次序；
    [val,&val2,...]	    以上 2 种方式还可以混合使用，变量之间没有前后次序。
    [=,&val1,...]	    表示除 val1 以引用传递的方式导入外，其它外部变量都以值传递的方式导入。
    [this]	            表示以值传递的方式导入当前的 this 指针。
 */

int main() {
    std::cout << "---------------lambda-------------" << std::endl;

    int num[] = {4, 2, 3, 1};

    std::sort(num, num + 4, [=](int x, int y) -> bool {
        return x < y;
    });

    for (auto n : num) {
        std::cout << n << " ";
    }
    std::cout << std::endl;

    auto display = [](int a, int b) -> void {
        cout << a << " " << b << endl;
    };
    display(10, 20);

    return 0;
}