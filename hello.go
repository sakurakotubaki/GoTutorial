package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 標準入力の準備
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("動物の鳴き声クイズ！")
	fmt.Println("犬、猫、または猿の鳴き声を当ててください。")

	// 問題のデータを作成
	questions := map[string]string{
		"1. ワンワン": "犬",
		"2. ニャーニャー": "猫",
		"3. キーキー": "猿",
	}

	correctCount := 0 // 正解の数をカウント
  // 問題を出題
	for question, answer := range questions {
		// 問題を表示
		fmt.Println(question)
    // ユーザーの回答を受け取る。, _ はエラーを無視するための記号
		userAnswer, _ := reader.ReadString('\n')
		// 回答を整形
		userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))
    // 回答をチェック
		if userAnswer == answer {
			fmt.Println("正解！")
			correctCount++
		} else {
			fmt.Println("不正解！正解は", answer, "です。")
		}
	}

	fmt.Println("クイズ終了！正解数:", correctCount)
}
