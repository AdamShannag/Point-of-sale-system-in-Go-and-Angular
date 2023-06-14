package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	PrivateKeyFilePath   string        `mapstructure:"PRIVATE_KEY_FILE_PATH"`
	PublicKeyFilePath    string        `mapstructure:"PUBLIC_KEY_FILE_PATH"`
	RsaKeysBits          int           `mapstructure:"RSA_KEYS_BITS"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	AppTimeZone          string        `mapstructure:"APP_TIME_ZONE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
