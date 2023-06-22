# Makefile: Build script for my hackerchallenge solutions.
# Copyright © 2023  Savio S. Machado <sav@loophole.cc>

DIRS = $(wildcard      \
    challenge*         \
    c-tutorial*        \
	introduction*      \
	other*             \
	stl*               \
    strings*           \
)

PROGS = $(DIRS:=/main)

TESTS = $(DIRS:=-test)

CC = clang -std=c1x

CXX = clang++

CFLAGS = -Wall

CXXFLAGS = -std=c++20 $(CFLAGS)

LDFLAGS = -lpthread -lcrypto

ifdef DEBUG
   CFLAGS += -ggdb3 -Og
   CXXFLAGS += -ggdb3 -Og
   LDFLAGS += -ggdb3
else
   CFLAGS += -O3 -DNDEBUG -Wextra -Wconversion -Wpedantic -Werror -pedantic
   CXXFLAGS += -O3 -DNDEBUG -Wextra -Wpedantic -Werror
endif

ifdef VERBOSE
   CC += -v
   CXX += -v
endif

.PHONY: clean %-test test

all: $(DIRS)

$(DIRS): $(PROGS)

%/main: %/main.cpp
	$(CXX) $(CFLAGS) $(CXXFLAGS) -o $@ $< $(LDFLAGS)

%-test: %/main
	@./test.sh $(shell dirname ./$<)

test: $(TESTS)

clean:
	@rm -f $(PROGS)
