package main

import "fmt"

// Person構造体の定義
type Person struct {
	Name string
	Age  int
}

// 値レシーバーを使用したメソッド
func (p Person) Greet() {
	fmt.Printf("こんにちは、私は%sです。%d歳です。\n", p.Name, p.Age)
}

// ポインタレシーバーを使用したメソッド
func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("お誕生日おめでとう！%sさんは%d歳になりました！\n", p.Name, p.Age)
}

func main() {
	// Personのインスタンスを作成
	person := Person{
		Name: "田中",
		Age:  25,
	}

	// メソッドの呼び出し
	person.Greet()    // 値レシーバー
	person.Birthday() // ポインタレシーバー
	person.Greet()    // 年齢が変更されていることを確認
}