package functions

import (
	"os"
)

func WriteFile(filepath string, data string) error {
	file, err := os.Create(filepath)
	defer file.Close()

	file.WriteString(data)
	return err
}

func ReadFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)

	return string(data), err
}

func SetReadOnly(filepath string) error {
	err := os.Chmod(filepath, 0444)
	return err
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
