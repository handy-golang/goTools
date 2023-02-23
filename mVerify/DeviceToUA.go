package mVerify

import (
	useragent "github.com/wenlng/go-user-agent"
)

type DeviceInfo struct {
	BrowserName string
	OsName      string
}

func DeviceToUA(ua string) DeviceInfo {
	BrowserName := useragent.GetBrowserName(ua)
	OsName := useragent.GetOsName(ua)
	return DeviceInfo{
		BrowserName: BrowserName,
		OsName:      OsName,
	}
}
