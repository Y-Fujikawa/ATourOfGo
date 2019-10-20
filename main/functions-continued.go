package main

import "fmt"

func main()  {
	fmt.Println(add2(55, 10))
}

// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できる
func add2(x, y int) int {
	return x + y
}
