# Golangでタイマーを作ってみる
Golangをキャッチアップするのに、いい方法はないかと考えていたときに思いついた方法は、Swiftの学習をしていたときに、やっていたビジネスロジックの学習でした。

## ⏲️を作ってみる
for文の学習をする課題として、タイマーを作ってみることにしました。10秒経つと、for文のbreak文で、ループから抜ける処理を実行します。

**サンプルコード**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) // 1秒ごとにカウントするタイマーを作成

	for count := 1; ; count++ {
		<-ticker.C // タイマーのチャネルからの受信を待機

		fmt.Println("Count:", count)

		// ここにループを抜ける条件を記述する場合、break文を使用します
		if count >= 10 {
		    break
		}
	}
}
```
go run  hello.goと入力して、コンパイルを実行
**実行結果**
```
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
Count: 6
Count: 7
Count: 8
Count: 9
Count: 10
```