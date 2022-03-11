package qr_common

import (
	"fmt"
	"image"
	"os"
	"qrcode/qr_windows"
)

const (
	win_black = 0
	win_white = 15
)

// -----------------------
// func Base64ToFile(src string) error {
// 	dst, err := base64.StdEncoding.DecodeString(src)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//	ioutil.WriteFile("test2_src.png", dst, fs.FileMode(777))

// 	img, str, err := image.Decode(strings.NewReader(string(dst)))
// 	fmt.Println(str, err)
// 	if err != nil {
// 		return err
// 	}
// 	result := getMatrixByAnalysisBoundAndLand(getMatrixFromImage(img))
// 	bound := getMatrixBound(result, 1)
// 	qr_windows.WinOutMatrix(bound)
// 	return nil
// }

// ---------------------
func PrintFromFile(path string) error {
	matrix, err := getSourceMatrixFromFile(path)
	if err != nil {
		return err
	}
	result := getMatrixByAnalysisBoundAndLand(matrix)
	bound := getMatrixBound(result, 1)
	qr_windows.WinOutMatrix(bound)
	//fmt.Println("----------------------------------------------")
	//WinOutMatrix(bound)
	return nil
}

//打印
func printSource(matrix [][]byte) {
	for k1 := range matrix {
		for k2 := range matrix[k1] {
			print(matrix[k1][k2])
		}
		println()
	}
}

//从文件中获取源矩阵
func getSourceMatrixFromFile(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return getMatrixFromImage(img), nil
}

//从image对象中获取二值化矩阵
//grey1 ==> [ 1 : >125    0 : <=125 ]
func getMatrixFromImage(img image.Image) [][]byte {
	rect := img.Bounds()
	matrix := [][]byte{}
	for y := rect.Min.Y; y < rect.Max.Y; y++ { //坐标系 左上(0,0)  右下 (max,max)
		bts := []byte{}
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if grey1(img.At(x, y).RGBA()) > 125 {
				bts = append(bts, common_white)
			} else {
				bts = append(bts, common_black)
			}
		}
		matrix = append(matrix, bts)
	}
	return matrix
}

/*
分析裁剪原始矩阵，获取二维码的边界 和 单个块的厚度
得到最终矩阵
*/
func getMatrixByAnalysisBoundAndLand(matrix [][]byte) [][]byte {
	result := [][]byte{}
	cutBoundMatrix(&matrix)
	land := getSingleBlockLand(matrix)
	for k := range matrix {
		rtl := []byte{}
		for kk := range matrix {
			if (kk+1)%land == 0 {
				rtl = append(rtl, matrix[k][kk])
			}
		}
		if (k+1)%land == 0 {
			result = append(result, rtl)
		}
	}
	//printSource(result)
	return result
}

//给矩阵增加边框，提高对比度
func getMatrixBound(matrix [][]byte, frame int) [][]byte {
	if frame > 0 {
		lastMatrix := make([][]byte, len(matrix)+frame*2)
		for k := range lastMatrix {
			lastMatrix[k] = make([]byte, len(matrix)+frame*2)
			for kk := range lastMatrix[k] {
				lastMatrix[k][kk] = common_white
				if (k >= frame && k < len(lastMatrix)-frame) &&
					(kk >= frame && kk < len(lastMatrix)-frame) {
					lastMatrix[k][kk] = matrix[k-frame][kk-frame]
				}
			}
		}
		return lastMatrix
	}
	return matrix
}

//裁剪矩阵边界
func cutBoundMatrix(matrix *[][]byte) {
	mtx := *matrix
	if len(mtx) == 0 || len(mtx[0]) == 0 {
		return
	}
	ws, we := 0, len(mtx)-1    //宽起始点和结束点
	ls, le := 0, len(mtx[0])-1 //长起始点和结束点
	//ws
	flag := false
	for k1 := 0; k1 < len(*matrix); k1++ {
		flag = false
		for k2 := range mtx[k1] {
			if mtx[k1][k2] == common_black {
				ws = k1
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	//we
	for k1 := len(*matrix) - 1; k1 >= 0; k1-- {
		flag = false
		for k2 := range mtx[k1] {
			if mtx[k1][k2] == common_black {
				we = k1
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	//ls
	for k1 := 0; k1 < len(mtx[0]); k1++ {
		flag = false
		for k2 := 0; k2 < len(*matrix); k2++ {
			if mtx[k2][k1] == common_black {
				ls = k1
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	//le
	for k1 := len(mtx[0]) - 1; k1 >= 0; k1-- {
		flag = false
		for k2 := 0; k2 < len(*matrix); k2++ {
			if mtx[k2][k1] == common_black {
				le = k1
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	fmt.Println("ws, we", ws, we)
	fmt.Println("ls, le", ls, le)

	//剪裁矩阵
	*matrix = mtx[ws : we+1]
	mtx = *matrix
	for k := range *matrix {
		mtx[k] = mtx[k][ls : le+1]
	}
}

//获取矩阵单格块的厚度
func getSingleBlockLand(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	//加上掩码图案后，仍然是3个方块，只需要取上下边最长的一条白线，即可拿到 定位坐标的长度
	upFlag, downFlag := true, true
	upLength, downLength := 0, 0
	for k := range matrix[0] {
		if !upFlag && !downFlag {
			break
		}
		if upFlag {
			if matrix[0][k] == common_black {
				upLength++
			} else {
				upFlag = false
			}
		}
		if downFlag {
			if matrix[len(matrix)-1][k] == common_black {
				downLength++
			} else {
				downFlag = false
			}
		}
	}
	//层数  7 -> 5 -> 3
	if upLength > downLength {
		return upLength / 7
	}
	return downLength / 7
}

// 二值化方案1  白字   三等分
// 灰度化后的R=（处理前的R + 处理前的G +处理前的B）/ 3
// 灰度化后的G=（处理前的R + 处理前的G +处理前的B）/ 3
// 灰度化后的B=（处理前的R + 处理前的G +处理前的B）/ 3
func grey1(r, g, b, a uint32) uint8 {
	r = r >> 8
	g = g >> 8
	b = b >> 8
	return uint8((r + b + g) / 3)
}

// 二值化方案2  黑字   特殊图片
// 灰度化后的R =  处理前的R * 0.3+ 处理前的G * 0.59 +处理前的B * 0.11
// 灰度化后的G =  处理前的R * 0.3+ 处理前的G * 0.59 +处理前的B * 0.11
// 灰度化后的B =  处理前的R * 0.3+ 处理前的G * 0.59 +处理前的B * 0.11
func grey2(r, g, b, a uint32) uint8 {
	r = (r >> 8) * 3000
	g = (g >> 8) * 5900
	b = (b >> 8) * 1100
	return uint8((r + b + g) / 300)
}
