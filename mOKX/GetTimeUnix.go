package mOKX

func GetKdataTime(TickerKdata map[string][]TypeKd) int64 {
	keys := []string{}

	for key, val := range TickerKdata {
		if len(val) > 0 {
			keys = append(keys, key)
		}
	}
	var TimeUnix int64

	if len(keys) > 0 {
		TimeUnix = TickerKdata[keys[0]][0].TimeUnix
	}

	return TimeUnix
}
