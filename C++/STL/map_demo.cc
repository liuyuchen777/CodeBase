/*
 * @Author: Liu Yuchen
 * @Date: 2021-07-01 10:51:59
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-07-01 14:06:27
 * @Description: 
 * @FilePath: /CodeBase/C++/STL/map_demo.cc
 * @GitHub: https://github.com/liuyuchen777
 */
#include "map_demo.h"

template<typename key, typename value>
ostream& operator<<(ostream& os, map<key, value> &m) {
    for (auto it = m.begin(); it != m.end(); it++) {
        cout << it->first << " => " << it->second << endl;
    }

    return os;
}

typedef struct UrlKey {
    uint64_t id;
    uint64_t version;
    uint64_t hash;
} UrlKey;

typedef struct UrlValue {
    string url;
} UrlValue;

ostream& operator<<(ostream &os, const UrlKey &key) {
    cout << key.id << " " << key.version << " " << key.hash;
    return os;
}

ostream& operator<<(ostream &os, const UrlValue &val) {
    cout << val.url;
    return os;
}

void map_basic_operations() {
    cout << "-----------map initialize---------" << endl;
    {
        // 1. initialize
        // 1.1 default constructor
        map<int, string> my_map;
        my_map[1] = "liuyuchen";
        my_map[2] = "zhaozihao";
        my_map[3] = "xiongshuangli";
        my_map[4] = "wangzihan";
        my_map[5] = "luanhaoxiang";
        cout << "my_map: " << endl << my_map;

        // 1.2 range constructor
        map<int, string> my_map2(my_map.find(2), my_map.find(4));
        cout << "my_map2: " << endl << my_map2;

        // 1.3 copy constructor
        auto my_map3 = my_map2;
        auto my_map4(my_map2);
        cout << "my_map3: " << endl << my_map3;
        cout << "my_map4: " << endl << my_map4;

        // 1.4 move constructor
        auto my_map5(move(my_map2));
        cout << "my_map5: " << endl << my_map5;
        // my_map2[2] = "zhujunyi";
        // cout << "my_map5: " << endl << my_map5;

        // 1.5 list constructor
        map<int, string> my_map6 {
            {1, "liuyuchen"},
            {2, "xiongshuangli"},
            {3, "xiangzihao"}
        };

        cout << "my_map6: " << endl << my_map6;
    }

    cout << "---------------map data access-----------" << endl;
    // 2. data access
    {
        // 2.1 through operator []
        // return value is a reference
        map<char,string> mymap;

        mymap['a'] = "an element";
        mymap['b'] = "another element";
        mymap['c'] = mymap['b'];

        cout << "mymap['a'] is " << mymap['a'] << '\n';
        cout << "mymap['b'] is " << mymap['b'] << '\n';
        cout << "mymap['c'] is " << mymap['c'] << '\n';
        cout << "mymap['d'] is " << mymap['d'] << '\n';

        cout << "mymap now contains " << mymap.size() << " elements.\n";
    }
    
    {
        // 2.2 through at
        map<string, int> my_map = {
            {"alpha", 0},
            {"beta", 0},
            {"gamma", 0}
        };
        my_map.at("alpha") = 10;
        my_map.at("beta") = 20;
        my_map.at("gamma") = 30;
        // my_map.at("miu") = 40;

        cout << "my_map: " << endl << my_map;
    }

    cout << "------------map with custom order----------" << endl;
    // 3. map with custom order
    {
        // 3.1 according to key
        struct cmp_key {
            bool operator()(const UrlKey &k1, const UrlKey &k2) const {
                if (k1.id != k2.id)
                    return k1.id > k2.id;
                if (k1.version != k2.version) 
                    return k1.version > k2.version;
                if (k1.hash != k2.hash)
                    return k1.hash > k2.hash;
                
                return false;
            }
        };

        map<UrlKey, UrlValue, cmp_key> my_map;

        my_map[UrlKey{.id=1, .version=2, .hash=3}] = UrlValue{.url="hello"};
        my_map[UrlKey{.id=1, .version=3, .hash=3}] = UrlValue{.url="world"};
        my_map[UrlKey{.id=1, .version=3, .hash=4}] = UrlValue{.url="are"};
        my_map[UrlKey{.id=2, .version=2, .hash=3}] = UrlValue{.url="you"};

        for (map<UrlKey, UrlValue, cmp_key>::iterator it = my_map.begin(); it != my_map.end(); it++) {
            cout << it->first << " => " << it->second << endl;
        }
    }

    {
        // 3.2 cutom value order

    }

    cout << "--------------map modifier----------" << endl;
    // 4. modifier
    {
        // 4.1 erase - move element in container
        map<char,int> mymap;
        map<char,int>::iterator it;

        // insert some values:
        mymap['a']=10;
        mymap['b']=20;
        mymap['c']=30;
        mymap['d']=40;
        mymap['e']=50;
        mymap['f']=60;

        it=mymap.find('b');
        mymap.erase (it);                   // erasing by iterator

        mymap.erase ('c');                  // erasing by key

        it=mymap.find ('e');
        mymap.erase (it, mymap.end());    // erasing by range

        cout << "mymap: " << endl << mymap;
    }

    {
        // 4.2 insert
        map<char,int> mymap;

        // 4.2.1 first insert function version (single parameter):
        mymap.insert( pair<char,int>('a', 100) );
        mymap.insert( pair<char,int>('z', 200) );

        pair<map<char,int>::iterator, bool> ret;
        ret = mymap.insert (pair<char,int>('z', 500));
        if (ret.second==false) {
            cout << "element 'z' already existed";
            cout << " with a value of " << ret.first->second << endl;
        }

        // 4.2.2 second insert function version (with hint position):
        map<char,int>::iterator it = mymap.begin();
        mymap.insert(it, pair<char,int>('b',300));  // max efficiency inserting
        mymap.insert(it, pair<char,int>('c',400));  // no max efficiency inserting

        // 4.2.3 third insert function version (range insertion):
        map<char,int> anothermap;
        anothermap.insert(mymap.begin(), mymap.find('c'));

        // showing contents:
        cout << "mymap contains: " << endl << mymap;
        cout << "anothermap contains: " << endl << anothermap;
    }
}