package utils

import (
	"clients/config"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSystemMonitor(t *testing.T) {
	data := SystemMonitor()
	if data == nil {
		t.Error("SystemMonitor failed!")
	}
	//fmt.Println(data)
	//pushData(Struct2Map(data), "")
	formatJson, _ := json.Marshal(data)
	fmt.Println(string(formatJson))
	t.Log("SystemMonitor test pass")
}

func TestPushData(t *testing.T) {
	err := config.Init("/Users/cengcanguang/work/clients/conf/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	data := SystemMonitor()
	if data == nil {
		t.Error("SystemMonitor failed!")
	}
	res := StructToMap(data)
	pushData(res, "/api/host/monitor/update")
	t.Log("PushData test pass")
}
