#include <iostream>

/*
C++ Abstract Class (Interface)
If a class has at least 1 pure virtual function, this class is abstract class
pure virtual function is defined through "=0"
*/

// base class
class Shape {
public:
    virtual double get_area() = 0;
    void set_width(int w) { this->width = w; }
    void set_height(int h) { this->height = h; }
    Shape(int w, int h): width(w), height(h) {}
protected:
    double width;
    double height;
};

// derived class
class Rectangle: public Shape {
public:
    double get_area() { return (width * height); }
    Rectangle(int w, int h): Shape(w, h){}
};

class Triangle: public Shape {
public:
    double get_area() { return (width * height) / 2; }
    Triangle(int w, int h): Shape(w, h){}
};

// error
class Hexangle: public Shape {
public:
    Hexangle(int r): Shape(0, 0) { this->radius = r; }
private:
    int radius;
};

/*
global / local * static / non-static

*/

int main() {

    Rectangle rect(10, 20);
    Triangle tri(8, 10);

    std::cout << rect.get_area() << std::endl;
    std::cout << tri.get_area() << std::endl;

    // Hexangle hex;

    return 0;
}