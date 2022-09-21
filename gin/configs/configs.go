package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.gin/pkg/env"
	"time"
)

var config = new(Config)

func Get() Config {
	return *config
}

type Config struct {
	Server struct {
		RunMode      string        `toml:"runMode"`
		HttpPort     int           `toml:"httpPort"`
		ReadTimeout  time.Duration `toml:"readTimeout"`
		WriteTimeout time.Duration `toml:"writeTimeout"`
	}
}

func init() {
	fmt.Println(env.Active().Value())
	viper.SetConfigName(env.Active().Value() + "-configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(config)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})

}
