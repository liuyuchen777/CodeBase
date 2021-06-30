/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-30 19:57:07
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-30 21:06:14
 * @Description: 
 * @FilePath: /CodeBase/C++/Temp/test.cc
 * @GitHub: https://github.com/liuyuchen777
 */
#include <vector>
#include <math.h>
#include <iostream>

using namespace std;

int ceiling(int val) {
    if (val % 2 == 0) {
        return val / 2;
    } else {
        return val / 2 + 1;
    }
}

void combine(const int target, vector<int> &res) {
    
}

vector<int> baseNeg2(int N) {
    vector<int> res;

    if (N  == 0)
        return res;
    
    while (N != 0)
    {
        int r = N % -2;
        N /= -2;
        if (r < 0) {
            r += 2;
            ++N;
        }
        res.push_back(r);
    }
    
    return res;
}

vector<int> solution(vector<int> &vec) {
    vector<int> res;
    // extreme condition
    if (vec.size() == 0) {
        return res;
    }
    // get val of vec
    int val = 0;
    for (int i = 0; i < vec.size(); i++) {
        val += vec[i] * (int)pow(-2, i);
    }
    // ceiling
    val = ceiling(val);
    // make up res
    return baseNeg2(val);
}

vector<int> solution2(vector<int> &vec) {
    vector<int> res(vec.size());
    vector<int> return_res;

    for(int i = 0; i < vec.size() - 1; i++) {
        if(vec[vec.size() - i - 1] == 1) {
            res[i]++;
            res[i + 1]++;
        }
    }

    if(vec[0] == 1) {
        res[vec.size() - 1]++;
    }

    for(int i = vec.size() - 1; i >= 0; --i) {
        if(res[i] >= 2) {
            if(res[i-1] >= 1) {
                res[i - 1]--;
                res[i] -= 2;
            } else {
                res[i - 1]++;
                res[i - 2]++;
            }
        }
    }

    // get result
    int k = 0;
    while (res[k] == 0) {
        k++;
    }
    for (int i = res.size() - 1; i >= k; i--) {
        return_res.push_back(res[i]);
    }

    return return_res;
}

int main() {
    vector<int> vec1 = {1, 0, 0, 1, 1, 1};
    vector<int> res1 = solution2(vec1);
    for (auto i : res1) {
        cout << i;
    }
    cout << endl;

    vector<int> vec2 = {1, 0, 0, 1, 1};
    vector<int> res2 = solution2(vec2);
    for (auto i : res2) {
        cout << i;
    }
    cout << endl;

    // solution1
    vector<int> res3 = solution(vec1);
    for (auto i : res3) {
        cout << i;
    }
    cout << endl;

    vector<int> res4 = solution(vec2);
    for (auto i : res4) {
        cout << i;
    }
    cout << endl;
    
}