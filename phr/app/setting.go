package app

import (
	"encoding/hex"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"sardines/config"
	"sardines/err"
	"strconv"
)

func SettingTab() fyne.CanvasObject {
	c := container.NewWithoutLayout()

	lab := widget.NewLabel("节点配置")
	lab.Alignment = fyne.TextAlignCenter
	lab.Move(fyne.NewPos(title, 0))

	ipEntry := widget.NewEntry()
	ipEntry.SetText("0.0.0.0")
	portEntry := widget.NewEntry()
	portEntry.SetText("8082")
	rsEntry := widget.NewEntry()
	rsEntry.SetText("1")
	bsEntry := widget.NewEntry()

	conf := &config.Config{}
	er := conf.Load()
	if er == nil {
		ipEntry.SetText(conf.IP)
		portEntry.SetText(conf.Port)
		rsEntry.SetText(strconv.FormatInt(conf.RandomSeed, 10))
		bsEntry.SetText(conf.BootstrapNode)
	}

	cForm := widget.NewForm()
	cForm.Append("IP地址", ipEntry)
	cForm.Append("端口", portEntry)
	cForm.Append("私钥种子", rsEntry)
	cForm.Append("引导节点", bsEntry)
	cForm.Move(fyne.NewPos(220, 50))
	cForm.Resize(fyne.NewSize(750, 500))

	// 保存配置
	submit := widget.NewButton("保存", func() {
		conf.IP = ipEntry.Text
		conf.Port = portEntry.Text
		conf.BootstrapNode = bsEntry.Text
		conf.RandomSeed, er = strconv.ParseInt(rsEntry.Text, 10, 64)
		if er != nil {
			ShowErr(er)
			return
		}
		if b := conf.Save(); b {
			ShowInfo("配置成功")
		} else {
			ShowErr(err.ConfFailed)
		}

	})
	submit.Move(fyne.NewPos(title-200, 250))
	submit.Resize(fyne.NewSize(80, 40))

	keyLab := widget.NewLabel("私钥")
	keyLab.Alignment = fyne.TextAlignCenter
	keyLab.Move(fyne.NewPos(title, 300))
	keyEntry := widget.NewEntry()
	keyEntry.Disable()
	keyEntry.MultiLine = true
	keyEntry.Wrapping = fyne.TextWrapBreak
	keyEntry.Move(fyne.NewPos(220, 350))
	keyEntry.Resize(fyne.NewSize(800, 150))

	// 生成密钥对
	genKey := widget.NewButton("生成私钥", func() {
		co := &config.Config{}
		er = co.Load()
		if er != nil {
			ShowErr(er)
			return
		}
		err2 := co.GenKey()
		if err2 != nil {
			ShowErr(err2)
			return
		}
		err2 = co.SaveKey()
		if err2 != nil {
			ShowErr(err2)
			return
		}
		ShowInfo("生成成功")
		raw, _ := co.PrvKey.Raw()
		keyEntry.SetText(hex.EncodeToString(raw))
	})
	genKey.Move(fyne.NewPos(title+100, 250))
	genKey.Resize(fyne.NewSize(80, 40))

	md := widget.NewRichTextFromMarkdown("**请妥善保管您的私钥!**")

	c.Add(lab)
	c.Add(cForm)
	c.Add(submit)
	c.Add(genKey)
	c.Add(keyLab)
	c.Add(keyEntry)
	c.Add(md)

	return c
}
