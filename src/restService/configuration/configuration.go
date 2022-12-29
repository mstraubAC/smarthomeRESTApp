package configuration

import (
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"go.uber.org/zap"
)

type Config struct {
	RestListener string `mapstructure:"REST_LISTENER" koanf:"rest_listener"`
	DBUrl        string `mapstructure:"DB_URL" koanf:"db_url"`
}

func LoadConfig(logger zap.Logger) (c Config, err error) {
	const envPrefix = "SMARTHOME_"
	var k = koanf.New(".")

	// set defaults
	k.Load(confmap.Provider(map[string]interface{}{
		"rest_listener": ":7777",
		"db_url":        "UNCONFIGURED",
	}, ","), nil)

	// load configuration fom environment file
	if err := k.Load(file.Provider("envs/dev"), dotenv.Parser()); err != nil {
		logger.Error("failed to load configuration from file: " + err.Error())
	}

	// load configuration from environment variables
	k.Load(env.ProviderWithValue(envPrefix, ".", func(s string, v string) (string, interface{}) {
		key := strings.ToLower(strings.TrimPrefix(s, envPrefix))
		return key, v
	}), nil)

	println(k)
	err = k.Unmarshal("", &c)

	return c, err
}
