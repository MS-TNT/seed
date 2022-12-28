package cfg

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("$SOAPA_APP_CONFIG_PATH")
	viper.AddConfigPath("D:\\code\\seed\\backend\\cmd")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
