package qr_windows

import (
	"syscall"
)

//https://docs.microsoft.com/en-us/windows/console/console-screen-buffers#character-attributes
/*  wincon.h
#define FOREGROUND_BLUE      0x0001 // text color contains blue.
#define FOREGROUND_GREEN     0x0002 // text color contains green.
#define FOREGROUND_RED       0x0004 // text color contains red.
#define FOREGROUND_INTENSITY 0x0008 // text color is intensified.
#define BACKGROUND_BLUE      0x0010 // background color contains blue.
#define BACKGROUND_GREEN     0x0020 // background color contains green.
#define BACKGROUND_RED       0x0040 // background color contains red.
#define BACKGROUND_INTENSITY 0x0080 // background color is intensified.
*/
const (
	black        = iota // 黑色
	blue                // 蓝色
	green               // 绿色
	cyan                // 青色
	red                 // 红色
	purple              // 紫色
	yellow              // 黄色
	light_gray          // 淡灰色（系统默认值）
	gray                // 灰色
	light_blue          // 亮蓝色
	light_green         // 亮绿色
	light_cyan          // 亮青色
	light_red           // 亮红色
	light_purple        // 亮紫色
	light_yellow        // 亮黄色
	white               // 白色
)

const (
	charsetPrint = "▇▇"
)

var (
	kernel32                    = syscall.NewLazyDLL(`kernel32.dll`)
	setConsoleTextAttributeProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	closeHandleProc             = kernel32.NewProc(`CloseHandle`)
)

func colorPrint(str string, color int) {
	handle, _, _ := setConsoleTextAttributeProc.Call(uintptr(syscall.Stdout), uintptr(color))
	print(str)
	closeHandleProc.Call(handle)
}

//
func WinOutMatrix(matrix [][]byte) {
	for k1 := range matrix {
		for k2 := range matrix[k1] {
			if matrix[k1][k2] == 0 {
				colorPrint(charsetPrint, white)
			} else {
				colorPrint(charsetPrint, black)
			}
		}
		println()
	}
}
