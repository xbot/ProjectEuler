#!/usr/bin/python
# -*- coding: utf-8 -*-

import math


def get_perm(digits, number):
    if len(digits) == 0:
        return ''
    (perm, counter) = ('', math.factorial(len(digits) - 1))
    for d in digits:
        if counter >= number:
            digits.remove(d)
            perm += str(d) + get_perm(digits, number)
            break
        number -= counter
    return perm


if __name__ == '__main__':
    import time
    startTime = time.time()
    digits = [d for d in range(10)]
    print get_perm(digits, 1000000), '%sms' % ((time.time()
            - startTime) * 1000)
