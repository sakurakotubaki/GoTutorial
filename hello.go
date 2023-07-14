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
