// Golangはパッケージ（package）で構成される
// プログラムは main パッケージから開始される
package main

// ここでは "fmt" , "math/rand" パッケージをインポート（import）している
import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Println("My favorite number is", rand.Intn(10))
}
