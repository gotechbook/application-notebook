package viper

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gotechbook/application-notebook/config"
	vip "github.com/spf13/viper"
)

const (
	ConfigEnv  = "application-notebook"
	ConfigFile = "config.yaml"
)

func Viper(path ...string) *vip.Viper {
	var (
		v         = vip.New()
		conf      string
		configEnv string
	)
	// 优先级: 命令行 > 环境变量 > 默认值

	if len(path) != 0 {
		conf = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", conf)
	} else {
		flag.StringVar(&conf, "c", "", "choose config file.")
		flag.Parse()
		if conf == "" {
			configEnv = os.Getenv(ConfigEnv)
			if configEnv == "" {
				conf = ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
			} else {
				conf = configEnv
				fmt.Printf("您正在使用application-notebook环境变量,config的路径为%v\n", conf)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", conf)
		}
	}
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config.C); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&config.C); err != nil {
		fmt.Println(err)
	}
	return v
}
