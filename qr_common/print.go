package qr_common

import (
	"fmt"
)

//通用
const (
	common_black                    byte = 1  //二值化存储 黑景色值
	common_white                    byte = 0  //二值化存储 白景色值
	common_console_black_background int  = 40 // 控制台 黑色 背景
	common_console_white_background int  = 47 // 控制台 白色 背景
	common_console_black_foreground int  = 30 // 控制台 黑色 前景
	common_console_white_foreground int  = 37 // 控制台 白色 前景
)

func sprintfColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}

func printColor(color int, str string) string {
	if color == 47 {
		return "\033[40;40m  \033[0m"
	} else {
		return "\033[47;30m  \033[0m"
	}
}
