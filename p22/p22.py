#!/usr/bin/python
# -*- coding: utf-8 -*-

if __name__ == '__main__':
    f = open('names.txt')
    names = sorted(f.readline().replace('"', '').split(','))
    f.close()

    print sum(map(lambda name: sum([ord(c) - 64 for c in name]) \
              * (names.index(name) + 1), names))
