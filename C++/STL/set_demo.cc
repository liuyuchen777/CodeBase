/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 10:14:52
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 11:57:59
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/set_demo.cc
 * @GitHub: https://github.com/liuyuchen777
 */

#include "set_demo.h"

template<typename T>
ostream& operator<<(ostream &os, set<T> &s) {
    for (auto it = s.begin(); it != s.end(); it++) {
        os << *it << "  ";
    }
    
    return os;
}

void set_modifier() {
    /** cheat sheet
     * insert()
     * earse()
     * swap()
     * clear()
     * emplace() - construct and insert element
     * emplcae_hint()
     */

    // insert - insert is after
    {
        cout << "-------------set insert-----------" << endl;
        set<int> my_set;
        set<int>::iterator it;
        pair<set<int>::iterator, bool> ret;
        
        for (int i = 1; i <= 5; i++)
            my_set.insert(i * 10);
        
        cout << "origin set: " << my_set << endl;
        
        ret = my_set.insert(20); // insert nothing

        cout << "after insert a duplicate element: " << my_set << endl;

        if (ret.second == false) it = ret.first; // it now pints to element 20
        
        my_set.insert(it, 25); // efficient insert
        my_set.insert(it, 26);

        cout << "after insert two element: " << my_set << endl;

        int arrays[] = {5, 10, 100};
        my_set.insert(arrays, arrays + 3);

        cout << "after insert an array" << my_set << endl;
    }

    // earse
    // earse return iterator of the next element
    {
        cout << "---------------set earse------------" << endl;
        set<int> my_set;
        set<int>::iterator it;

        // insert some values:
        for (int i=1; i<10; i++) my_set.insert(i*10);  // 10 20 30 40 50 60 70 80 90

        cout << "origin set: " << my_set << endl;

        it = my_set.begin();
        ++it;                                         // "it" points now to 20

        my_set.erase (it);

        cout << "after earse: " << my_set << endl;

        my_set.erase (40);

        cout << "after earse: " << my_set << endl;

        it = my_set.find(60);
        my_set.erase (it, my_set.end());

        cout << "my_set contains: " << my_set << endl;
    }
    
    // swap - swap two contents of set

    // clear - clear all elements

    // emplace
    {
        cout << "----------------set emplace--------------" << endl;
        set<string> my_set;

        my_set.emplace("foo");
        my_set.emplace("bar");
        auto ret = my_set.emplace("foo");

        if (!ret.second) cout << "foo already exists in my_set\n";
    }

    // emplace_hint - hint for where the element could be insert, improve efficiency
    
}

void set_operations() {
    /** cheat sheet
     * find()
     * count() 
     * lower_bound()
     * upper_bound()
     * equal_range()
     */
    
    // find - return iterator of
    // if not found, return set::end
    {
        cout << "--------------set find-----------" << endl;
        set<int> my_set;
        set<int>::iterator it;

        // set some initial values:
        for (int i = 1; i <= 5; i++) 
            my_set.insert(i * 10);    // set: 10 20 30 40 50

        it = my_set.find(20);
        my_set.erase (it);
        my_set.erase (my_set.find(40));

        cout << "my_set: " << my_set << endl;
    }

    // count -> return 1 if have this val, return 0 if val doesn't exist
    
    // upper_bound, lower_bound
    {
        cout << "-------------bound-----------" << endl;
        set<int> my_set;
        set<int>::iterator itlow, itup;

        for (int i = 1; i < 10; i++) 
            my_set.insert(i * 10); // 10 20 30 40 50 60 70 80 90

        itlow = my_set.lower_bound (32);          
        itup = my_set.upper_bound (58);                 

        my_set.erase(itlow, itup);                     // 10 20 70 80 90

        cout << "changed set: " << my_set << endl;
    }

    // equal_range - return bound of range taht includes all the elements val
    {
        cout << "--------------set equal_range-----------" << endl;
        set<int> my_set;

        for (int i=1; i<=5; i++) my_set.insert(i * 10);   // my_set: 10 20 30 40 50

        pair<set<int>::const_iterator, set<int>::const_iterator> ret;
        ret = my_set.equal_range(30);

        cout << "the lower bound points to: " << *ret.first << '\n';
        cout << "the upper bound points to: " << *ret.second << '\n';
    }
}

void set_observations() {
    // key_comp and value_comp return comparsion ov=bject used by the container
    {
        cout << "--------------self define set struct compare---------" << endl;
        struct my_comp {
            // add const at end of function declartion
            bool operator()(const int &a, const int &b) const {
                return a > b; // decrease order
                // return a < b; // increase order
            }
        };
        
        set<int, my_comp> s;
        for (int i = 1; i < 6; i++) {
            s.insert(i * i);
        }

        for(auto it = s.begin(); it != s.end(); it++)
        {
            cout << *it << " ";
        }
        
        cout << endl;
    }

    {
        cout << "----------self define set class compare-------" << endl;
        struct Info {
            string name;
            float score;

            bool operator<(const Info &a) const {
                // return this->score > a.score; // decrease order
                return this->score < a.score; // increase order
            }
        };

        set<Info> s;
        Info info{.name = "liuyuchen", .score = 100.0};
        s.insert(info);
        info = {.name = "zhaozihao", .score = 90.0};
        s.insert(info);
        info = {.name = "xiongshuangli", .score = 95.0};
        s.insert(info);
        
        for (auto it = s.begin(); it != s.end(); it++) {
            cout << it->name << " : " << it->score << endl;
        }
    }
    
    // get compare
    {
        cout << "--------------set get comapre-----------" << endl;
        set<int> my_set;
        int highest;

        set<int>::key_compare mycomp = my_set.key_comp(); // just <

        for (int i = 0; i <= 5; i++) 
            my_set.insert(i); // set: 1 2 3 4 5 

        cout << "my_set contains: ";

        highest = *my_set.rbegin();
        set<int>::iterator it = my_set.begin();
        do {
            cout << ' ' << *it;
        } while (mycomp(*(++it), highest));

        cout << endl;
    }
}