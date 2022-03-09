package qr_common

import "testing"

func Test_cut(t *testing.T) {
	matrix := [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
	}
	cutBoundMatrix(&matrix)
	printSource(matrix)
}
