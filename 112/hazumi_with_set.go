package main

import (
    "fmt"
    "strconv"
    set "github.com/deckarep/golang-set"
)

func isHazumi(i int) bool {
    c := strconv.Itoa(i)

    // 重複排除をするためinterfaceで定義
    type any []interface{}
    chk := any{}
    for i := 0; i < len(c); i++ {

        // i + 1 が len(c)を上回る場合はエラー
        if i + 1 >= len(c) {
            continue
        }
        if c[i] < c[i+1] {
            chk = append(chk, 1)
        } else if c[i] > c[i+1] {
            chk = append(chk, 0)
        }
    }

    // 重複を排除
    chks := set.NewSetFromSlice(chk)

    // 弾み数は重複排除した後1以上になる
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

