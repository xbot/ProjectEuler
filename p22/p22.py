#!/usr/bin/python
# -*- coding: utf-8 -*-

if __name__ == '__main__':
    f = open('names.txt')
    names = sorted(f.readline().replace('"', '').split(','))
    f.close()

    # Solution 1: using map() and lambda expression

    print sum(map(lambda name: sum([ord(c) - 64 for c in name]) \
              * (names.index(name) + 1), names))

    # Solution 2: using enumerate()

    print sum(sum([ord(c) - 64 for c in name]) * (i + 1) for (i,
              name) in enumerate(names))
