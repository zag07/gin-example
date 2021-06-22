package config

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var vp *viper.Viper

func init() {
	vp = viper.New()
	vp.SetConfigName(".env")
	vp.SetConfigType("env")
	vp.AddConfigPath(".")

	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func Get(key string, defaultValue ...interface{}) interface{} {
	if vp.IsSet(key) && vp.Get(key) != "" {
		return vp.Get(key)
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}

func GetConfig(k string, v interface{}) error {
	err := vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
