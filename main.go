package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	for i := 0; i < 10; i++ {
		str := mEncrypt.RandStr(8)
		fmt.Println(str)

	}

	fmt.Println(" =========   END   ========= ")
}

/*
1L6olEDHzaUAHg2r1yJeYiyasmLbtA98pkPChumdWGI=
1L6olEDHzaUAHg2r1yJeYiyasmLbtA98pkPChumdWGI=



nHQQuqFyWopraPbQgBUChWncXhPAclWlOfOHLnPDG0c=
nHQQuqFyWopraPbQgBUChWncXhPAclWlOfOHLnPDG0c=


*/
