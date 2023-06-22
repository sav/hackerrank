// stl-prettyprint/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <algorithm>
#include <bits/stdc++.h>
#include <cmath>
#include <cstdio>
#include <ios>
#include <iostream>
#include <iomanip>

int main() {
	int T;
	double A, B, C;
	std::ios_base::fmtflags f(std::cout.flags());

	std::cin >> T;
	while(T--) {
		std::cin >> A >> B >> C;

		std::cout << std::left << std::showbase << std::nouppercase << std::hex
				  << static_cast<long long>(trunc(A)) << std::endl;
		std::cout.flags(f);

		std::cout << std::setprecision(2) << std::setw(15)
				  << std::right << std::showpos << std::fixed
				  << std::setfill('_') << B << std::endl;
		std::cout.flags(f);

		std::cout << std::setprecision(9) << std::noshowpos << std::uppercase
				  << std::scientific << C << std::endl;
   		std::cout.flags(f);
	}

	return 0;
}
