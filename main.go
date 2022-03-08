package main

import (
	"fmt"
	"qrcode/qr_windows"
)

func main() {
	matrix := [][]byte{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	qr_windows.WinOutMatrix(matrix)
	fmt.Scanln()
}
