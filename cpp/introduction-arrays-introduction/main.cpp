// introduction-arrays-introduction/main.cpp
// Copyright © 2023  Savio Sena <savio.sena@acm.org>

#include <bits/stdc++.h>
#include <iostream>

using namespace std;

int main() {

  int n;

  cin >> n;

  int i = n;
  std::vector<int> a;

  while (i--) {
    int t;
    scanf("%d", &t);
    a.push_back(t);
  }

  bool space = false;
  std::for_each(a.rbegin(), a.rend(), [&space](const int n) {
	  if (space) std::cout << ' ';
	  std::cout << n;
	  space = true;
  });
  std::cout << std::endl;

  return 0;
}
