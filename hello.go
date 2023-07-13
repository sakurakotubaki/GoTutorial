package main

import "fmt"

func main() {
    // numsは、int型のスライス
    nums := []int{2, 3, 4}
    // sumは、int型
    sum := 0
    // _,は、インデックスを無視することを示す
    // rangeは、スライスやマップを反復処理する
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    // rangeは、インデックスと値の両方を返す
    for i, num := range nums {
        // num == 3のとき、iを出力する
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    // kvsは、map[string]string型
    kvs := map[string]string{"a": "apple", "b": "banana"}
    // for k, v := range kvsは、mapのキーと値を反復処理する
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    // for k := range kvsは、mapのキーを反復処理する
    for k := range kvs {
        fmt.Println("key:", k)
    }
    // for i, c := range "go"は、文字列の文字とそのインデックスを反復処理する
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

// まとめ
// rangeは、スライスやマップを反復処理する