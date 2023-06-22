#include <iostream>
#include <string>
#include <sstream>
#include <vector>

std::vector<int> parseInt(std::string s) {
  std::stringstream ss(s);
  std::vector<int> v;
  char c;
  do {
    int n;
    if(ss >> n)
      v.push_back(n);
  } while(ss >> c);
  return v;
}

int main() {
  std::string line;
  std::getline(std::cin, line);
  std::vector<int> v = parseInt(line);
  for (auto const& i: v)
    std::cout << i << std::endl;
  return 0;
}
  