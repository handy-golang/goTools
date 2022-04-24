package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")
	message := "ab123cdef"
	secretKey := "meicha44ngliang"
	shaStr := mEncrypt.Sha256(message, secretKey)

	fmt.Println("shaStr", shaStr)

	fmt.Println(" =========   END   ========= ")
}

/*
1L6olEDHzaUAHg2r1yJeYiyasmLbtA98pkPChumdWGI=
1L6olEDHzaUAHg2r1yJeYiyasmLbtA98pkPChumdWGI=



nHQQuqFyWopraPbQgBUChWncXhPAclWlOfOHLnPDG0c=
nHQQuqFyWopraPbQgBUChWncXhPAclWlOfOHLnPDG0c=


*/
