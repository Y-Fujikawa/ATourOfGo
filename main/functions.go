package main

import "fmt"

func main()  {
	fmt.Println(add(42, 13))
}

// 関数は0個以上の引数を受け取ることができる
// 変数名の後ろに型名を書くことに注意すること
func add(x int, y int) int {
	return x + y
}
