package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mFile"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	fileName := mFile.GetName(mFile.GetNameOpt{
		FileName: "package.json",
		SavePath: ".",
	})

	fmt.Println(fileName)

	fmt.Println(" =========   END   ========= ")
}
