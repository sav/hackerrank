// c-tutorial-for-loop/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <iostream>
#include <bits/stdc++.h>

using namespace std;

const char *word[] = {
	"one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine"
};

int main(void) {
    int a, b;

    cin >> a;
	cin >> b;

	for (int i = a; i <= b; ++i) {
		if (i >= 1 && i <= 9) {
			cout << word[i - 1] <<  endl;
		} else if (i % 2 == 0) {
			cout << "even" << endl;
		} else {
			cout << "odd" << endl;
		}
	}

    return 0;
}
