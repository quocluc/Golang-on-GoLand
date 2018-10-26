package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type gps struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
	Time int     `json:"time"`
}
type res []gps

func Get(c *gin.Context) {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/loglag_gps")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT latitude, longitude, created_at FROM gps_event_logs WHERE order_vehicle_id = 29 AND driver_id = 395")
	if err != nil {
		panic(err.Error())
	}
	var ress res
	for result.Next() {
		var latitude, longitude float64
		var created_at int
		err = result.Scan(&latitude, &longitude, &created_at)
		if err != nil {
			panic(err.Error())
		}
		re := gps{latitude, longitude, created_at}
		ress = append(ress, re)
	}
	//fmt.Print(ress)

	dataJson, err := json.Marshal(ress)
	c.Data(200, "application/json", dataJson)
}
func JSONMarshal(v interface{}, backslashEscape bool) (string, error) {
	b, err := json.Marshal(v)
	var b_string string
	if backslashEscape {
		b_string = strings.Replace(string(b), "/", "", -1)
	}
	return b_string, err
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", Def)
	r.POST("/", Get)
	r.GET("/map", Map)
	r.GET("/test", func(c *gin.Context) {
		//c.Data("200","JSON",)
	})
	r.Run()
}
func Map(c *gin.Context) {
	mr, err := Fetch(MetricSet{17, 26, 34, 33})
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(mr, "", "  ")
	c.JSON(200, string(b))

}
func Def(c *gin.Context) {
	listGps := GetListGps()
	data, err := JSONMarshal(listGps, true)
	if (err != nil) {
		panic(err.Error())
		return
	}
	c.JSON(200, gin.H{
		"data": data,
	})
}
func GetListGps() gps {
	return gps{333, 333, 333,}
}

type MapStr map[string]interface{}
type MetricSet struct {
	cpu_status     uint64
	memory_status  uint64
	lte_status     uint64
	network_status uint64
}

func Fetch(m MetricSet) (MapStr, error) {
	k := [...]MapStr{
		{"name": "cpu", "status": m.cpu_status},
		{"name": "LTE", "status": m.lte_status},
		{"name": "Network", "status": m.network_status},
		{"name": "Memory", "status": m.memory_status},
	}
	event := MapStr{
		"cpu_status":     (m.cpu_status % 4),
		"memory_status":  (m.memory_status % 4),
		"lte_status":     (m.lte_status % 4),
		"network_status": (m.network_status % 4),
		"summary":        k,
	}

	m.cpu_status++
	m.memory_status = m.memory_status + 2
	m.lte_status = m.lte_status + 7
	m.network_status = m.network_status + 13

	return event, nil
}
