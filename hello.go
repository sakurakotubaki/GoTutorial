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
