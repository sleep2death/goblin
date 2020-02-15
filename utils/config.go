package utils

import (
	"time"

	"github.com/spf13/viper"
)

func InitConfig() error {
	// you can comment these default settings, if you had written a config file.
	viper.SetDefault("port", 9000)
	viper.SetDefault("readtimeout", time.Minute*5)  // 5 minutes
	viper.SetDefault("idletimeout", time.Minute*5)  // 5 minutes
	viper.SetDefault("writetimeout", time.Second*1) // 1 second
	viper.SetDefault("dbname", "goblin")
	viper.SetDefault("dbaddr", "mongodb://localhost:27017")
	viper.SetDefault("dbreadtimeout", time.Second*5)
	viper.SetDefault("dbwritetimeout", time.Second*5)
	viper.SetDefault("tokenexpiretime", time.Hour*24)

	viper.SetConfigName(".goblin")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	return viper.ReadInConfig() // Find and read the config file
}
