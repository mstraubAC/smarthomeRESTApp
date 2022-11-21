package configuration

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	RestListener string `mapstructure:"REST_LISTENER"`
	DBUrl        string `mapstructure:"DB_URL"`
}

func LoadConfig(logger zap.Logger) (c Config, err error) {
	viper.SetEnvPrefix("SMARTHOME")
	viper.AddConfigPath("../envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logger.Error("Failed to read configuration: " + err.Error())
		return c, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		logger.Error("Failed to unmarshal configuration: " + err.Error())
		return c, err
	}

	return c, err
}
