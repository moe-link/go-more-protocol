package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Config() {

	viper.SetDefault("server", map[string]string{"ip": "127.0.0.1", "port": "8080"})

}

func Test() {
	fmt.Println("")
}
