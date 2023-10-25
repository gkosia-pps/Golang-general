package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv string `mapstructure:"APP_ENV"`
	DBUser string `mapstructure:"DB_USER"`
}

func main() {

	// os
	app := os.Getenv("App_name")
	fmt.Printf("Env var passed from cmd is %v \n", app)

	// godotenv
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Cannot load the env file")
	}
	value_from_env_file := os.Getenv("STRONGEST_AVENGER")
	fmt.Printf("Value readed from env file is %v \n", value_from_env_file)

	// Viber
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper_value := viper.Get("STRONGEST_AVENGER")
	fmt.Printf("Value readed using viber %v \n", viper_value)

	// Viber from json
	viper.SetConfigName("19_env_json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	fmt.Printf("Readed from viper %v \n", viper.GetString("APP_ENV"))

	// viper from struct
	var vp_conf Config
	err = viper.Unmarshal(&vp_conf)

	if err != nil {
		fmt.Print("Cannot Unmarshal config")
	}

	fmt.Printf("From viper Unmarshal readed %v", vp_conf.AppEnv)

}
