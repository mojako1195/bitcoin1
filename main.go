package main

import (
	"bitcoin1/bitflyer"
	"bitcoin1/config"
	"bitcoin1/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
