package config

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int16  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	Timeout  string `json:"timeout"`
}

type WechatConfig struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}
