package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PeersTab() fyne.CanvasObject {
	c := container.NewWithoutLayout()

	lab := widget.NewLabel("对等点列表")
	lab.Alignment = fyne.TextAlignCenter
	lab.Move(fyne.NewPos(title, 0))

	list := widget.NewList(
		length,
		func() fyne.CanvasObject {
			e := widget.NewEntry()
			e.Disable()
			return e
		},
		updateItem)
	list.Move(fyne.NewPos(220, 50))
	list.Resize(fyne.NewSize(750, 500))

	btnUP := widget.NewButton(`↑`, func() {
		list.ScrollToTop()
	})
	btnUP.Alignment = widget.ButtonAlignCenter
	btnUP.Move(fyne.NewPos(1000, 60))
	btnUP.Resize(fyne.NewSize(30, 20))

	c.Add(lab)
	c.Add(list)
	c.Add(btnUP)

	return c
}

func length() int {
	if hNode != nil {
		//fmt.Println(hNode.Router.Sum())
		return hNode.Router.Sum()
	} else {
		return 1
	}
}

func updateItem(i widget.ListItemID, o fyne.CanvasObject) {
	str := ""
	if hNode != nil {
		nodes := hNode.Router.AllNodes()
		//fmt.Println(nodes)
		if nodes[i] != nil {
			str = nodes[i].String()
		}
	}
	o.(*widget.Entry).SetText(str)
}
