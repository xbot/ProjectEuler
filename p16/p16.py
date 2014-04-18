#!/usr/bin/python
# -*- coding: utf-8 -*-

import math

if __name__ == '__main__':
    str = format(math.pow(2, 1000), 'f')
    sum = 0
    for c in str[:str.index('.')]:
        sum += int(c)
    print sum
