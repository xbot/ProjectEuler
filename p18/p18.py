#!/usr/bin/python
# -*- coding: utf-8 -*-


def main():
    matrix = []
    file = open('data_p18.txt')
    for line in file.readlines():
        matrix.append(map(int, line.replace('\n', '').split(' ')))

    (x, y) = (0, 0)
    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            if y > 0:
                greaterParentPathValue = 0
                if x > 0:
                    greaterParentPathValue = matrix[y - 1][x - 1]
                if x < len(matrix[y - 1]) and matrix[y - 1][x] \
                    > greaterParentPathValue:
                    greaterParentPathValue = matrix[y - 1][x]
                matrix[y][x] += greaterParentPathValue

    print max(matrix[-1])


if __name__ == '__main__':
    main()
