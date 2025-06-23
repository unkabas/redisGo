package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unkabas/redisGo/config"
	"io"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getWeather(c *gin.Context) {
	city := c.Param("city")

	value, err := config.Rdb.Get(config.Ctx, city).Result()
	if err != nil {
		fmt.Println("city wasnt found")
	} else {
		fmt.Println("city found")
		c.JSON(200, gin.H{
			"message": value,
		})
	}

	link := os.Getenv("weatherUrl") + city + "&appid=" + os.Getenv("apiKey") + "&units=metric"
	data, err := http.Get(link)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "fuck",
		})
		return
	}
	body, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	err = config.Rdb.Set(config.Ctx, city, weather.Main.Temp, 100*time.Second).Err()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "redis error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": weather.Main.Temp,
	})

}
func main() {
	config.RedisConnect()
	config.LoadEnv()

	r := gin.Default()

	r.GET("/weather/:city", getWeather)
	r.Run(":8080")
}
