// cmd/以下にバイナリビルド用のmainパッケージが配置される
package main

import (
	"fmt"
	"greeting"
)

func main() {
	fmt.Println(greeting.Do())
}
