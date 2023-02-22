package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"sardines/storage"
)

var (
	fileTree *widget.Tree
)

func FilesTreeTab() fyne.CanvasObject {
	c := container.NewWithoutLayout()

	lab := widget.NewLabel("文件树")
	lab.Alignment = fyne.TextAlignCenter
	lab.Move(fyne.NewPos(title, 0))

	tree := widget.NewTree(childUIDs, isBranch, createNode, updateNode)
	tree.Move(fyne.NewPos(220, 50))
	tree.Resize(fyne.NewSize(750, 500))
	tree.OpenAllBranches()
	fileTree = tree

	refBtn := widget.NewButton("刷新", func() {
		tree.Refresh()
	})
	refBtn.Resize(fyne.NewSize(40, 20))
	refBtn.Move(fyne.NewPos(1000, 60))

	c.Add(tree)
	c.Add(lab)
	c.Add(refBtn)
	return c
}

func childUIDs(uid string) (c []string) {
	data := storage.FileStoreTree()
	c = data[uid]
	return
}
func isBranch(uid string) (b bool) {
	data := storage.FileStoreTree()
	_, b = data[uid]
	return
}
func createNode(branch bool) fyne.CanvasObject {
	e := widget.NewEntry()
	e.Disable()
	e.TextStyle = fyne.TextStyle{
		Bold:      false,
		Italic:    false,
		Monospace: true,
		Symbol:    false,
		TabWidth:  0,
	}

	return e
}
func updateNode(uid string, branch bool, node fyne.CanvasObject) {
	node.(*widget.Entry).SetText(uid)
}
