package window

import (
	"os"
	"satisfactory_patcher/functions"
	"satisfactory_patcher/ui"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	uiWindow *ui.UIMainMainWindow

	stopCipher bool
}

func NewWindow(ui *ui.UIMainMainWindow) *Window {
	return &Window{
		uiWindow: ui,
	}
}

func (w *Window) Connect() {
	ww := w.uiWindow

	// Setup
	_translate := core.QCoreApplication_Translate

	var imageReader = gui.NewQImageReader3("resources/satisifactory.png", core.NewQByteArray())
	ww.Image.SetPixmap(gui.QPixmap_FromImageReader(imageReader, core.Qt__AutoColor))
	var Completed = false

	data, err := functions.ReadFile(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Game.ini")
	if err != nil {
		widgets.QMessageBox_Critical(nil, "Error", "Error Reading Files", widgets.QMessageBox__Abort, widgets.QMessageBox__Abort)
	} else {
		if data == functions.Game_Ini {
			ww.Status.SetStyleSheet("color: rgb(144, 238, 144);")
			ww.Status.SetText(_translate("MainWindow", "Status: Patched", "", -1))
			Completed = true
		} else {
			ww.Status.SetText(_translate("MainWindow", "Status: Not Patched", "", -1))
			ww.Status.SetStyleSheet("color: rgb(255, 127, 127);")
		}
		ww.PushButton.ConnectClicked(func(checked bool) {
			if Completed == false {
				//  Check ALL File Perms
				// Writing Game.Ini
				functions.WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Game.ini", functions.Game_Ini)
				functions.SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Game.ini")

				// Writing Scalability.Ini
				functions.WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Scalability.ini", functions.Scalability_Ini)
				functions.SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Scalability.ini")

				// Getting Existing Data for Engine.Ini
				data, _ := functions.ReadFile(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Engine.ini")

				functions.WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Engine.ini", data+functions.Engine_Ini)
				functions.SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Engine.ini")

				// Writing Input.ini
				functions.WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Input.ini", functions.Input_Ini)
				functions.SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Input.ini")
				widgets.QMessageBox_Information(nil, "Success!", "Sucessfully Patched!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				ww.Status.SetStyleSheet("color: rgb(144, 238, 144);")
				ww.Status.SetText(_translate("MainWindow", "Status: Patched", "", -1))
				Completed = true
			} else {
				widgets.QMessageBox_Warning(nil, "Warning", "Already Patched!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			}
		})
	}

}
