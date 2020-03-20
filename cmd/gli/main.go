package main

import (
	"flag"
	"github.com/nikitsenka/gli/command"
	"github.com/spf13/viper"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yml", "Defines the path, name and extension of the config file")
	flag.Parse()
	if configFile != "" {
		viper.SetConfigFile(configFile)
		viper.ReadInConfig()
	}
	command.Execute()
}
