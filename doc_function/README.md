# Goの関数（Function）について

## 関数の基本構文
```go
func 関数名(引数 型) 戻り値の型 {
    // 処理
    return 戻り値
}
```

## 基本的な関数の例
```go
// 単純な足し算を行う関数
func add(x int, y int) int {
    return x + y
}

// 複数の引数の型が同じ場合は、まとめて書くことができます
func add2(x, y int) int {
    return x + y
}
```

## 複数の戻り値を持つ関数
```go
// 複数の戻り値を返す関数
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return x / y, nil
}

// 使用例
result, err := divide(10, 2)
if err != nil {
    fmt.Println("エラー:", err)
    return
}
fmt.Println("結果:", result)
```

## 名前付き戻り値
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 名前付き戻り値を使用する場合、空のreturnで良い
}
```

## 可変長引数を持つ関数
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 使用例
result := sum(1, 2, 3, 4, 5)  // 15
```

## 関数を値として扱う
```go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

// 使用例
hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(compute(hypot))  // 5
```

## クロージャ（無名関数）
```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

// 使用例
pos := adder()
fmt.Println(pos(1))  // 1
fmt.Println(pos(2))  // 3
fmt.Println(pos(3))  // 6
```

## defer文の使用
```go
func readFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        return
    }
    defer file.Close()  // 関数の終了時に必ず実行される
    
    // ファイルの読み込み処理
}
```

## 関数の使用に関するベストプラクティス
1. 関数は一つの責務のみを持つようにする
2. 適切な名前付けを行う（動詞+名詞の形式が一般的）
3. エラーハンドリングを適切に行う
4. 関数の長さは適度に保つ
5. 引数の数は必要最小限に抑える

## まとめ
- Goの関数は柔軟で強力な機能を提供します
- 複数の戻り値、名前付き戻り値、可変長引数など、様々な機能があります
- defer文を使用することで、リソースの解放を確実に行うことができます
- 関数型プログラミングの要素も取り入れることができます
