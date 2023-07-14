package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var choices = []string{"✊", "✌️", "✋"}

func main() {
	// 乱数のシードを設定
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("じゃんけんゲームを始めます！")
	fmt.Println("✊: グー, ✌️: チョキ, ✋: パー")
	fmt.Print("選択してください（✊, ✌️, ✋）: ")

	// プレイヤーの選択
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.TrimSpace(playerChoice)

	// コンピュータの選択
	computerChoice := choices[rand.Intn(len(choices))]

	fmt.Println("あなた:", playerChoice)
	fmt.Println("コンピュータ:", computerChoice)

	// 結果の判定
	result := judge(playerChoice, computerChoice)
	fmt.Println("結果:", result)
}

func judge(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "引き分け"
	} else if (playerChoice == "✊" && computerChoice == "✌️") ||
		(playerChoice == "✌️" && computerChoice == "✋") ||
		(playerChoice == "✋" && computerChoice == "✊") {
		return "あなたの勝ち"
	} else {
		return "コンピュータの勝ち"
	}
}
