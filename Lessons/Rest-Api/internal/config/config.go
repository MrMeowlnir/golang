package config

type Config struct {
	IsDebug *bool
	Listen struct{
		Type string
		BindIp string
		Port string
	}
}