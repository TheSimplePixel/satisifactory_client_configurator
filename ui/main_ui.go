// WARNING! All changes made in this file will be lost!
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
)

type UIMainMainWindow struct {
	Centralwidget *widgets.QWidget
	MainFrame *widgets.QFrame
	Image *widgets.QLabel
	ButtonBg *widgets.QFrame
	PushButton *widgets.QPushButton
	Footer *widgets.QFrame
	LabelCredits *widgets.QLabel
	Status *widgets.QLabel
	Version *widgets.QLabel
}

func (this *UIMainMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 630, 292))
	MainWindow.SetMinimumSize(core.NewQSize2(630, 292))
	MainWindow.SetMaximumSize(core.NewQSize2(630, 292))
	MainWindow.SetStyleSheet("QMainWindow {background: transparent; }\nQToolTip {\n\tbackground-color: rgba(27, 29, 35, 160);\n\tborder: 1px solid rgb(40, 40, 40);\n\tborder-radius: 2px;\n}")
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.MainFrame = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.MainFrame.SetObjectName("MainFrame")
	this.MainFrame.SetGeometry(core.NewQRect4(10, 10, 611, 261))
	this.MainFrame.SetStyleSheet("background-color: rgb(41, 45, 56);\nborder-radius: 5px;")
	this.MainFrame.SetFrameShape(widgets.QFrame__StyledPanel)
	this.MainFrame.SetFrameShadow(widgets.QFrame__Raised)
	this.Image = widgets.NewQLabel(this.MainFrame, core.Qt__Widget)
	this.Image.SetObjectName("Image")
	this.Image.SetGeometry(core.NewQRect4(100, 50, 421, 91))
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetFamily("Microsoft YaHei UI")
	font.SetPointSize(24)
	font.SetWeight(50)
	this.Image.SetFont(font)
	this.Image.SetScaledContents(true)
	this.ButtonBg = widgets.NewQFrame(this.MainFrame, core.Qt__Widget)
	this.ButtonBg.SetObjectName("ButtonBg")
	this.ButtonBg.SetGeometry(core.NewQRect4(50, 180, 511, 51))
	this.ButtonBg.SetStyleSheet("background-color: rgb(38, 42, 52);\nborder-radius: 5px;")
	this.ButtonBg.SetFrameShape(widgets.QFrame__StyledPanel)
	this.ButtonBg.SetFrameShadow(widgets.QFrame__Raised)
	this.PushButton = widgets.NewQPushButton(this.ButtonBg)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetGeometry(core.NewQRect4(10, 10, 493, 31))
	this.PushButton.SetMinimumSize(core.NewQSize2(150, 30))
	var palette *gui.QPalette
	var brush *gui.QBrush
	palette = gui.NewQPalette()
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Active, gui.QPalette__WindowText, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Active, gui.QPalette__Button, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Active, gui.QPalette__Text, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Active, gui.QPalette__ButtonText, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Active, gui.QPalette__Base, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Active, gui.QPalette__Window, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Inactive, gui.QPalette__WindowText, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Inactive, gui.QPalette__Button, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Inactive, gui.QPalette__Text, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Inactive, gui.QPalette__ButtonText, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Inactive, gui.QPalette__Base, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Inactive, gui.QPalette__Window, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Disabled, gui.QPalette__WindowText, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Disabled, gui.QPalette__Button, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Disabled, gui.QPalette__Text, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(255, 255, 255, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Disabled, gui.QPalette__ButtonText, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Disabled, gui.QPalette__Base, brush)
	brush = gui.NewQBrush3(gui.NewQColor3(52, 59, 72, 255), core.Qt__SolidPattern)
	palette.SetBrush2(gui.QPalette__Disabled, gui.QPalette__Window, brush)
	this.PushButton.SetPalette(palette)
	font = gui.NewQFont()
	font.SetFamily("Segoe UI")
	font.SetPointSize(9)
	this.PushButton.SetFont(font)
	this.PushButton.SetStyleSheet("QPushButton {\n\tborder: 2px solid rgb(52, 59, 72);\n\tborder-radius: 5px;\t\n\tbackground-color: rgb(52, 59, 72);\n\tcolor: rgb(255, 255, 255);\n}\nQPushButton:hover {\n\tbackground-color: rgb(57, 65, 80);\n\tborder: 2px solid rgb(61, 70, 86);\n}\nQPushButton:pressed {\t\n\tbackground-color: rgb(35, 40, 49);\n\tborder: 2px solid rgb(43, 50, 61);\n}")
	this.PushButton.SetAutoDefault(false)
	this.ButtonBg.Raise()
	this.Image.Raise()
	this.Footer = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.Footer.SetObjectName("Footer")
	this.Footer.SetGeometry(core.NewQRect4(10, 260, 611, 25))
	this.Footer.SetMinimumSize(core.NewQSize2(0, 25))
	this.Footer.SetMaximumSize(core.NewQSize2(16777215, 25))
	this.Footer.SetStyleSheet("background-color: rgb(33, 37, 43);")
	this.Footer.SetFrameShape(widgets.QFrame__NoFrame)
	this.Footer.SetFrameShadow(widgets.QFrame__Raised)
	this.LabelCredits = widgets.NewQLabel(this.Footer, core.Qt__Widget)
	this.LabelCredits.SetObjectName("LabelCredits")
	this.LabelCredits.SetGeometry(core.NewQRect4(10, 0, 221, 21))
	font = gui.NewQFont()
	font.SetFamily("Segoe UI")
	this.LabelCredits.SetFont(font)
	this.LabelCredits.SetStyleSheet("color: rgb(180, 180, 180);")
	this.Status = widgets.NewQLabel(this.Footer, core.Qt__Widget)
	this.Status.SetObjectName("Status")
	this.Status.SetGeometry(core.NewQRect4(370, 0, 161, 21))
	font = gui.NewQFont()
	font.SetFamily("Segoe UI")
	this.Status.SetFont(font)
	this.Status.SetStyleSheet("color: rgb(180, 180, 180);")
	this.Version = widgets.NewQLabel(this.Footer, core.Qt__Widget)
	this.Version.SetObjectName("Version")
	this.Version.SetGeometry(core.NewQRect4(570, 0, 31, 21))
	font = gui.NewQFont()
	font.SetFamily("Segoe UI")
	this.Version.SetFont(font)
	this.Version.SetStyleSheet("color: rgb(180, 180, 180);")
	MainWindow.SetCentralWidget(this.Centralwidget)


    this.RetranslateUi(MainWindow)

}

func (this *UIMainMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "Satisfactory Patcher", "", -1))
	this.Image.SetText(_translate("MainWindow", "", "", -1))
	this.PushButton.SetText(_translate("MainWindow", "Patch Satisfactory", "", -1))
	this.LabelCredits.SetText(_translate("MainWindow", "Registered by: github.com/TheSimplePixel", "", -1))
	this.Status.SetText(_translate("MainWindow", "Status: Patched", "", -1))
	this.Version.SetText(_translate("MainWindow", "v1.0.0", "", -1))
}
