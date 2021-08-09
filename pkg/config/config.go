package config

import (
	"github.com/spf13/viper"
	"path"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig(file string) (*Config, error) {
	ext := string([]rune(path.Ext(file))[1:])

	vp := viper.New()
	vp.SetConfigName(file)
	vp.SetConfigType(ext)
	vp.AddConfigPath(".")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{vp}, nil
}

func (c *Config) Scan(v interface{}) error {
	return c.vp.Unmarshal(v)
}

func (c *Config) ScanKey(k string, v interface{}) error {
	return c.vp.UnmarshalKey(k, v)
}
