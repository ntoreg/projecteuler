#!/usr/bin/env python
# -*- coding: utf-8 -*-


def is_hazumi(a):
    if not (a == sorted(a) or a == sorted(a, reverse=True)):
        return True


if __name__ == "__main__":

    x = 0
    i = 0
    while True:
        i += 1
        a = list(str(i))
        if is_hazumi(a):
            x += 1
            if x * 1.0 / i == 0.99:
                print i
                break

"""
$ time python hazumi.py 
1587000

real    0m9.243s
user    0m9.151s
sys     0m0.046s
"""
