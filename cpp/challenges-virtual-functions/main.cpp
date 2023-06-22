// challenges-virtual-functions/main.cpp
// Copyright © 2023 Savio Sena <savio.sena@acm.org>

#include <algorithm>
#include <bits/stdc++.h>
#include <cassert>
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <thread>

namespace hr {

class Person {
protected:
  std::string name;
  int age;

public:
  Person() : name(), age() {}
  Person(Person const &other) : name(other.name), age(other.age) {}
  virtual ~Person() {}

  virtual void getdata() = 0;
  virtual void putdata() = 0;
}; // class Person

class Professor : public Person {
public:
  Professor() : Person() { cur_id = ++Professor::professor_next_id; }
  Professor(Professor const &other)
      : Person(other), publications(other.publications) {
    Professor();
  }

  void getdata();

  void putdata();

protected:
  int cur_id;
  int publications;
  inline static thread_local int professor_next_id;
}; // class Professor

void Professor::getdata() { std::cin >> name >> age >> publications; }

void Professor::putdata() {
  std::cout << name << " " << age << " " << publications << " " << cur_id
            << std::endl;
}

class Student : public Person {
public:
  Student() : Person() {
    std::fill(marks, marks + marks_size, 0);
    cur_id = ++Student::student_next_id;
  }

  Student(Student const &other) : Person(other) {
    new (this) Student();

    for (std::size_t i{}; i < marks_size; ++i) {
      marks[i] = other.marks[i];
    }
  }

  void getdata();

  void putdata();

protected:
  inline static thread_local int student_next_id;
  int cur_id;
  uint64_t marks[6];
  constexpr static std::size_t marks_size = sizeof(marks) / sizeof(marks[0]);

  uint64_t total() {
    uint64_t sum = 0;
#define USE 1
#if USE <= 1
    for (auto& mark : marks) {
		sum += mark;
    }
#elif USE == 2
    for (std::size_t i{}; i < marks_size; ++i) {
      sum += marks[i];
    }
#elif USE >= 3
    std::for_each(marks, &marks[6], [&sum](int mark) { sum += mark; });
#endif
    return sum;
  }
}; // class Student

void Student::getdata() {
  std::cin >> name >> age;

  for (std::size_t i{}; i < marks_size; ++i) {
    std::cin >> marks[i];
  }
}

void Student::putdata() {
  std::cout << name << " " << age << " " << total() << " " << cur_id
            << std::endl;
}

} // namespace hr

int main() {
  int n;
  std::cin >> n;

  while (n--) {
    int kind;

    std::cin >> kind;
    assert(kind >= 1 && kind <= 2);

    if (kind == 1) {
      hr::Professor professor;
      professor.getdata();
      professor.putdata();
    } else {
      hr::Student student;
      student.getdata();
      student.putdata();
    }
  }

  return 0;
}
