#!/usr/bin/python
# -*- coding: utf-8 -*-


def sum_proper_factors(n):
    (result, sqrt) = (1, n ** 0.5)

    (start, step) = n % 2 == 1 and (3, 2) or (2, 1)
    for i in range(start, int(sqrt) + 1, step):
        if n % i == 0:
            result += i + n / i

    if sqrt == int(sqrt):
        result -= sqrt

    return result


def main():
    result = 0
    for i in range(1, 10000):
        sum1 = sum_proper_factors(i)
        if sum1 > i:
            if i == sum_proper_factors(sum1):
                result += i + sum1
    print result


if __name__ == '__main__':
    import time
    startTime = time.time()
    main()
    print time.time() - startTime
