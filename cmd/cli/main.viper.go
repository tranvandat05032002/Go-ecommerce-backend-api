package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	vp := viper.New()

	vp.AddConfigPath("./config/")
	vp.SetConfigName("local")
	vp.SetConfigType("yaml")

	//read configuration
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server Port::", vp.GetInt("server.port"))

	var config Config
	if err = vp.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Config Port::", config.Server.Port)

	// get config is array
	for _, db := range config.Databases {
		fmt.Printf("\ndatabase ==>  user: %s, password: %s, host: %s", db.User, db.Password, db.Host)
	}
}
