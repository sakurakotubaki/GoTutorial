package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // リファレンスカウンタを1増やす
	go func() {
		fmt.Println("Hello, World!")
		wg.Done() // リファレンスカウンタを1減らす
		// 重たい処理
	}()
	// 別の重たい処理
	wg.Wait() // リファレンスカウンタが0になるまで待つ
}
