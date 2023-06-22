// c-tutorial-pointer/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <iostream>
#include <cstdlib>
#include <stdlib.h>

void update(int *a, int *b) {
	int x = *a + *b;
	int y = ::abs(*a - *b);

	*a = x;
	*b = y;
}

int main() {
	int a, b;

	std::cin >> a;
	std::cin >> b;

	update(&a, &b);

	std::cout << a << std::endl;
	std::cout << b << std::endl;

	return 0;
}
