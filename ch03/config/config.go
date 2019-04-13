package config

type MysqlConfig struct {
	ServerName string
	ServerPort int
	UserName   string
	Password   string
	DbName     string
	Charset    string
}
