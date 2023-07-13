package main

import "fmt"

func main() {
    // int型の変数を宣言
    var num int = 42
    // int型のポインタを宣言し、numのアドレスを代入
    var ptr *int = &num

    // numの値とアドレスを表示
    fmt.Println("num =", num)
    fmt.Println("ptr =", ptr)
    fmt.Println("ptrのアドレス =", &ptr)

    // ポインター経由でnumの値を変更
    *ptr = 100

    // 変更された値を表示
    fmt.Println("num =", num)
}
