package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func init() {
	icon, _ = fyne.LoadResourceFromPath("./rsrc/4.jpg")
}

var (
	CX        float32 = 1030
	CY        float32 = 650
	CD        float32 = 27
	title     float32 = 600
	w         fyne.Window
	circleOn  *canvas.Circle
	circleOff *canvas.Circle
	hostEntry *widget.Entry
	icon      fyne.Resource
	a         fyne.App
)

func App() {

	a = app.New()

	a.Settings().SetTheme(theme.LightTheme())

	w = a.NewWindow("sardines")
	w.Resize(fyne.Size{Width: 1280, Height: 720})
	w.CenterOnScreen()
	//w.SetFixedSize(true)
	w.SetIcon(icon)

	c := container.NewWithoutLayout()

	hostEntry = widget.NewEntry()
	hostEntry.Disable()
	hostEntry.SetPlaceHolder("未启动")
	hostInfo := widget.NewForm()
	hostInfo.Resize(fyne.NewSize(770, 30))
	hostInfo.Move(fyne.NewPos(CX-800, CY-5))
	hostInfo.Append("ID", hostEntry)
	initCircle()

	tabs := container.NewAppTabs(
		container.NewTabItem("控制面板", PanelTab()),
		container.NewTabItem("文件", FilesTreeTab()),
		container.NewTabItem("对等点", PeersTab()),
		container.NewTabItem("配置", SettingTab()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.Resize(fyne.Size{
		Width:  1080,
		Height: 650,
	})

	c.Add(circleOn)
	c.Add(circleOff)
	c.Add(tabs)
	c.Add(hostInfo)

	w.SetContent(c)
	w.ShowAndRun()
}

func initCircle() {
	circleOn = canvas.NewCircle(color.RGBA{
		R: 69,
		G: 196,
		B: 190,
		A: 255,
	})
	circleOn.StrokeColor = color.Gray{Y: 0x99}
	circleOn.StrokeWidth = 2
	circleOn.Position1 = fyne.NewPos(CX, CY)
	circleOn.Position2 = fyne.NewPos(CX+CD, CY+CD)

	circleOff = canvas.NewCircle(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})
	circleOff.StrokeColor = color.Gray{Y: 0x99}
	circleOff.StrokeWidth = 2
	circleOff.Position1 = fyne.NewPos(CX, CY)
	circleOff.Position2 = fyne.NewPos(CX+CD, CY+CD)

	turnOff()
}

func turnOn(host string) {
	circleOn.Show()
	circleOff.Hide()
	hostEntry.SetText(host)
}

func turnOff() {
	circleOff.Show()
	circleOn.Hide()
	hostEntry.SetText("")
}
