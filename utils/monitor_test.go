package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSystemMonitor(t *testing.T) {
	data := SystemMonitor()
	if data == nil {
		t.Error("SystemMonitor failed!")
	}
	formatJson, _ := json.Marshal(data)
	fmt.Println(string(formatJson))
	t.Log("SystemMonitor test pass")
}

func TestSystemInfo(t *testing.T) {
	data := SystemInfo()
	if data == nil {
		t.Error("SystemInfo failed!")
	}

	formatJson, _ := json.Marshal(data)
	fmt.Println(string(formatJson))
	t.Log("SystemInfo test pass")
}
