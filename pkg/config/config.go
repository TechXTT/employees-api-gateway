package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT" default:"3001"`
	AuthSvcUrl string `mapstructure:"AUTH_SVC_URL" default:"localhost:50051"`
	DataSvcUrl string `mapstructure:"DATA_SVC_URL" default:"localhost:50052"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("envs/")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
