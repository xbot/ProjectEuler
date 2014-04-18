#!/usr/bin/python
# -*- coding: utf-8 -*-


def is_leap(year):
    return year % 400 == 0 or year % 100 != 0 and year % 4 == 0


def count_days(year, month, day):
    days = 0
    for y in range(1900, year):
        days += is_leap(y) and 366 or 365
    thirties = [4, 6, 9, 11]
    for m in range(1, month):
        if m == 2:
            days += is_leap(year) and 29 or 28
        else:
            days += m in thirties and 30 or 31
    return days + day


def main():
    count = 0
    for year in range(1901, 2001):
        for month in range(1, 13):
            if count_days(year, month, 1) % 7 == 0:
                count += 1
    print count


if __name__ == '__main__':
    main()
