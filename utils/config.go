package utils

import "github.com/spf13/viper"

// Config is the struct that contains the configuration of the application
// The values are loaded from environment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type TestConfig struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE_TEST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // can be json, yaml, toml or env

	viper.AutomaticEnv() // read values from environment variables

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		return
	}

	err = viper.Unmarshal(&config) // Unmarshal config into struct
	if err != nil {
		return
	}

	return
}

func LoadTestConfig(path string) (config TestConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // can be json, yaml, toml or env

	viper.AutomaticEnv() // read values from environment variables

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		return
	}

	err = viper.Unmarshal(&config) // Unmarshal config into struct
	if err != nil {
		return
	}

	return
}
