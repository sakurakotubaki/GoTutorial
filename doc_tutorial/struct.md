# 構造体とポインター
```go
package main

import "fmt"

// 人間という構造体を定義
type Human struct {
		name string
		age int
}

// 構造体を表示したポインターがある関数を定義する
func (h *Human) HPoint() string {
		return fmt.Sprintf("My name is %v", h.name)
}

func main() {
	// 構造体の初期化
	i := Human{"Taro", 20}
	j := 00000000
	fmt.Println(i.name)
	fmt.Println(i.age)
	fmt.Println("変数jのメモリの場所", &j)
}
```