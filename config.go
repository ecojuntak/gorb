package main

import "github.com/spf13/viper"

func getAddress() string {
	return viper.GetString("HOST") + ":" + viper.GetString("PORT")
}
