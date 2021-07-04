package config

import (
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig(configName string) (*Config, error) {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.SetConfigType("dotenv")
	vp.AddConfigPath(".")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	c := &Config{vp}
	go func() {
		vp.WatchConfig()
		vp.OnConfigChange(func(e fsnotify.Event) {
			c.ResetConfig()
		})
	}()

	return c, nil
}

var sections = make(map[string]func(config *Config))

func (c *Config) SetConfig(k string, f func(config *Config)) {
	f(c)

	if _, ok := sections[k]; !ok {
		sections[k] = f
	}
}

func (c *Config) ResetConfig() {
	for _, f := range sections {
		f(c)
	}
}

func (c *Config) get(key string, defaultValue interface{}) interface{} {
	if c.vp.IsSet(key) && c.vp.Get(key) != "" {
		return c.vp.Get(key)
	}
	c.vp.SetDefault(key, defaultValue)

	return defaultValue
}

func (c *Config) GetString(path string, defaultValue string) string {
	return cast.ToString(c.get(path, defaultValue))
}

func (c *Config) GetStringSlice(path string, defaultValue string) []string {
	return strings.Split(cast.ToString(c.get(path, defaultValue)), ",")
}

func (c *Config) GetInt(path string, defaultValue int) int {
	return cast.ToInt(c.get(path, defaultValue))
}

func (c *Config) GetInt64(path string, defaultValue int64) int64 {
	return cast.ToInt64(c.get(path, defaultValue))
}

func (c *Config) GetUint(path string, defaultValue uint) uint {
	return cast.ToUint(c.get(path, defaultValue))
}

func (c *Config) GetBool(path string, defaultValue bool) bool {
	return cast.ToBool(c.get(path, defaultValue))
}

func (c *Config) GetDuration(path string, defaultValue time.Duration) time.Duration {
	return cast.ToDuration(c.get(path, defaultValue))
}
