package qrcode2console

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"

	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
)

//通用 跨平台
const (
	common_black                    byte = 1  //二值化存储 黑景色值
	common_white                    byte = 0  //二值化存储 白景色值
	common_console_black_background int  = 40 // 控制台 黑色 背景
	common_console_white_background int  = 47 // 控制台 白色 背景
	common_console_black_foreground int  = 30 // 控制台 黑色 前景
	common_console_white_foreground int  = 37 // 控制台 白色 前景

	default_character = "▇▇"
)

const (
	file_type_png  = "image/png"  //png文件格式
	file_type_gif  = "image/gif"  //png文件格式
	file_type_jpg  = "image/jpeg" //png文件格式
	file_type_webp = "image/webp" //png文件格式
	file_type_bmp  = "image/bmp"  //png文件格式
)

type qrcode struct {
	img                                                   image.Image //img对象  用于base64对象可以方便输出文件
	sourceMinBound, genBound                              int         //原始最小边界值，生成后的边界
	isBoundSet                                            bool
	character                                             string   //输出字符 一般2个 ▇▇ 可根据控制台大小调整 个数
	imgBinaryMatrix, cutMatrix, shrinkMatrix, boundMatrix [][]byte //二值化图像矩阵,剪裁后的矩阵，缩放后的矩阵，加边界后的矩阵
	shrinkRate                                            int      //缩放比例
}

//从base64中获取qr对象
func NewQrcodeFromBase64(src string) (*qrcode, error) {
	bts, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return newQRFromBytes(bts)
}

//从文件中获取qr对象
func NewQrcodeFromPath(path string) (*qrcode, error) {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return newQRFromBytes(bts)
}

//创建
func newQRFromBytes(bts []byte) (*qrcode, error) {
	qr := &qrcode{character: default_character}
	if err := qr.genImg(bts); err != nil {
		return qr, err
	}
	qr.binaryImg()
	if len(qr.imgBinaryMatrix) == 0 ||
		len(qr.imgBinaryMatrix[0]) == 0 {
		return qr, errors.New("图片非二维码")
	}
	qr.cut()
	qr.shrink()
	return qr, nil
}

// //向windows控制台输出图像
// //边框
// func (qr *qrcode) PrintOnWindowsConsole() {
// 	qr.bound()
// 	defer qr_windows.WinColorPrint("", qr_windows.Win_light_gray) //输出后，如果最后一个色块是黑色，将导致控制台为黑，此处恢复默认值
// 	for k1 := range qr.boundMatrix {
// 		for k2 := range qr.boundMatrix[k1] {
// 			if qr.boundMatrix[k1][k2] == common_white {
// 				WinColorPrint(qr.character, qr_windows.Win_white)
// 			} else {
// 				WinColorPrint(qr.character, qr_windows.Win_black)
// 			}
// 		}
// 		println()
// 	}
// }

//向全平台控制台输出图像
func (qr *qrcode) PrintForConsole() {
	qr.bound()
	for k1 := range qr.boundMatrix {
		str := ""
		for k2 := range qr.boundMatrix[k1] {
			if qr.boundMatrix[k1][k2] == common_white {
				str += CommonColorPrint(qr.character, Common_white)
			} else {
				str += CommonColorPrint(qr.character, Common_black)
			}
		}
		println(str)
	}
}

// //打印控制台输出的源矩阵
// func (qr *qrcode) PrintSourceBinaryMatrix() {
// 	for k1 := range qr.imgBinaryMatrix {
// 		for k2 := range qr.imgBinaryMatrix[k1] {
// 			print(qr.imgBinaryMatrix[k1][k2])
// 		}
// 		println()
// 	}
// }

// //打印控制台输出的源矩阵
// func (qr *qrcode) PrintSourceCutMatrix() {
// 	for k1 := range qr.cutMatrix {
// 		for k2 := range qr.cutMatrix[k1] {
// 			print(qr.cutMatrix[k1][k2])
// 		}
// 		println()
// 	}
// }

// //打印控制台输出的源矩阵
// func (qr *qrcode) PrintSourceShrinkMatrix() {
// 	for k1 := range qr.shrinkMatrix {
// 		for k2 := range qr.shrinkMatrix[k1] {
// 			print(qr.shrinkMatrix[k1][k2])
// 		}
// 		println()
// 	}
// }

//设置外边框
func (qr *qrcode) SetBound(bound int) {
	qr.genBound = bound
	qr.isBoundSet = true
}

//设置控制台输出字符，否则使用默认 ▇▇
func (qr *qrcode) SetOutputCharacter(character string) {
	qr.character = character
}

//外边框修饰  优先使用指定边框，其次使用默认缩放时计算的原比例边框
func (qr *qrcode) bound() {
	frame := 0
	if qr.genBound >= 0 && qr.isBoundSet {
		frame = qr.genBound
	} else if qr.sourceMinBound > 0 {
		frame = qr.sourceMinBound / qr.shrinkRate
	}
	qr.boundMatrix = qr.shrinkMatrix
	if frame > 0 {
		qr.boundMatrix = make([][]byte, len(qr.shrinkMatrix)+frame*2)
		for k := range qr.boundMatrix {
			qr.boundMatrix[k] = make([]byte, len(qr.shrinkMatrix)+frame*2)
			for kk := range qr.boundMatrix[k] {
				qr.boundMatrix[k][kk] = common_white
				if (k >= frame && k < len(qr.boundMatrix)-frame) &&
					(kk >= frame && kk < len(qr.boundMatrix)-frame) {
					qr.boundMatrix[k][kk] = qr.shrinkMatrix[k-frame][kk-frame]
				}
			}
		}
	}
}

//缩放图像  必须保证是正常图像，图像外无乱码，否则缩放异常
func (qr *qrcode) shrink() {
	//加上掩码图案后，仍然是3个方块，只需要取上下边最长的一条白线，即可拿到 定位坐标的长度
	upFlag, downFlag := true, true
	upLength, downLength := 0, 0
	for k := range qr.cutMatrix[0] {
		if !upFlag && !downFlag {
			break
		}
		if upFlag {
			if qr.cutMatrix[0][k] == common_black {
				upLength++
			} else {
				upFlag = false
			}
		}
		if downFlag {
			if qr.cutMatrix[len(qr.cutMatrix)-1][k] == common_black {
				downLength++
			} else {
				downFlag = false
			}
		}
	}
	//层数  7 -> 5 -> 3
	qr.shrinkRate = downLength / 7
	if upLength > downLength {
		qr.shrinkRate = upLength / 7
	}
	//缩放
	qr.shrinkMatrix = make([][]byte, len(qr.cutMatrix)/qr.shrinkRate)
	for k, si, si2 := 0, 0, 0; k < len(qr.cutMatrix); k++ {
		if (k-qr.shrinkRate/2)%qr.shrinkRate == 0 { //防止边缘像素识别异常，取像素中值为色块
			qr.shrinkMatrix[si] = make([]byte, len(qr.cutMatrix[k])/qr.shrinkRate)
			si2 = 0
			for kk := 0; kk < len(qr.cutMatrix[k]); kk++ {
				if (kk-qr.shrinkRate/2)%qr.shrinkRate == 0 {
					qr.shrinkMatrix[si][si2] = qr.cutMatrix[k][kk]
					si2++
				}
			}
			si++
		}
	}
}

//剪裁边界 二维码图案外如果不是纯白则剪裁异常
func (qr *qrcode) cut() {
	//寻找有效矩阵的边界
	ws, we := 0, len(qr.imgBinaryMatrix)-1    //宽起始点和结束点
	ls, le := 0, len(qr.imgBinaryMatrix[0])-1 //长起始点和结束点
	//ws
	flag := false
	for k1 := 0; k1 < len(qr.imgBinaryMatrix); k1++ {
		flag = false
		for k2 := range qr.imgBinaryMatrix[k1] {
			if qr.imgBinaryMatrix[k1][k2] == common_black {
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
	for k1 := len(qr.imgBinaryMatrix) - 1; k1 >= 0; k1-- {
		flag = false
		for k2 := range qr.imgBinaryMatrix[k1] {
			if qr.imgBinaryMatrix[k1][k2] == common_black {
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
	for k1 := 0; k1 < len(qr.imgBinaryMatrix[0]); k1++ {
		flag = false
		for k2 := 0; k2 < len(qr.imgBinaryMatrix); k2++ {
			if qr.imgBinaryMatrix[k2][k1] == common_black {
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
	for k1 := len(qr.imgBinaryMatrix[0]) - 1; k1 >= 0; k1-- {
		flag = false
		for k2 := 0; k2 < len(qr.imgBinaryMatrix); k2++ {
			if qr.imgBinaryMatrix[k2][k1] == common_black {
				le = k1
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	//生成剪裁矩阵
	qr.cutMatrix = make([][]byte, len(qr.imgBinaryMatrix))
	copy(qr.cutMatrix, qr.imgBinaryMatrix)
	qr.cutMatrix = qr.cutMatrix[ws : we+1]
	for k := range qr.cutMatrix {
		qr.cutMatrix[k] = qr.cutMatrix[k][ls : le+1]
	}
	//存储原外边框的最小值,
	qr.sourceMinBound = ws
	if ls < ws {
		qr.sourceMinBound = ls
	}
	if len(qr.imgBinaryMatrix)-we-1 < qr.sourceMinBound {
		qr.sourceMinBound = len(qr.imgBinaryMatrix) - we - 1
	}
	if len(qr.imgBinaryMatrix[0])-le-1 < qr.sourceMinBound {
		qr.sourceMinBound = len(qr.imgBinaryMatrix[0]) - le - 1
	}
}

//二值化图像
func (qr *qrcode) binaryImg() {
	rect := qr.img.Bounds()
	qr.imgBinaryMatrix = make([][]byte, rect.Dy())
	for y := rect.Min.Y; y < rect.Max.Y; y++ { //坐标系 左上(0,0)  右下 (max,max)
		qr.imgBinaryMatrix[y] = make([]byte, rect.Dx())
		for x := rect.Min.X; x < rect.Max.X; x++ {
			qr.imgBinaryMatrix[y][x] = common_black
			if grey1(qr.img.At(x, y).RGBA()) > 125 {
				qr.imgBinaryMatrix[y][x] = common_white
			}
		}
	}
}

//生成图像
func (qr *qrcode) genImg(bts []byte) (err error) {
	switch http.DetectContentType(bts) {
	case file_type_png:
		qr.img, err = png.Decode(bytes.NewReader(bts))
	case file_type_jpg:
		qr.img, err = jpeg.Decode(bytes.NewReader(bts))
	case file_type_gif:
		qr.img, err = gif.Decode(bytes.NewReader(bts))
	case file_type_bmp:
		qr.img, err = bmp.Decode(bytes.NewReader(bts))
	case file_type_webp:
		qr.img, err = webp.Decode(bytes.NewReader(bts))
	default:
		return errors.New("暂不支持的文件类型")
	}
	return
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
