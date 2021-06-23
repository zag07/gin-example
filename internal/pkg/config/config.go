package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
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

	go func() {
		vp.WatchConfig()
		vp.OnConfigChange(func(in fsnotify.Event) {
			// TODO
		})
	}()
}

func get(key string, defaultValue interface{}) interface{} {
	if vp.IsSet(key) && vp.Get(key) != "" {
		return vp.Get(key)
	}
	vp.SetDefault(key, defaultValue)

	return defaultValue
}

func GetString(path string, defaultValue string) string {
	return cast.ToString(get(path, defaultValue))
}

func GetInt(path string, defaultValue int) int {
	return cast.ToInt(get(path, defaultValue))
}

func GetInt64(path string, defaultValue int64) int64 {
	return cast.ToInt64(get(path, defaultValue))
}

func GetUint(path string, defaultValue uint) uint {
	return cast.ToUint(get(path, defaultValue))
}

func GetBool(path string, defaultValue bool) bool {
	return cast.ToBool(get(path, defaultValue))
}
