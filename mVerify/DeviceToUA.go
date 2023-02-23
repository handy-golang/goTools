package mVerify

import (
	"fmt"

	useragent "github.com/wenlng/go-user-agent"
)

type DeviceInfo struct {
	BrowserName string
	OsName      string
}

func DeviceToUA(ua string) DeviceInfo {
	BrowserName := useragent.GetBrowserName(ua)
	fmt.Println("BrowserName", BrowserName)

	OsName := useragent.GetOsName(ua)
	fmt.Println("OsName", OsName)

	return DeviceInfo{
		BrowserName: BrowserName,
		OsName:      OsName,
	}
}
