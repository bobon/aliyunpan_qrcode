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
	Win_black        = iota // 黑色
	Win_blue                // 蓝色
	Win_green               // 绿色
	Win_cyan                // 青色
	Win_red                 // 红色
	Win_purple              // 紫色
	Win_yellow              // 黄色
	Win_light_gray          // 淡灰色（系统默认值）
	Win_gray                // 灰色
	Win_light_blue          // 亮蓝色
	Win_light_green         // 亮绿色
	Win_light_cyan          // 亮青色
	Win_light_red           // 亮红色
	Win_light_purple        // 亮紫色
	Win_light_yellow        // 亮黄色
	Win_white               // 白色
)

var (
	kernel32                    = syscall.NewLazyDLL(`kernel32.dll`)
	setConsoleTextAttributeProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	closeHandleProc             = kernel32.NewProc(`CloseHandle`)
)

func WinColorPrint(str string, color int) {
	handle, _, _ := setConsoleTextAttributeProc.Call(uintptr(syscall.Stdout), uintptr(color))
	print(str)
	closeHandleProc.Call(handle)
}
func WinOutMatrix(matrix [][]byte) {
	for k1 := range matrix {
		for k2 := range matrix[k1] {
			if matrix[k1][k2] == 0 {
				WinColorPrint("▇▇", Win_white)
			} else {
				WinColorPrint("▇▇", Win_black)
			}
		}
		println()
	}
}
