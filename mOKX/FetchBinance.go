package mOKX

type TypeBinanceKey struct {
	Name      string `bson:"Name"`
	ApiKey    string `bson:"ApiKey"`
	SecretKey string `bson:"SecretKey"`
	IsTrade   bool   `bson:"IsTrade"`
	UserID    string `bson:"UserID"`
}

type OptFetchBinance struct {
	Path          string
	Data          map[string]any
	Method        string
	IsLocalJson   bool
	LocalJsonPath string // 本地的数据源
	Event         func(string, any)
	BinanceKey    TypeBinanceKey
}
