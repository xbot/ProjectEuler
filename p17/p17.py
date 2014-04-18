#!/usr/bin/python
# -*- coding: utf-8 -*-


def translate_number(num):
    words = {
        0: '',
        1: 'one',
        2: 'two',
        3: 'three',
        4: 'four',
        5: 'five',
        6: 'six',
        7: 'seven',
        8: 'eight',
        9: 'nine',
        10: 'ten',
        11: 'eleven',
        12: 'twelve',
        13: 'thirteen',
        14: 'fourteen',
        15: 'fifteen',
        16: 'sixteen',
        17: 'seventeen',
        18: 'eighteen',
        19: 'nineteen',
        20: 'twenty',
        30: 'thirty',
        40: 'forty',
        50: 'fifty',
        60: 'sixty',
        70: 'seventy',
        80: 'eighty',
        90: 'ninety',
        }

    english = ''

    if num / 1000 > 0:
        english += translate_number(num / 1000) + ' thousand'
        tmp = translate_number(num % 1000)
        english += ((tmp.strip() == '' or tmp.find('hundred') > -1)
                    and ' ' or ' and ') + tmp
    elif num / 100 > 0:
        english += translate_number(num / 100) + ' hundred'
        tmp = translate_number(num % 100)
        if tmp.strip() != '':
            english += ' and ' + tmp
    else:
        if words.has_key(num):
            english += words[num]
        else:
            english += words[num / 10 * 10] + '-' + words[num % 10]

    return english


if __name__ == '__main__':
    count = 0
    for i in range(1, 1001):
        count += len(translate_number(i).replace(' ', '').replace('-',
                     ''))
    print count
