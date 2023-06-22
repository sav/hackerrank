CXX := clang++

CXXFLAGS := -Og -ggdb3 -DDEBUG -Wextra -Wconversion -Wpedantic -Werror -pedantic -std=c++20

LDFLAGS := -lpthread -lcrypto 

all: main

main: main.cpp
	$(CXX) $(CXXFLAGS) $< -o $@ $(LDFLAGS)

.PHONY: clean
clean:
	rm -f main
