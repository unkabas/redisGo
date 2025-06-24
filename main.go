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
		c.JSON(500, gin.H{
			"message": "something wrong",
		})
	} else {
		fmt.Println("city found")
		c.JSON(200, gin.H{
			city: value,
		})
	}

	link := os.Getenv("weatherUrl") + city + "&appid=" + os.Getenv("apiKey") + "&units=metric"
	data, err := http.Get(link)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "fuck",
		})
		return
	}
	body, err := io.ReadAll(data.Body)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
		fmt.Println("Error reading response body:", err)
		return
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	err = config.Rdb.Set(config.Ctx, city, weather.Main.Temp, 100*time.Second).Err()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "redis error",
		})
		return
	}

}
func main() {
	config.RedisConnect()
	config.LoadEnv()

	r := gin.Default()

	r.GET("/weather/:city", getWeather)
	r.Run(":8080")
}
