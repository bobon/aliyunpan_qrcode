package main

import (
	"fmt"
	"qrcode/qr_common"
)

func main() {
	src := `test_src.png`
	fmt.Println(qr_common.PrintFromFile(src))
}
