#ifndef _ANIMAL_
#define _ANIMAL_

#include <iostream>

class Animal {
private:
    std::string name;
    double height;
    double weight;
    // static variable
    static int num_of_animals;
public:
    std::string get_name() { return this->name; }
    void set_name(std::string name) { this->name = name; }
    double get_height() { return this->height; }
    void set_height(double height) { this->height = height; }
    double get_weight() { return this->weight; }
    void set_weight(double weight) { this->weight = weight; }

    void set_all(std::string, double, double);
    
    Animal(std::string, double, double);
    Animal();

    ~Animal(); 
    // static method only able to access static member variable
    static int get_num_of_animals() { return num_of_animals; }
    
    void to_string();
};

int Animal::num_of_animals = 0;

void Animal::set_all(std::string name, double height, double weight) {
    this->name = name;
    this->weight = weight;
    this->height = height;
}

Animal::Animal(std::string name, double height, double weight) {
    this->name = name;
    this->weight = weight;
    this->height = height;
    Animal::num_of_animals++;
}

Animal::Animal() {
    this->name = "";
    this->weight = 0;
    this->height = 0;
    Animal::num_of_animals++;
}

Animal::~Animal() {
    std::cout << "Animal " << this->name << " destroyed!\n";
    Animal::num_of_animals--;
}

void Animal::to_string() {
    std::cout << this->name << " is " << this->height << 
        " cms tall and " << this->weight << " kgs in weight.\n";
}

class Dog: public Animal {
private:
    std::string sound = "Woof";
public:
    void make_sound() {
        std::cout << "The dog " << this->get_name() << " say " << this->sound << "\n";
    }
    Dog(std::string, double, double, std::string);
    Dog(): Animal() {};
    void to_string();
};

Dog::Dog(std::string name, double height, double weight, std::string sound): Animal(name, height, weight){
    this->sound = sound;
}

void Dog::to_string() {
    std::cout << get_name() << " is " << Animal::get_height() << 
        " cms tall and " << Animal::get_weight() << " kgs in weight and says " 
        << this->sound << "\n";
}

#endif