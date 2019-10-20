package main

import "fmt"

func main()  {
	fmt.Println(split(17))
}

// 戻り値となる変数に名前をつける( named return value )ことができる
// 戻り値に名前をつけると、関数の最初で定義した変数名として扱われる
// 名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができる
// これを naked return と呼ぶ
// naked return ステートメントは、短い関数でのみ使用するべきである
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
