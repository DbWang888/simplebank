package util

import (
	"time"

	"github.com/spf13/viper"
)

//保存所有的配置信息
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

//从指定路径的配置文件中读取配置信息
func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)  //配置文件路劲
	viper.SetConfigName("app") //配置文件名称
	viper.SetConfigType("env") //配置文件类型

	viper.AutomaticEnv()

	err = viper.ReadInConfig()  //自动覆盖环境变量的默认值
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
