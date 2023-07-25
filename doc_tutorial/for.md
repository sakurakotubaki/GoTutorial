# 繰り返し処理をする

```go
package main

import "fmt"

func main() {
	// 配列用の変数
	j := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// jの数だけ配列をループするが、8のところでbreakする
	// _, v := range j で、_はインデックス番号、vは値を取得する
	// range j で、jの配列の数だけループする
	for _, v := range j {
		fmt.Println(v)
		// 8のところでbreakする
		if v == 8 {
			break
		}
	}
}
```