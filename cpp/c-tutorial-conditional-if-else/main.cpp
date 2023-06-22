// c-tutorial-conditional-if-else/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <iostream>
#include <bits/stdc++.h>

using namespace std;

int main() {
    int n;
    cin >> n;

    if (n > 9) {
        cout << "Greater than 9" << endl;
    } else if (n >= 1 && n <= 9) {
        static const char *word[] = {
            "one", "two", "three", "four", "five",
            "six", "seven", "eight", "nine"
        };
        cout << word[n-1] << endl;
    } else {
        cerr << "Invalid input: " << n << endl;
    }

    return 0;
}
