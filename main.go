package main

import (
	"os"
	"satisfactory_patcher/ui"
	"satisfactory_patcher/window"

	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	ui.UIMainMainWindow
	Widget *widgets.QMainWindow
}

func NewWidget(parent widgets.QWidget_ITF) *Window {
	window := &Window{
		Widget: widgets.NewQMainWindow(parent, core.Qt__Window),
	}

	window.SetupUI(window.Widget)
	return window
}

func main() {
	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetWindowIcon(gui.NewQIcon5("resources/smart_splitter.ico"))
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetMinimumSize2(250, 200)

	uiWindow := &ui.UIMainMainWindow{}
	uiWindow.SetupUI(mainWindow)

	w := window.NewWindow(uiWindow)
	w.Connect()

	// make the mainWindow visible
	mainWindow.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the mainWindow is closed by the user
	app.Exec()
}
