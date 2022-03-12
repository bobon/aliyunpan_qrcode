package qrcode2console

import "fmt"

//跨平台通用
const (
	Common_black = 30 //黑
	Common_white = 37 //二值化存储 白景色值

	Common_console_black_background int = 40 // 控制台 黑色 背景
	Common_console_white_background int = 47 // 控制台 白色 背景
)

//打印
func CommonColorPrint(str string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}

//显示有点异常
func CommonColorPrintWithWhiteBackground(str string, color int) string {
	return fmt.Sprintf("\x1b[%d;%dm%s", Common_console_white_background, color, str)
}
