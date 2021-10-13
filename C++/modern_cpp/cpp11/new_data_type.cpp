//
// Created by Yuchen Liu on 2021/07/14.
//

#include "c++11_test.h"

/*  C++11 Integer Type
    整数类型	            ｜等价类型	                              C++11标准规定占用最少位数
    -----------------------------------------------------------------------------
    short	            ｜short int                              ｜至少 16 位（2 个字节）
    signed short        ｜                                       ｜
    --------------------｜---------------------------------------｜
    signed short int    ｜                                       ｜
    unsigned short	    ｜unsigned short int                     ｜
    unsigned short int  ｜                                       ｜
    ------------------------------------------------------------------------------
    int	                ｜int                                    ｜至少 16 位（2 个字节）
    signed              ｜                                       ｜
    --------------------｜---------------------------------------｜
    signed int          ｜                                       ｜
    unsigned	        ｜unsigned int                           ｜
    unsigned int        ｜                                       ｜
    ------------------------------------------------------------------------------
    long	            ｜long int                               ｜至少 32 位（4 个字节）
    long int            ｜                                       ｜
    signed long         ｜                                       ｜
    --------------------｜---------------------------------------｜
    signed long int     ｜                                       ｜
    unsigned long	    ｜unsigned long int                      ｜
    unsigned long int   ｜                                       ｜
    -------------------------------------------------------------------------------
    (Below is all C++11 New)
    long long           ｜long long int                          ｜至少 64 位（8 个字节）
    long long int       ｜                                       ｜
    signed long long    ｜                                       ｜
    --------------------｜---------------------------------------｜
    signed long long int｜                                       ｜
    unsigned long long  ｜unsigned long long int                 ｜
    unsigned long long  ｜                                       ｜
    int                 ｜                                       ｜
 */

int main() {
    cout <<"long long maximum: " << LLONG_MIN <<" "<< hex << LLONG_MIN <<"\n";
    cout << dec <<"long long minimum: " << LLONG_MAX << " " << hex << LLONG_MAX << "\n";
    cout << dec << "unsigned long long maximum: " << ULLONG_MAX << " " << hex << ULLONG_MAX;

    return 0;
}