package mOKX

func GetKdataTime(TickerKdata map[string][]TypeKd) int64 {
	keys := []string{}

	for key, val := range TickerKdata {
		if len(val) > 1 {
			keys = append(keys, key)
		}
	}
	var TimeUnix int64

	if len(keys) > 0 {
		key := keys[len(keys)-1]
		KdataList := TickerKdata[key]
		TimeUnix = KdataList[len(KdataList)-1].TimeUnix
	}

	return TimeUnix
}
