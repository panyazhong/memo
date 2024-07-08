package config

func NewDefaultConf() *DatabaseConfig {
	c := &DatabaseConfig{}

	c.Host = "127.0.0.1"
	c.Port = 3306
	c.User = "root"
	c.Password = "123456"
	c.DbName = "PotatoMenu"
	c.Charset = "utf8mb4"
	c.Timeout = "10s"

	return c
}

func InitWechatConfig() *WechatConfig {
	c := &WechatConfig{}
	c.Appid = "wx830953df29a79fb9"
	c.Secret = "da77b2e0276ca7b339f30a75677de00e"

	return c
}
