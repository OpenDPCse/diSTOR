package app

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"sardines/core"
	"sardines/err"
	"sardines/tool"
	"strings"
)

type Mode int

const (
	KW Mode = iota
	CID
)

var (
	ctx    context.Context    = nil
	cancel context.CancelFunc = nil
	btnW   float32            = 500
	btnH   float32            = 40
	btnX   float32            = 350
	// 节点句柄
	hNode      *core.HostNode
	searchMode Mode
)

// 启动按钮
func btnOn() *widget.Button {
	btn := widget.NewButton("启动", func() {
		if cancel != nil || hNode != nil {
			return
		}
		ctx, cancel = context.WithCancel(context.Background())
		msg := make(chan bool, 1)
		go Run(ctx, msg)
		if b := <-msg; b {
			turnOn(hNode.NodeAddr.String())
		} else {
			cancel = nil
		}
	})
	btn.Alignment = widget.ButtonAlignCenter
	btn.Move(fyne.NewPos(btnX, 50))
	btn.Resize(fyne.NewSize(btnW, btnH))
	return btn
}

// 关闭按钮
func btnOff() *widget.Button {
	btn := widget.NewButton("关闭", func() {
		if cancel != nil && hNode != nil {
			Ensure("are you sure?", func(b bool) {
				if b {
					cancel()
					turnOff()
					cancel = nil
				}
			})
		}
	})
	btn.Alignment = widget.ButtonAlignCenter
	btn.Move(fyne.NewPos(btnX, 100))
	btn.Resize(fyne.NewSize(btnW, btnH))
	return btn
}

// 上传文件
func btnUpload() *widget.Button {
	btn := widget.NewButton("上传文件", func() {

		if cancel != nil && hNode != nil {

			uploadWindow()

		} else {
			ShowErr(err.NodeNotStarted)
		}

	})
	btn.Alignment = widget.ButtonAlignCenter
	btn.Move(fyne.NewPos(btnX, 150))
	btn.Resize(fyne.NewSize(btnW, btnH))
	return btn
}

// 检索文件
func search() fyne.CanvasObject {
	c := container.NewVBox()
	c.Move(fyne.NewPos(btnX, 200))
	c.Resize(fyne.NewSize(btnW, btnH*2))

	combo := widget.NewSelect([]string{"Keyword", "CID"}, func(value string) {
		switch value {
		case "CID":
			searchMode = CID
		case "Keyword":
			searchMode = KW

		}
	})
	combo.SetSelectedIndex(0)
	e := widget.NewEntry()

	btnSearch := widget.NewButton("检索文件", func() {
		if cancel != nil && hNode != nil {
			words := e.Text
			words = strings.Trim(words, " ")
			if words == "" || words == " " {
				return
			}
			files := make([]*tool.File, 0, 5)
			switch searchMode {
			case CID:
				f, err2 := hNode.SearchFileByCid(words)
				if err2 != nil {
					ShowErr(err2)
				}
				files = append(files, f)
			case KW:
				f, err2 := hNode.SearchFileByKey(words)
				if err2 != nil {
					ShowErr(err2)
				}
				files = append(files, f...)
			}

			if files != nil && len(files) != 0 {
				subWin := a.NewWindow(e.Text)
				subWin.Resize(fyne.NewSize(900, 600))
				subWin.CenterOnScreen()
				subTab := container.NewDocTabs()

				for i, file := range files {
					tabItem := FileTab(i, file, subWin)
					subTab.Append(tabItem)
				}

				subTab.SetTabLocation(container.TabLocationTop)
				subTab.Resize(fyne.NewSize(890, 590))

				subWin.SetContent(subTab)
				subWin.Show()
			} else {
				ShowInfo("nothing could be found")
			}
		} else {
			ShowErr(err.NodeNotStarted)
		}
	})

	c.Add(e)
	c.Add(combo)
	c.Add(btnSearch)
	return c
}

func PanelTab() fyne.CanvasObject {
	c := container.NewWithoutLayout()

	lab := widget.NewLabel("节点控制")
	lab.Alignment = fyne.TextAlignCenter
	lab.Move(fyne.NewPos(title, 0))

	c.Add(lab)
	c.Add(btnOn())
	c.Add(btnOff())
	c.Add(btnUpload())
	c.Add(search())

	return c
}

func Run(ctx context.Context, msg chan<- bool) {
	h, er := core.GenerateNode()
	if er != nil {
		ShowErr(er)
		hNode = nil
		msg <- false
		return
	}
	if f := <-h.JoinNetwork(); f == 0 {
		ShowErr(err.JoinNetworkFailed)
		hNode = nil
		msg <- false
		return
	}
	hNode = h
	msg <- true

	// keytab distribute
	errC := hNode.KeyTableDistributeOn(false, 15, ctx)
	for err2 := range errC {
		ShowErr(err2)
	}

	hNode.Close()
	hNode = nil
}
