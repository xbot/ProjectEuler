#!/usr/bin/python
# -*- coding: utf-8 -*-

if __name__ == '__main__':
    product = 1
    for i in range(2, 101):
        product *= i
    print sum([int(x) for x in str(product)])
