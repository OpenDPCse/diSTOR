package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowErr(e error, win ...fyne.Window) {
	if len(win) == 0 {
		dialog.ShowError(e, w)
	} else {
		dialog.ShowError(e, win[0])
	}
}

func ShowInfo(info string, win ...fyne.Window) {
	if len(win) == 0 {
		dialog.ShowInformation("info", info, w)
	} else {
		dialog.ShowInformation("info", info, win[0])
	}
}

func Ensure(info string, f func(b bool)) {
	dialog.ShowConfirm("ensure", info, f, w)
}

func ShowData(info, data string) {
	e := widget.NewEntry()
	e.SetText(data)
	e.Disable()

	custom := dialog.NewCustom(info, "关闭", e, w)
	custom.Resize(fyne.NewSize(700, 50))
	custom.Show()
}
