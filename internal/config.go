package internal

import "github.com/spf13/viper"

type Config struct {
	Root   string
	Config string
}

func (c *Config) init() error {
	viper.AddConfigPath(c.Root)
	viper.SetConfigName(c.Config)

	return viper.ReadInConfig()
}

func NewConfig() (*Config, error) {
	config := Config{
		Root:   "configs",
		Config: "config",
	}

	err := config.init()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Get(key string) string {
	return viper.GetString(key)
}

func (c *Config) GetSlice(key string) []string {
	return viper.GetStringSlice(key)
}
