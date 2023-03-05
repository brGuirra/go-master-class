package utils

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by vyper from a config file or env variables
type Config struct {
	DBDriver      string `mapstructure:"DATABASE_DRIVER"`
	DBSource      string `mapstructure:"DATABASE_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads the configuration from file or environment variables
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
