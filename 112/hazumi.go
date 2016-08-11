package main

import (
    "fmt"
    "strconv"
)

func isUp(n int) bool {
    c := strconv.Itoa(n)
    for i := 0; i < len(c); i++ {
        if i + 1 >= len(c) {
            continue
        }
        if c[i] > c[i+1] {
            return false
        }
    }
    return true
}

func isDown(n int) bool {
  c := strconv.Itoa(n)
    for i := 0; i < len(c); i++ {
        if i + 1 >= len(c) {
            continue
        }
        if c[i] < c[i+1] {
            return false
        }
    }
    return true
}

func main() {
    n := 0
    x := 0
    for {
        n++
        if !isUp(n) && !isDown(n) {
            x++
            per := float64(x) / float64(n)
            if per == 0.99 {
                fmt.Println(n)
                break
            }
        }
    }
}

