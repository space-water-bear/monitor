package models

import (
	"github.com/shirou/gopsutil/load"
)

type Server struct {
	Percent  StatusPercent            `json:"percent"`
	CPU      CPUInfo                  `json:"cpu"`
	Mem      MemInfo                  `json:"mem"`
	Swap     SwapInfo                 `json:"swap"`
	Disk     []*DiskInfo              `json:"disk"`
	Load     *load.AvgStat            `json:"load"`
	Network  map[string]InterfaceInfo `json:"network"`
	BootTime uint64                   `json:"boot_time"`
	Uptime   uint64                   `json:"uptime"`
}

type StatusPercent struct {
	CPU    float64        `json:"cpu"`
	Disk   []*DiskPercent `json:"disk"`
	DiskIO []*DiskIO      `json:"disk_io"`
	Mem    float64        `json:"mem"`
	Swap   float64        `json:"swap"`
}

type CPUInfo struct {
	ModelName string `json:"model_name"`
	Cores     int32  `json:"cores"`
}

type MemInfo struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

type SwapInfo struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

type DiskInfo struct {
	Path   string `json:"path"`
	FsType string `json:"fs_type"`
	Total  uint64 `json:"total"`
	Free   uint64 `json:"free"`
	User   uint64 `json:"user"`
}

type DiskPercent struct {
	Path string  `json:"path"`
	User float64 `json:"user"`
}

type DiskIO struct {
	Device     string `json:"device"`
	ReadCount  uint64 `json:"read_count"`
	WriteCount uint64 `json:"write_count"`
	ReadBytes  uint64 `json:"read_bytes"`
	WriteBytes uint64 `json:"write_bytes"`
}

type InterfaceInfo struct {
	Addrs    []string `json:"addrs"`
	ByteSent uint64   `json:"byte_sent"`
	ByteRecv uint64   `json:"byte_recv"`
}
