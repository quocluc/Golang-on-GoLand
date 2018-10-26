package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type GPS struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
	Time int     `json:"time"`
}

func main2() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Me": "This Is Seasasadsad",
		})
	})

	r.GET("/data", GetGPS)

	r.Run()
}
func GetGPS(c *gin.Context) {
	gps := GPS{1, 1, 1}
	b, _ := json.Marshal(gps)

	//c.JSON(200, gin.H{
	//	"data": b,
	//})
	c.Data(200, "GPS", b)
}
