// other-concepts-bitset-1/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <algorithm>
#include <bits/stdc++.h>
#include <cmath>
#include <cstdio>
#include <iostream>

using namespace std;

int main() {
    const uint_fast64_t M = 1UL << 31;
    uint_fast64_t N, S, P, Q;
    cin >> N >> S >> P >> Q;
    uint_fast64_t curr = M, prev = S % M, i = 1;

    for (; i <= N - 1; i++, prev = curr) {
        curr = ((uint_fast64_t) prev * P + Q) % M;
		if (curr == prev)
			break;
	}

    cout << i << endl;
    return 0;
}
