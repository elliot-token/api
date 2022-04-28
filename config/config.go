package config

type Config struct {
	Server        Server        `mapstructure:"server"`
	SmartContract SmartContract `mapstructure:"smartContract"`
	Database      Database      `mapstructure:"database"`
}

type Server struct {
	Port int64 `mapstructure:"port"`
}

type SmartContract struct {
	Address string `mapstructure:"address"`
}

type Database struct {
	Connection string `mapstructure:"connection"`
}
