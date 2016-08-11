package main

import (
    "fmt"
    "strconv"
    set "github.com/deckarep/golang-set"
)

func isHazumi(i int) bool {
    c := strconv.Itoa(i)

    // $B=EJ#GS=|$r$9$k$?$a(Binterface$B$GDj5A(B
    type any []interface{}
    chk := any{}
    for i := 0; i < len(c); i++ {

        // i + 1 $B$,(B len(c)$B$r>e2s$k>l9g$O%(%i!<(B
        if i + 1 >= len(c) {
            continue
        }
        if c[i] < c[i+1] {
            chk = append(chk, 1)
        } else if c[i] > c[i+1] {
            chk = append(chk, 0)
        }
    }

    // $B=EJ#$rGS=|(B
    chks := set.NewSetFromSlice(chk)

    // $BCF$_?t$O=EJ#GS=|$7$?8e(B1$B0J>e$K$J$k(B
    if chks.Cardinality() > 1 {
        return true
    }
    return false
}

func main() {
    i := 0
    x := 0
    for {
        i++
        if isHazumi(i) {
            x++
            per := float64(x) / float64(i)
            if per == 0.99 {
                fmt.Println(i)
                break
            }
        }
    }
}

