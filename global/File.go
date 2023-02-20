package global

import (
	_ "embed"
)

//go:embed AppInfo.txt
var AppInfo string

type AppInfoParam struct {
	Version string
}
