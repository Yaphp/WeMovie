package config

// Database DatabaseConfig 数据库配置
type Database struct {
	Dsn    string
	Prefix string
}

// Config ServerConfig 服务器配置
type Config struct {
	Database Database
}
