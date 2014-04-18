#!/usr/bin/python
# -*- coding: utf-8 -*-

if __name__ == '__main__':
    (steps, a, b) = (20, 1, 1)

    i = steps * 2
    while i > steps:
        a *= i
        i -= 1
    while steps > 1:
        b *= steps
        steps -= 1

    print a / b
