package mStr

func Fuzzy(cont any) string {
	str := ToStr(cont)
	Len := len(str)

	if Len < 2 {
		return "*"
	}

	xlen := Len / 3

	if xlen < 1 {
		return Join(
			str[0], '*',
		)
	}

	startStr := str[0:xlen]
	endStr := str[Len-xlen : Len]

	centerStr := ""
	for i := 0; i < xlen; i++ {
		centerStr += "*"
	}

	resData := Join(
		startStr,
		centerStr,
		endStr,
	)

	return resData
}

func GetKeyFuzzy(Ket string, startIdx, endIdx int) string {
	return Ket[0:startIdx] + "******" + Ket[len(Ket)-endIdx:]
}
