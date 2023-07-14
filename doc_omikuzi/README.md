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