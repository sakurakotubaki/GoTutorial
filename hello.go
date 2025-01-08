package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// コンストラクタ関数
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// 名前を設定するメソッド
func (p *Person) SetName(name string) {
	p.Name = name
}

// 年齢を設定するメソッド
func (p *Person) SetAge(age int) {
	p.Age = age
}

// 情報を取得するメソッド
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// メインプログラムでの使用例
func main() {
	// 方法1: コンストラクタを使用
	person1 := NewPerson("田中", 25)
	fmt.Println(person1.GetInfo()) // 出力: 田中 is 25 years old

	// 方法2: 構造体を直接初期化
	person2 := &Person{
		Name: "佐藤",
		Age:  30,
	}

	// メソッドを使って値を変更
	person2.SetName("鈴木")
	person2.SetAge(35)
	fmt.Println(person2.GetInfo()) // 出力: 鈴木 is 35 years old
}
