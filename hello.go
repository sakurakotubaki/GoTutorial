package main

import "fmt"

// 配列の中の挨拶をfor文で表示する関数
func hello() {
    // 配列の中の挨拶をfor文で表示する
    // greetings := []stringは、string型の配列を作るという意味
    greetings := []string{"Good morning", "Good afternoon", "Good night"}
    // forの後の_,は、index番号を使わないという意味。
    // greeting := range greetingsは、配列の中身をgreetingに入れるという意味
    for _, greeting := range greetings {
        // 配列の中身を表示する
        fmt.Println(greeting)
    }
}

func main() {
    // hello関数を呼び出す
    hello()
}