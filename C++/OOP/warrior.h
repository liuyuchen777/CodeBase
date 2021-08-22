#ifndef _WARRIER_
#define _WARRIER_

#include <iostream>
#include <math.h>

class Warrior {
private:
    int attack_max;
    int block_max;
public:
    int health;
    std::string name;

    Warrior(std::string name, int health, int attack_max, int block_max)
        : name(name), health(health), attack_max(attack_max), block_max(block_max) {
        // constructor
    }

    int attack() {
        return std::rand() % this->attack_max;
    }

    int block() {
        return std::rand() % this->block_max;
    }
};

class Battle {
public:
    static void start_fight(Warrior &w1, Warrior &w2) {
        while (true) {
            // warrior 1 attack warrior 2
            if (Battle::get_attack_result(w1, w2).compare("Game Over") == 0) {
                std::cout << "Game Over!\n";
                break;
            }
            // warrior 2 attack warror 1
            if (Battle::get_attack_result(w2, w1).compare("Game Over") == 0) {
                std::cout << "Game Over!\n";
                break;
            }
        }
    }

    static std::string get_attack_result(Warrior &w1, Warrior &w2) {
        int w1_attack = w1.attack();
        int w2_block = w2.block();
        int damage = ceil(w1_attack - w2_block);
        damage = (damage <= 0) ? 0 : damage;
        w2.health = ((w2.health - damage) <= 0) ? 0 : (w2.health - damage);
        // print out info
        printf("%s attacks %s and deals %d damage.\n", 
            w1.name.c_str(), w2.name.c_str(), damage);
        printf("%s is down to %d health.\n", w2.name.c_str(), w2.health);
        if (w2.health == 0) {
            printf("%s has died and %s is winner.\n", w2.name.c_str(), w1.name.c_str());
            return "Game Over";
        } else {
            return "Fight Again";
        }
    }
};

#endif