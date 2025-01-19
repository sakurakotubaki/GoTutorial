# Goで配列を操作するには?
Goには配列とスライスがある。似ているが異なる。配列は固定長。スライスは可変長な配列と
考える。０からはじまる。配列の長さを超えた位置を参照すると`panic`が発生する。

```go
var a [4]int
a[0] = 1
```

intの長さ４の配列を宣言。長さを変更することはできない。

## 配列を定義する
配列の使用例。

```go
package main

import "fmt"

func main() {
	var colors [3]string = [3]string{"red", "green", "blue"}
	
	// 配列の要素数を省略することもできる
	// var colors [3]string = [...]string{"red", "green", "blue"}
  
	// logに出す
	fmt.Println(colors)
}
```

## スライスを定義する
スライスは長さを指定せずに宣言する。
```go
var a[]int
```
宣言するだけでは`a`には`nil`という特別な値が格納される。そのまま要素にアクセスすると
`panic`が発生する。スライスを作成するには`make`を使用する。
```go
a := make([]int, 3)
a[0] = 1
a[1] = 2
a[2] = 3
```

また、スライスの初期化は初期値を代入することでも可能。
```go
a := []int{1, 2, 3}
```

使用例。
```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}
```

このようにも使える:
```go
package main

import "fmt"

func main() {
	// スライスの作成
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("元のスライス:", numbers)

	// スライスの要素を変更
	numbers[0] = 10
	fmt.Println("要素を変更したスライス:", numbers)

	// スライスの一部を取得
	slice1 := numbers[1:3]
	fmt.Println("一部を取得したスライス:", slice1)

	// スライスの要素を追加
	slice1 = append(slice1, 6)
	fmt.Println("要素を追加したスライス:", slice1)

	// スライスの要素を削除
	slice1 = append(slice1[:1], slice1[2:]...)
	fmt.Println("要素を削除したスライス:", slice1)

	// スライスの長さと容量
	fmt.Println("スライスの長さ:", len(slice1))
	fmt.Println("スライスの容量:", cap(slice1))
}
```

-------

Go言語では、スライスを使って配列を操作することが多いそうです。

## 📦要素の追加・更新・削除をやってみる
```go
package main

import "fmt"

func main() {
	// 空のスライスを定義
	var fruits []string

	// 要素の追加
	// fruits = append と書くのは、スライスの末尾に要素を追加するという意味
	// ()の中に、fruits, と書くのは、スライスの中身を展開している
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "grape")
	fruits = append(fruits, "orange")

	// 要素の取得
	// fruits[0] は、スライスの先頭の要素を取得するという意味
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])
	fmt.Println(fruits[3])

	// 全ての要素を取得
	// fruits[:] は、スライスの全ての要素を取得するという意味
	fmt.Println(fruits[:])
	// 配列の数を取得するには、len()を使う
	fmt.Println(len(fruits))

	// 要素の上書き
	fruits[0] = "lemon"
	fmt.Println(fruits[0])

	// 要素の削除
	// 0番目を削除するには、最初に0番目の要素を取得して、それ以降の要素を取得して、それをスライスに入れる
	/// fruits[1:] は、スライスの2番目の要素から最後までを取得するという意味
	fruits = append(fruits[:0], fruits[1:]...)

	fmt.Println(fruits[:])
	fmt.Println(len(fruits))
}
```

**実行結果**
```
apple
banana
grape
orange
[apple banana grape orange]
4
lemon
[banana grape orange]
3
```