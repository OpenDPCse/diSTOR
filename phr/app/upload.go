package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func uploadWindow() {
	uploadWin := a.NewWindow("上传文件")
	uploadWin.CenterOnScreen()
	uploadWin.Resize(fyne.NewSize(900, 600))

	c := container.NewWithoutLayout()
	c.Resize(fyne.NewSize(890, 590))

	fileNameEntry := widget.NewEntry()
	fileNameEntry.SetPlaceHolder("文件名")
	fileNameEntry.Resize(fyne.NewSize(850, btnH))
	fileNameEntry.Move(fyne.NewPos(0, 0))
	fileNameEntry.Disable()

	filePathEntry := widget.NewEntry()
	filePathEntry.SetPlaceHolder("文件路径")
	filePathEntry.Resize(fyne.NewSize(850, btnH))
	filePathEntry.Move(fyne.NewPos(0, 50))
	filePathEntry.Disable()

	selectFileBtn := widget.NewButton("选择文件", func() {
		fileDialog := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err2 error) {
			if closer != nil {
				p := closer.URI().Path()
				n := closer.URI().Name()

				fileNameEntry.SetText(n)
				filePathEntry.SetText(p)

			}
		}, uploadWin)
		fileDialog.Resize(fyne.NewSize(800, 540))
		fileDialog.Show()
	})
	selectFileBtn.Alignment = widget.ButtonAlignCenter
	selectFileBtn.Resize(fyne.NewSize(btnW, btnH))
	selectFileBtn.Move(fyne.NewPos(btnX-130, 100))

	fileIndexEntry := widget.NewEntry()
	fileIndexEntry.SetPlaceHolder("文件索引")
	fileIndexEntry.Resize(fyne.NewSize(850, btnH))
	fileIndexEntry.Move(fyne.NewPos(0, 150))

	uploadFileBtn := widget.NewButton("上传", func() {
		defer uploadWin.Close()
		fid, e := hNode.UploadFile(filePathEntry.Text, fileNameEntry.Text, fileIndexEntry.Text)
		if e != nil {
			ShowErr(e)
			return
		}
		ShowData("上传成功", fid+":"+fileIndexEntry.Text)
		fileTree.Refresh()
	})
	uploadFileBtn.Alignment = widget.ButtonAlignCenter
	uploadFileBtn.Resize(fyne.NewSize(btnW, btnH))
	uploadFileBtn.Move(fyne.NewPos(btnX-130, 200))

	c.Add(fileNameEntry)
	c.Add(filePathEntry)
	c.Add(selectFileBtn)
	c.Add(fileIndexEntry)
	c.Add(uploadFileBtn)

	uploadWin.SetContent(c)
	uploadWin.Show()

}
