package tmpl

import (
	_ "embed"
)

//go:embed tikker.sh
var TikkerSh string

type TikkerShParam struct {
	Path      string
	ShellCont string
}
