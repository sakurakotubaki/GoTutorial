# goroutine(ゴルーチン)
Goのランタイムで管理された軽量なスレッド

```go
message := "hi"

go sendMessage(message)
```

`goroutine`を実行するには、関数呼び出しの前に`go`をつけるだけ。その呼び出し時には、引数がキャプチャされる。以下のように`message`を変換しても呼び出しされる引数は変わらない。

```go
message := "hi"// (1)

go sendMessage(message)// (2)
message = "ho"// (3)
```

(1) (2) (3) の順番で実行されることを期待するが、Goのランタイムスケジューラが(3)よりも(2)を先に実行するかはわからない。以下の例では`ho`が２回出力されることがあります。

```bash
hashimotojunichi@hashimochinoMBP GoTutorial % go run hello.go 
ho
ho
```

これは、race conditionと呼ばれ、goroutineとその呼び出しも元との間でデータ競合が起きていることを意味している。

race conditionは実行していても不整合が起きるまでは気づくことはできないが、Goには、race condition detectorが付属されていて、コンパイル時に^raceをつけて実行することで、ランタイムがrace conditionを検出してくれる。

```bash
go build -race hello.go
```

`goroutine`は関数の中で実行しても動き続ける。逆言うと、main関数の中で実行すると、`goroutine`が実行中にもかかわらずmain関数の終了とともに強制終了してしまう。main関数が、`goroutine`の終了を待つためには、`syncパッケージ`を使う必要がある。

```go
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
```