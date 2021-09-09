package config

import (
	"bytes"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/zag07/gin-example/internal/conf"
)

func load(content string, isPath bool) (*conf.Bootstrap, error) {
	config := &conf.Bootstrap{}
	var err error

	if isPath == true {
		viper.SetConfigFile(content)
		err = viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	} else {
		viper.SetConfigType("json")

		err = viper.ReadConfig(bytes.NewBuffer([]byte(content)))

		if err != nil {
			return nil, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			if err = viper.Unmarshal(&config); err != nil {
				zap.L().Error("配置变更发生错误")
			}
		})
	}()

	return config, nil
}

// Load creates a Config struct from a config file path
func Load(path string) (*conf.Bootstrap, error) {
	return load(path, true)
}

// LoadFromContent creates a Config struct from a config content
func LoadFromContent(content string) (*conf.Bootstrap, error) {
	return load(content, false)
}

// DefaultConfig returns a default config instance
func DefaultConfig() *conf.Bootstrap {
	return &conf.Bootstrap{
		Trace: &conf.Trace{
			Endpoint: "http://127.0.0.1:14268/api/traces",
		},
		Http: &conf.HTTP{
			Name:                 "gin-example",
			Port:                 ":8080",
			Timeout:              &durationpb.Duration{Seconds: 1},
			Debug:                false,
			PageSize:             25,
			MaxPageSize:          100,
			UploadSavePath:       "storage/app/uploads",
			UploadServerUrl:      "http://127.0.0.1:8080/static",
			UploadImageMaxSize:   5,
			UploadImageAllowExts: []string{".jpg", ".jpeg", ".png"},
			TraceName:            "gin-example",
			TracePort:            "127.0.0.1:6831",
			WsWriteWait:          &durationpb.Duration{Seconds: 10},
			WsPongWait:           &durationpb.Duration{Seconds: 60},
			WsMaxMessageSize:     512,
			WsMessageQueue:       1024,
			WsOfflineNum:         10,
			JwtSecret:            "echo",
			JwtIssuer:            "gin-example",
			JwtExpire:            &durationpb.Duration{Seconds: 7200},
		},
		Data: &conf.Data{
			Database: &conf.Data_Database{
				Driver: "mysql",
				Source: "root:root@tcp(127.0.0.1:33061)/gin_example2?charset=utf8mb4&parseTime=True&loc=Local",
			},
			Redis: &conf.Data_Redis{
				Addr:         "127.0.0.1:6379",
				Password:     "",
				Db:           0,
				DialTimeout:  &durationpb.Duration{Seconds: 1},
				ReadTimeout:  &durationpb.Duration{Nanos: 400 * 1000 * 1000},
				WriteTimeout: &durationpb.Duration{Nanos: 600 * 1000 * 1000},
			},
		},
	}
}
