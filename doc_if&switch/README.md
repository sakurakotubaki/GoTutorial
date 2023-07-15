# Goで条件分岐を使う
if文を使った学習をするために、いろんなパターンのコードを書いてみる

## 🧷ifを使った例
¥500あったらワンコインランチが、帰る条件分岐
```go
package main

import "fmt"

func main() {
	price := 450

	if price == 500 {
		fmt.Println("ワンコインランチが食べられる！")
	} else if price < 500 {
		fmt.Println("ワンコインランチが買えない...")
	}
}
```

**実行結果**
変数priceの値は、450なので、500未満の時の処理が実行される
```
ワンコインランチが買えない...
```

## 🧷switch文を使う
Go言語でもswitch文を使うことができる。
ラーメンの券売機をイメージして作ってみた。
```go
package main

import "fmt"

func main() {
	var price int = 700

	switch price {
		case 700:
			fmt.Println("ラーメン: ¥700")
		case 800:
			fmt.Println("チャーシューメン: ¥800")
		case 750:
			fmt.Println("のりラーメン: ¥750")
		case 300:
			fmt.Println("チャーハン: ¥300")
		default:
			fmt.Println("メニューがありません!")
	}
}
```

**実行結果**
変数priceが、700の時の実行結果
```
ラーメン: ¥700
```