package initialize

import (
	"fmt"
	"gate/global"
	"github.com/spf13/viper"
	"log"
)

func Viper() {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config.docker")
	//添加读取的配置文件路径
	v.AddConfigPath("connect/")
	//设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Panicf("err:%s\n", err)
	}
	if err := v.Unmarshal(&global.FXConfig); err != nil {
		log.Panicf("配置文件获取失败", err)
	}
	x := global.FXConfig
	fmt.Println(x)
}
