package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"os"
	"raselper"
	"strings"
)

var nameWillSet = "RasltyMing" // 将要设置的Git显示名称

func init() {
	//设置中文字体:解决中文乱码问题
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	// 查看当前用户名
	err, outStr, errStr := raselper.RunCmd("git config user.name")
	if err != nil || errStr != "" {
		fmt.Println(err)
		fmt.Println(errStr)
		return
	}

	// 更改用户名
	if strings.Contains(outStr, "RasltyMing") { // 确认要更改为哪个用户名
		nameWillSet = "chenming"
	}
	err, _, errStr = raselper.RunCmd("git config --global user.name " + nameWillSet)
	if err != nil || errStr != "" {
		fmt.Println(err)
		fmt.Println(errStr)
		return
	}

	// 最后查看当前用户名
	_, outStr, _ = raselper.RunCmd("git config user.name")
	myApp := app.New()
	myWindow := myApp.NewWindow("成功")
	myWindow.SetContent(widget.NewLabel("新的用户名: " + outStr))
	myWindow.Resize(fyne.NewSize(335, 120))
	myWindow.ShowAndRun()
}
