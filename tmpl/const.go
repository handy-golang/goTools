package tmpl

import (
	_ "embed"
)

//go:embed tikker.sh
var TikkerSh string

type TikkerShParam struct {
	Path      string
	FileName  string
	ShellCont string
	LogPath   string
}

//go:embed instPm2.sh
var InstPm2 string

type InstPm2Param struct {
	Path     string
	FileName string
}
