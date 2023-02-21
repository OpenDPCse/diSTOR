package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"path/filepath"
	"sardines/tool"
	"strconv"
	"strings"
)

func FileTab(index int, file *tool.File, win fyne.Window) *container.TabItem {

	c := container.NewWithoutLayout()
	c.Resize(fyne.NewSize(890, 590))

	originEntry := widget.NewEntry()
	originEntry.SetText(file.Origin)
	originEntry.Disable()
	originEntry.Resize(fyne.NewSize(890, btnH))
	originEntry.Move(fyne.NewPos(0, 0))

	cidEntry := widget.NewEntry()
	cidEntry.SetText(file.CID)
	cidEntry.Disable()
	cidEntry.Resize(fyne.NewSize(890, btnH))
	cidEntry.Move(fyne.NewPos(0, 50))

	var contentEntry *widget.Entry
	if file.IsText() {
		contentEntry = widget.NewEntry()
		contentEntry.SetText(string(file.Content))
		contentEntry.Disable()
	} else {
		contentEntry = widget.NewEntry()
		contentEntry.SetText("该文件不可预览")
		contentEntry.Disable()
	}
	contentEntry.Resize(fyne.NewSize(890, 390))
	contentEntry.Move(fyne.NewPos(0, 100))

	download := widget.NewButton("下载文件", func() {
		folderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			path := uri.Path()

			name := strings.Split(file.Origin, ":")[0]
			filePath := filepath.Join(path, name)
			er := tool.WriteFile(file.Content, filePath)
			if er != nil {
				ShowErr(er, win)
			} else {
				ShowInfo("下载成功", win)
			}

		}, win)
		folderDialog.Resize(fyne.NewSize(800, 540))
		folderDialog.Show()
	})
	download.Resize(fyne.NewSize(btnW, btnH))
	download.Move(fyne.NewPos(btnX-130, 505))

	c.Add(originEntry)
	c.Add(cidEntry)
	c.Add(contentEntry)
	c.Add(download)

	return container.NewTabItem("file"+strconv.Itoa(index), c)
}
