package main

// カッコでパッケージのインポートをグループ化している
// import "fmt"
// import "math"
// と書くこともできる
import (
	"fmt"
	"math"
)

func main()  {
	fmt.Printf("Now you have %g problems", math.Sqrt(7))
}
