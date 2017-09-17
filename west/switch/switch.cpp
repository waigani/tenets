#include <iostream>

int main() {
    int a = 20;

    switch (a) {
        case 10:
        case 20:
            std::cout << "20 called" << std::endl;
        case 40:
            std::cout << "40 called" << std::endl;
            goto endswitch;
        case 60:
            std::cout << "60 called" << std::endl;
            return 0;
        case 80:
            std::cout << "80 called" << std::endl;
        case default:
            std::cout << "default called" << std::endl;
            break;
    }
    endswitch:
    return 0;
}