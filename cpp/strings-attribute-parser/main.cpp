// strings-attribute-parser/main.cpp -*- mode: c++; -*-
// Copyright © 2023 Savio Sena <savio.sena@acm.org>

#include <algorithm>
#include <bits/stdc++.h>
#include <cassert>
#include <cctype>
#include <cmath>
#include <cstdio>
#include <iomanip>
#include <ios>
#include <iostream>
#include <map>
#include <sstream>
#include <string>
#include <vector>

namespace hrml {

// Hrml attributes type.
using attributes = std::map<std::string, std::string>;

struct node;

using nodes = std::map<std::string, hrml::node>;

// Hrml node type.
struct node {
    std::string tag;
    hrml::attributes attrs;
    hrml::nodes children;
};

} // namespace hrml

void skip_spaces(std::string const& line, std::string::size_type& pos,
                 std::string::size_type const& len) {
    while (pos < len && isspace(line[pos]))
        pos++;
}

std::string read_token(std::string const& line, std::string::size_type& pos,
                       std::string::size_type const& len) {
    std::string::size_type beg = pos;
    while (pos < len && !isspace(line[pos]) && line[pos] != '<' &&
           line[pos] != '>') // TODO: if you could conjugate the condition
                             // "line[pos] != '='" here, you'd be able to
                             // eliminate read_attrs_lhs function.
        pos++;
    return line.substr(beg, pos - beg);
}

std::string read_attrs_lhs(std::string const& line, std::string::size_type& pos,
                           std::string::size_type const& len) {
    std::string::size_type beg = pos;
    while (pos < len && !isspace(line[pos]) && line[pos] != '<' &&
           line[pos] != '>' && line[pos] != '=')
        pos++;
    return line.substr(beg, pos - beg);
}

void read_attrs(hrml::attributes& attrs, std::string const& line,
                std::string::size_type& pos,
                std::string::size_type const& len) {
    // TODO implement with find/trim instead of char-by-char.
    do {
        skip_spaces(line, pos, len);
        std::string lhs = read_attrs_lhs(line, pos, len);
        skip_spaces(line, pos, len);
        if (line[pos] == '>')
            break;
        assert(line[pos] == '=');
        pos++;
        skip_spaces(line, pos, len);
        std::string rhs = read_token(line, pos, len);
        if (rhs[0] == '"' || rhs[0] == '\'')
            rhs = rhs.substr(1, rhs.length() - 2);
        attrs[lhs] = rhs;
    } while (line[pos] != '>');
}

void parse(hrml::node& node, std::istream& in, uint32_t& nlines) {
    while (nlines) {
        std::string line;
        std::getline(in, line);
        nlines--;

        std::string::size_type len = line.length();
        std::string::size_type pos = 0;

        skip_spaces(line, pos, len);

        assert(line[pos] == '<');
        pos++; // skip '<'

        skip_spaces(line, pos, len);

        bool is_closing_tag = line[pos] == '/';
        if (is_closing_tag)
            pos++;

        std::string tag = read_token(line, pos, len);

        if (!is_closing_tag) {
            node.children[tag] = std::move(hrml::node{tag, {}, {}});
            read_attrs(node.children[tag].attrs, line, pos, len);
            pos++; // skip '>'
            parse(node.children[tag], in, nlines);
        } else {
            pos++; // skip '>'
            return;
        }
    }
}

std::string query(hrml::node& node, std::string const& cmd) {
    std::string::size_type pos = cmd.find('.');
    if (pos == std::string::npos) {
        std::string::size_type pos = cmd.find('~');
        std::string tag = cmd.substr(0, pos);
        std::string attr = cmd.substr(pos + 1);
        if (node.children.find(tag) != node.children.end() &&
            node.children[tag].attrs.find(attr) !=
                node.children[tag].attrs.end())
            return node.children[tag].attrs[attr];
    } else {
        std::string tag = cmd.substr(0, pos);
        if (node.children.find(tag) != node.children.end())
            return query(node.children[tag], cmd.substr(pos + 1));
    }
    return "Not Found!";
}

void dbg_print(std::string parent, hrml::node const& node) {
    std::string name = parent + (parent != "" ? "." : "") + node.tag;
    std::clog << "-> " << name << ": ";
    for (auto const& [key, val] : node.attrs) {
        std::clog << key << "=" << val << " ";
    }
    std::clog << std::endl;
    for (auto const& [_, val] : node.children) {
        dbg_print(name, val);
    }
}

int main() {
    uint32_t N, Q;
    scanf("%u %u\n", &N, &Q);

    hrml::node document = {"<document>", {}, {}};
    std::string line;

    parse(document, std::cin, N);
    // dbg_print("", document);
    while (Q--) {
        std::getline(std::cin, line);
        std::cout << query(document, line) << std::endl;
    }

    return 0;
}

// vim:ft=cpp:
