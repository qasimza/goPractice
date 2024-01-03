package main

import (
	"fmt"

	"github.com/schollz/wifiscan"
)

type WifiInformation struct {
	SSID string `json:"SSID"`
	RSSI int    `json:"RSSI"`
}

func getWifiInformation() {
	var wifiList []WifiInformation
	wifis, err := wifiscan.Scan()
	if err != nil {
		fmt.Println(err)
	}
	for _, w := range wifis {
		fmt.Println(w.SSID, w.RSSI)
	}
}

func main() {

	getWifiInformation()
}
