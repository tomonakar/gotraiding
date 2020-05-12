package main

import (
	"app/app/controllers"
	"app/app/models"
	"app/config"
	"app/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
}

// tickerの確認
// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

// tickerChannel := make(chan bitflyer.Ticker)
// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
// for ticker := range tickerChannel {
// 	fmt.Println(ticker)
// 	fmt.Println(ticker.GetMidPrice())
// 	fmt.Println(ticker.DateTime())
// 	fmt.Println(ticker.TruncateDateTime(time.Second))
// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
// }
