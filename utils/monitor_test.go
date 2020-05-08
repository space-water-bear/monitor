package utils

import (
	"fmt"
	"testing"
)

func TestSystemMonitor(t *testing.T) {
	data := SystemMonitor()
	if data == nil {
		t.Error("SystemMonitor failed!")
	}
	fmt.Println(data)
	t.Log("SystemMonitor test pass")
}

func TestSystemInfo(t *testing.T) {
	data := SystemInfo()
	if data == nil {
		t.Error("SystemInfo failed!")
	}
	fmt.Println(data)
	t.Log("SystemInfo test pass")
}