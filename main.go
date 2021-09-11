package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mattn/go-tty"
)

var Game_Ini = `[/Script/Engine.GameNetworkManager]
TotalNetBandwidth=104857600
MaxDynamicBandwidth=104857600
MinDynamicBandwidth=10485760
`
var Scalability_Ini = `[NetworkQuality@3]
TotalNetBandwidth=104857600
MaxDynamicBandwidth=104857600
MinDynamicBandwidth=10485760
`

var Engine_Ini = `

[/Script/Engine.Player]
ConfiguredInternetSpeed=104857600
ConfiguredLanSpeed=104857600

[/Script/OnlineSubsystemUtils.IpNetDriver]
MaxClientRate=104857600
MaxInternetClientRate=104857600

[/Script/SocketSubsystemEpic.EpicNetDriver]
MaxClientRate=104857600
MaxInternetClientRate=104857600
`
var Input_Ini = `[/script/engine.inputsettings]
ConsoleKey=F6
ConsoleKeys=F6
`

var colorRed = "\033[31m"
var colorYellow = "\033[33m"
var colorGreen = "\033[32m"

func WaitExit() {
	tty, err := tty.Open()
	if err != nil {
		fmt.Println(string(colorRed), err)
	}
	defer tty.Close()
	fmt.Println(string(colorYellow), "\n Press any key to close the terminal.")
	for {
		r, err := tty.ReadRune()
		val := int(r)
		if err != nil {
			log.Fatal(err)
		} else if val != 0 {
			os.Exit(0)
		}
	}
}

func WriteFile(filepath string, data string) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(string(colorRed), "Error creating file data;", err)
		WaitExit()
	}
	defer file.Close()

	file.WriteString(data)
}

func ReadFile(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(string(colorRed), "Error reading file data;", err)
		WaitExit()
	}
	return string(data)
}

func SetReadOnly(filepath string) {
	err := os.Chmod(filepath, 0444)
	if err != nil {
		fmt.Println(string(colorRed), "Error setting file to READONLY.")
		WaitExit()
	}
}

func CanWrite(filepath string) (bool, error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false, err
		}
	}
	file.Close()
	return true, nil

}

func main() {
	cmd := exec.Command("cmd", "/C", "title", "Satisifactory Configuration Client")
	err := cmd.Run()
	if err != nil {
		fmt.Println(string(colorRed), "Error setting title of terminal.")
	}
	cmd2 := exec.Command("cmd", "/C", "mode con:cols=130 lines=20")
	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println(string(colorRed), "Error setting size of terminal.")
	}
	//	fmt.Println(string(colorYellow), ``)
	fmt.Println(string(colorYellow), ` ________  ________  _________  ___  ________  ___  ________ ________  ________ _________  ________  ________      ___    ___ `)
	fmt.Println(string(colorYellow), `|\   ____\|\   __  \|\___   ___\\  \|\   ____\|\  \|\  _____\\   __  \|\   ____\\___   ___\\   __  \|\   __  \    |\  \  /  /|`)
	fmt.Println(string(colorYellow), `\ \  \___|\ \  \|\  \|___ \  \_\ \  \ \  \___|\ \  \ \  \__/\ \  \|\  \ \  \___\|___ \  \_\ \  \|\  \ \  \|\  \   \ \  \/  / /`)
	fmt.Println(string(colorYellow), ` \ \_____  \ \   __  \   \ \  \ \ \  \ \_____  \ \  \ \   __\\ \   __  \ \  \       \ \  \ \ \  \\\  \ \   _  _\   \ \    / / `)
	fmt.Println(string(colorYellow), `  \|____|\  \ \  \ \  \   \ \  \ \ \  \|____|\  \ \  \ \  \_| \ \  \ \  \ \  \____   \ \  \ \ \  \\\  \ \  \\  \|   \/  /  /  `)
	fmt.Println(string(colorYellow), `    ____\_\  \ \__\ \__\   \ \__\ \ \__\____\_\  \ \__\ \__\   \ \__\ \__\ \_______\  \ \__\ \ \_______\ \__\\ _\ __/  / /    `)
	fmt.Println(string(colorYellow), `   |\_________\|__|\|__|    \|__|  \|__|\_________\|__|\|__|    \|__|\|__|\|_______|   \|__|  \|_______|\|__|\|__|\___/ /     `)
	fmt.Println(string(colorYellow), `   \|_________|                        \|_________|                                                              \|___|/      `)
	fmt.Println(string(colorYellow), `------------------------------------------------------------------------------------------------------------------------------`)
	fmt.Println(string(colorYellow), ``)
	fmt.Println(string(colorYellow), ``)
	//	Checking If Already complete

	_, err3 := CanWrite(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Game.ini")
	if err3 != nil {
		fmt.Println(string(colorRed), "Already ran application or file is locked.")
		WaitExit()
	} else {
		// Writing Game.Ini
		WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Game.ini", Game_Ini)
		SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Game.ini")

		// Writing Scalability.Ini
		WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Scalability.ini", Scalability_Ini)
		SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Scalability.ini")

		// Getting Existing Data for Engine.Ini
		data := ReadFile(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Engine.ini")

		WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Engine.ini", data+Engine_Ini)
		SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Engine.ini")

		// Writing Input.ini
		WriteFile(os.Getenv("LOCALAPPDATA")+"\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Input.ini", Input_Ini)
		SetReadOnly(os.Getenv("LOCALAPPDATA") + "\\FactoryGame\\Saved\\Config\\WindowsNoEditor\\Input.ini")
	}
	fmt.Println(string(colorGreen), "Successfully edited all required files! You may close this terminal with any key.")
	tty, err := tty.Open()
	if err != nil {
		fmt.Println(string(colorRed), err)
	}
	defer tty.Close()
	for {
		r, err := tty.ReadRune()
		val := int(r)
		if err != nil {
			log.Fatal(err)
		} else if val != 0 {
			os.Exit(0)
		}
	}
}
