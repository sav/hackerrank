// c-tutorial-functions/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <iostream>
#include <algorithm>

int max_of_four(int a, int b, int c, int d) {
	return std::max(std::max(std::max(a, b), c), d);
}

int main() {
	int a, b, c, d;

	std::cin >> a;
	std::cin >> b;
	std::cin >> c;
	std::cin >> d;

	std::cout << max_of_four(a, b, c, d) << std::endl;

	return 0;
}
