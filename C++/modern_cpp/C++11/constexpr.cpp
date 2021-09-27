//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"

/*
 * constexpr to decorate function
 */
constexpr int display1(int x) {
    int ret = 1 + 2 + x;
    return ret;
}
constexpr int display2(int x) {
    return 1 + 2 + x;
}
//constexpr void display3() {
//
//}
// error: return value couldn't be void
int num = 3;
//constexpr int display3(int x) {
//    return num + x;
//}
// error: non-constant variable is not allowed in constexpr function

/*
 * constexpr in struct and class
 */
struct MyType {
    constexpr MyType(char *name, int age) : name(name), age(age) {};
    const char *name;
    int age;
};

// constexpr is not able to decorate self-define type

class MyType2 {
public:
    constexpr MyType2(const char *name, int age) : name(name), age(age) {};
    constexpr const char * get_name() {
        return name;
    }
    constexpr int get_age() {
        return age;
    }
private:
    const char *name;
    int age;
};

/*
 * constexpr in template
 */
// constexpr can decorate template but is not decided to be const
template<typename T>
constexpr T display(T t) {
    return t;
}

int main() {
    {
        /*
         * constexpr to decorate variable
         */
        constexpr int num = 1 + 2 + 3;
        int url[num] = {1, 2, 3, 4, 5};
        std::cout << url[1] << std::endl;
    }
    {
        // use constexpr
        int a[display2(3)] = {1, 2, 3, 4};
        // constexpr struct
        constexpr struct MyType mt {"liuyuchen", 10};
        std::cout << mt.name << "  " << mt.age << std::endl;
        // constexpr class
        constexpr MyType2 mt2 {"liuyuchen", 22};
        const char *tmp_name = mt2.get_name();
        constexpr int tmp_age = mt2.get_age();
        std::cout << tmp_name << "  " << tmp_age << std::endl;
        // constexpr template
        // normal function
        struct MyType ret1 = display(mt);
        std::cout << ret1.name << "  " << ret1.age << std::endl;
        // constexpr function
        constexpr int ret2 = display(10);
        std::cout << ret2 << std::endl;
    }

    return 0;
}