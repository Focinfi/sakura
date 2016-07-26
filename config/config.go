package config

import (
	"os"
	"path"

	"github.com/jinzhu/configor"
)

// DataBase contains psql info
type DataBase struct {
	Adapter  string `default:"postgres"`
	Name     string `default:"sakura_test" env:"POSTGRESQL_INSTANCE_NAME"`
	Host     string `default:"127.0.0.1" env:"POSTGRESQL_PORT_5432_TCP_ADDR"`
	Port     string `default:"5432" env:"POSTGRESQL_PORT_5432_TCP_PORT"`
	User     string `default:"frank" env:"POSTGRESQL_USERNAME"`
	Password string `env:"POSTGRESQL_PASSWORD"`
}

// Redis contains redis info
type Redis struct {
	Name     string `default:"sakura_test"`
	Addr     string `env:"REDIS_PORT_6379_TCP_ADDR" default:"127.0.0.1"`
	Port     string `env:"REDIS_PORT_6379_TCP_PORT" default:"6379"`
	Password string `env:"REDIS_PASSWORD"`
}

// Config contains application configuration
var Config = struct {
	Port       uint `default:"1024" env:"PORT"`
	DataBase   DataBase
	Redis      Redis
	Env        string `default:"test" env:"APP_ENV"`
	BaseSecret string `default:"eslMep28nkvfiYTuWuF7OIp2A5sMb5ewOu2UwO/1PEI=" env:"BASE_SECRET"`
}{}

var (
	// Root is the application root directory
	Root = path.Join(os.Getenv("GOPATH"), "src", "github.com", "focinfi", "sakura")
	// ConfigRoot represents the config directory
	ConfigRoot = Root
)

func init() {
	// use the first arg as the config directory
	if len(os.Args) >= 2 && os.Args[1] != "" {
		ConfigRoot = os.Args[1]
	}

	// load configuration file
	if err := configor.Load(&Config, path.Join(ConfigRoot, "config/config.yml")); err != nil {
		panic(err)
	}
}

// IsProduction check if production environment
func IsProduction() bool {
	return Config.Env == "production"
}
