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
