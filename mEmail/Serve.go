package mEmail

type ServeType struct {
	Account  string
	Password string
	Port     string
	Host     string
}

func Gmail(account, password string) ServeType {
	return ServeType{
		Account:  account,
		Password: password,
		Port:     "587",
		Host:     "smtp.gmail.com",
	}
}

func WorkWeiXin(account, password string) ServeType {
	return ServeType{
		Account:  account,
		Password: password,
		Port:     "587",
		Host:     "smtp.exmail.qq.com",
	}
}

func QQ(account, password string) ServeType {
	return ServeType{
		Account:  account,
		Password: password,
		Port:     "587",
		Host:     "smtp.qq.com",
	}
}
