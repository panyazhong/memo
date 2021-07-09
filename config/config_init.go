package config

func NewDefaultConf() *DatabaseConfig {
	c := &DatabaseConfig{}

	c.Host = "192.168.120.53"
	c.Port = "3306"
	c.User = "root"
	c.Password = "123456"
	c.DbName = "memo"
	c.Charset = "utf8mb4"

	return c
}
