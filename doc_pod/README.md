# Goで電気ポットを作る
Golangを使用して、電気ポッドのお湯が沸いたら音が鳴るビジネスロジックを考えてみた。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 電気ポットの初期温度
	currentTemperature := 20

	// 目標温度
	targetTemperature := 100

	// カチッと音を鳴らす関数
	playSound := func() {
		fmt.Println("カチッ")
	}

	// 温度が目標温度に達するまでループ
	for currentTemperature < targetTemperature {
		// 温度を上昇させる処理（例：1度ずつ上昇させる）
		currentTemperature++

		// 1度上昇するごとに温度を表示
		fmt.Printf("現在の温度：%d度\n", currentTemperature)

		// お湯が100度になったらカチッと音を鳴らす
		if currentTemperature == targetTemperature {
			playSound()
		}

		// 1秒待機
		time.Sleep(time.Second)
	}
}
```

**実行結果**
```
hashimotojunichi@hashimotojunichinoMacBook-Pro go-example % go run hello.go
現在の温度：21度
現在の温度：22度
現在の温度：23度
現在の温度：24度
現在の温度：25度
現在の温度：26度
# 100度になるまで実行される...
```

このコードでは、currentTemperature変数を使って現在の温度を表現します。targetTemperature変数には目標の温度（ここでは100度）を設定します。

playSoundという関数は、カチッと音を鳴らすための関数です。ここでは、単純に文字列を表示していますが、実際のアプリケーションでは音声ファイルの再生などを行うことが考えられます。

ループの中で、currentTemperatureがtargetTemperatureに達するまで温度を上昇させます。1度上昇するたびに温度を表示し、100度に達したらplaySound関数を呼び出します。

最後に、1秒待機してからループを繰り返します。これにより、温度の上昇を時間的に表現しています。

上記のコードをGoのソースファイル（例：electric_kettle.go）に保存し、コンパイルして実行すると、電気ポットの温度が上昇し、100度になった時にカチッと音が鳴る様子が再現されます。