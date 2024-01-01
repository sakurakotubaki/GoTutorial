package main

import "fmt"

/*
関数の定義する場所はどこでも良い。実行する場合はmain関数の中で呼び出す必要がある。
counterというfunctionは、int型の引数を受け取り、ループ処理と分岐処理に応じて値を返す。
*/
func counter(v int) int {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("5になりましたので処理を終了します👻")
			break
    } else if i < 5 {
      fmt.Println("5になる前の処理が実行されました。値は:", i)
		} else {
			fmt.Println("処理が実行されませんでした😇")
		}
	}
	return v
}

// greetingというfunctionは、string型の引数を2つ受け取り、それぞれの値を返す。
func greeting(a string, b string) (string, string) {
	return a, b
}

func main() {
	// 引数として、1を渡す。for文の中で、iが5になったらbreakする。
	counter(1)
  // greetingという関数を呼び出し、それぞれの値をx, yに代入する。logを出力する。
	x, y := greeting("私のお名前は", "Jboyです☀️")
	fmt.Println(x, y)
}