package config

type Config struct {
	ListenAddr   string
	Debug        bool
	DatabasePath string
}

var Default = &Config{
	ListenAddr:   ":8080",
	Debug:        true,
	DatabasePath: "gorm.db",
}
