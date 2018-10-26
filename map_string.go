package main

import (
	"encoding/json"
	"fmt"
)

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
func main() {
	fmt.Println("Hello, playground")
	mr, err := Fetch(MetricSet{17, 26, 34, 33})
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(mr, "", "  ")
	fmt.Println(string(b))

}
