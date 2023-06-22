// introduction-variable-sized-arrays/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <bits/stdc++.h>
#include <iostream>
#include <vector>

using namespace std;

int main() {

    unsigned long n, q;
    cin >> n >> q;

    vector<vector<unsigned long>> a;
    for (unsigned long ni = 0; ni < n; ++ni) {
        unsigned long k;
        cin >> k;

        a.push_back(vector<unsigned long>());
        for (unsigned long ki = 0; ki < k; ++ki) {
            unsigned long tmp;
            cin >> tmp;
            a[ni].push_back(tmp);
        }
    }

    for (unsigned long qi = 0; qi < q; ++qi) {
        unsigned long i, j;
        cin >> i >> j;
        cout << a[i][j] << endl;
    }

    return 0;
}
