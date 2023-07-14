# Go言語理解するための学習
おみくじをプログラムで、再現してみた

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数のシードを設定
	rand.Seed(time.Now().UnixNano())

	// おみくじの結果のリスト
	results := []string{
		"大吉",
		"中吉",
		"小吉",
		"吉",
		"凶",
	}

	// 乱数を生成して結果を表示
	randomIndex := rand.Intn(len(results))
	result := results[randomIndex]
	fmt.Println("おみくじの結果:", result)
}
```

**実行結果**
```
hashimotojunichi@hashimotojunichinoMacBook-Pro go-example % go run hello.go
おみくじの結果: 小吉
hashimotojunichi@hashimotojunichinoMacBook-Pro go-example % go run hello.go
おみくじの結果: 小吉
hashimotojunichi@hashimotojunichinoMacBook-Pro go-example % go run hello.go
おみくじの結果: 吉
```

このプログラムでは、rand.Seed(time.Now().UnixNano())によって乱数生成器のシードを設定しています。これにより、実行ごとに異なる結果が得られます。

resultsにはおみくじの結果のリストを格納しています。rand.Intn(len(results))によって、結果のリストの中からランダムにインデックスを選びます。

最後に、選ばれた結果を表示します。

上記のコードをGoのソースファイル（例：omikuji.go）に保存し、コンパイルして実行すると、おみくじの結果が表示されます。

注意点として、おみくじの結果をより多く用意したい場合は、resultsの要素を追加することで対応できます。