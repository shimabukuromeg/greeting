// cmd/以下にバイナリビルド用のmainパッケージが配置される
package main

import (
	"fmt"

	"github.com/shimabukuromeg/greeting"
)

func main() {
	fmt.Println(greeting.Do())
}
