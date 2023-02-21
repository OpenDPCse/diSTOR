package main

import (
	"golang.org/x/sys/windows"
	"os"
	"sardines/app"
	"sardines/config"
	"sardines/err"
	"sardines/tool"
	"syscall"
)

func init() {
	os.Setenv("FYNE_FONT", "C:/Windows/Fonts/msyh.ttc")
}

func Error(e error) {
	p, _ := syscall.UTF16PtrFromString(e.Error())
	windows.MessageBox(0, p, p, windows.EVENTLOG_SUCCESS)
}

func main() {

	if er := tool.CreateDir(config.Dir); er != nil && er != err.DirExists {
		Error(er)
		return
	}

	if er := tool.CreateDir(config.Downloads); er != nil && er != err.DirExists {
		Error(er)
		return
	}

	app.App()

	defer os.Unsetenv("FYNE_FONT")
}
