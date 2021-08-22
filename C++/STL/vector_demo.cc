/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 00:06:25
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 11:31:18
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/vector_demo.cc
 * @GitHub: https://github.com/liuyuchen777
 */
#include "vector_demo.h"

ostream& operator<<(ostream &os, vector<int> &vec) {
    for (auto it = vec.begin(); it != vec.end(); it++) {
        os << *it << "  ";
    }
    
    return os;
}

void vector_initialization() {
    cout << "------------Initialization of STL Vector----------" << endl;
    // 1. initialize
    // 1.1 initialize by pushing values one by one
    {
        vector<int> vec;
        vec.push_back(10);
        vec.push_back(20);
        vec.push_back(30);
        cout << "1.1: " << vec << endl;
    }
    // 1.2 specify size and initialize all value
    {
        vector<int> vec(3, 10);
        cout << "1.2: " << vec << endl;
    }
    // 1.3 initialize like arrays, c++11
    {
        vector<int> vec{ 10, 20, 30 };
        // vector<int> vec = {10, 20, 30};
        cout << "1.3: " << vec << endl;
    }
    // 1.4 initializeing from an array
    {
        int arr[] = {10, 20, 30};
        int n = sizeof(arr) / sizeof(arr[0]);

        vector<int> vec(arr, arr + n);
        
        cout << "1.4: " << vec << endl;
    }
    // 1.5 initialzing from other vector
    {
        vector<int> vec1{ 10, 20, 30};
        vector<int> vec2(vec1.begin(), vec1.end());
        // vector<int> vec2(vec1);
        
        cout << "1.5: " << vec2 << endl;
    }
    // 1.6 initialize all elements with particular value
    {
        vector<int> vec(5);
        int val = 5;
        fill(vec.begin(), vec.end(), val);
        
        cout << "1.6: " << vec << endl;
    }
    // 1.7 =
    {
        vector<int> v1{1, 2, 3, 4, 5};
        vector<int> v2;
        v2 = v1;
        cout << "1.7: " << v2 << endl;
    }
}

void vector_iterator() {
    cout << "------------Iterator of STL Vector----------" << endl;
    /** cheat sheet
     * begin(), end()
     * rbegin(), rend()
     * cbegin(), cend()
     * crbegin(), crend()
     */
    vector<int> g1;

    for (int i = 1; i <= 5; i++)
        g1.push_back(i);

    cout << "Output of begin and end: ";
    for (auto i = g1.begin(); i != g1.end(); ++i)
        cout << *i << " ";

    cout << "\nOutput of cbegin and cend: ";
    for (auto i = g1.cbegin(); i != g1.cend(); ++i) {
        cout << *i << " ";
        // *i = 10; // not allow
    }

    cout << "\nOutput of rbegin and rend: ";
    for (auto ir = g1.rbegin(); ir != g1.rend(); ++ir)
        cout << *ir << " ";

    cout << "\nOutput of crbegin and crend : ";
    for (auto ir = g1.crbegin(); ir != g1.crend(); ++ir)
        cout << *ir << " ";

    cout << endl;
}

void vector_capacity() {
    cout << "------------Capacity of STL Vector----------" << endl;
    /** cheat sheet
     * size()
     * max_size() - return maximum numbers this vector can hold
     * capacity()
     * resize(n) - resize containers to have n element
     * empty() - return bool
     * shrink_to_fit() - reduce capacity to fit size
     * reverse() - request vector capacity at least enough for n elements
     */
    vector<int> g1;

    for (int i = 1; i <= 5; i++)
        g1.push_back(i);

    cout << "Size : " << g1.size() << endl;
    cout << "Capacity : " << g1.capacity() << endl;
    cout << "Max_Size : " << g1.max_size() << endl;

    // resizes the vector size to 4
    g1.resize(4);

    // prints the vector size after resize()
    cout << "Size : " << g1.size() << endl;

    // checks if the vector is empty or not
    if (g1.empty() == false)
        cout << "\nVector is not empty";
    else
        cout << "\nVector is empty";

    // Shrinks the vector
    g1.shrink_to_fit();
    cout << "\nVector elements are: ";
    for (auto it = g1.begin(); it != g1.end(); it++)
        cout << *it << " ";
    cout << endl << "capacity: " << g1.capacity() << endl;

    // reverse
    g1.reserve(10);
    cout << "capacity: " << g1.capacity() << endl;
}  

void vector_element_access() {
    cout << "------------Element Access of STL Vector----------" << endl;
    /** cheat sheet
     * []
     * at(n)
     * front()
     * back()
     * data() - return a direct pointer to the memory array used internally by the vector
     */
    vector<int> g1;

    for (int i = 1; i <= 10; i++)
        g1.push_back(i * 10);

    cout << "\nReference operator [g] : g1[2] = " << g1[2];

    cout << "\nat : g1.at(4) = " << g1.at(4);

    cout << "\nfront() : g1.front() = " << g1.front();

    cout << "\nback() : g1.back() = " << g1.back();

    // pointer to the first element
    int* pos = g1.data();

    cout << "\nThe first element is " << *pos << endl;
}

void vector_modifier() {
    cout << "------------Modifier of STL Vector----------" << endl;
    /** cheat sheet
     * assign(pos, val) - assgin new val to pos
     * push_back()
     * pop_back()
     * insert()
     * earse(it)
     * swap()
     * clear() - remove all elements in a vector
     * emplace(it, val)
     * emplace_back(val)
     * emplace_front(val)
     */
    // Assign vector
    vector<int> v;

    // fill the array with 10 five times
    v.assign(5, 10);

    cout << "The vector elements are: ";
    cout << v << endl;

    // inserts 15 to the last position
    v.push_back(15);
    cout << "The last element is: " << v[v.size() - 1] << endl;

    // removes last element
    v.pop_back();
    cout << "The last element is: " << v[v.size() - 1] << endl;

    // prints the vector
    cout << "The vector elements are: " << v << endl;

    // inserts 5 at the beginning
    v.insert(v.begin(), 5);

    cout << "Afetr inserting at begin: " << v << endl;

    // removes the first element
    v.erase(v.begin());

    cout << "After earse: " << v << endl;

    // inserts at the beginning
    v.emplace(v.begin(), 5);
    cout << "The first element is: " << v[0] << endl;

    // Inserts 20 at the end
    v.emplace_back(20);
    cout << "The last element is: " << v[v.size() - 1] << endl;

    // erases the vector
    v.clear();
    cout << "Vector size after erase(): " << v.size() << endl;

    // two vector to perform swap
    vector<int> v1, v2;
    v1.push_back(1);
    v1.push_back(2);
    v2.push_back(3);
    v2.push_back(4);

    cout << "Vector 1: ";
    cout << v1 << endl;

    cout << "Vector 2: ";
    cout << v2 << endl;

    // Swaps v1 and v2
    v1.swap(v2);

    cout << "After Swap: \nVector 1: ";
    cout << v1 << endl;

    cout << "Vector 2: ";
    cout << v2 << endl;

    /*
    emplace_back() and push_back():
    push_back in the above case will create a temporary object and move it into the 
    container. However, in-place construction used for emplace_back would be more performant 
    than constructing and then moving the object (which generally involves some copying).
     */
}