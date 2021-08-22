#include <iostream>
#include <vector>
#include <string>
#include "animal.h"
#include "warrior.h"

int main() {
    // std::cout << "Hello World" << std::endl;

    // Animal fred;
    // fred.set_height(33);
    // fred.set_weight(10);
    // fred.set_name("Fred");
    // fred.to_string();

    // Animal tom("Tom", 36, 15);
    // tom.to_string();

    // Dog spot("Spot", 38, 16, "Woooooof");
    // spot.to_string();

    // std::cout << "Number of Animals: " << Animal::get_num_of_animals() << "\n";

    // srand(time(NULL));
    // Warrior thor("Thor", 100, 30, 15);
    // Warrior hulk("hulk", 155, 25, 10);

    // Battle::start_fight(thor, hulk);

    Animal an = Dog("spot", 10, 5, "Woooof");
    an.to_string();

    return 0;
}