package config

import "github.com/spf13/viper"

var vp *viper.Viper

func init() {
	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("configs/")

	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}
}

// TODO config 文件中的 k 未设置怎么办？
func GetConfig(k string, v interface{}) error {
	err := vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
