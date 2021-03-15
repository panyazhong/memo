package config

type DatabaseConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User 	   string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
}
