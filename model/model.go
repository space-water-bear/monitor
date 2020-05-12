package models

import "github.com/shirou/gopsutil/load"

type Server struct {
	//Percent  StatusPercent            `json:"percent"`
	CPU     CPUInfo                  `json:"cpu"`
	Mem     MemInfo                  `json:"mem"`
	Swap    SwapInfo                 `json:"swap"`
	Disk    []*DiskInfo              `json:"disk"`
	Network map[string]InterfaceInfo `json:"network"`
	Kernel  string                   `json:"kernel"`
	OS      string                   `json:"os"`
}

type ServerPercent struct {
	CPU      float64         `json:"cpu"`
	Disk     []*DiskPercent  `json:"disk"`
	DiskIO   []*DiskIO       `json:"disk_io"`
	Mem      float64         `json:"mem"`
	Swap     float64         `json:"swap"`
	Network  *NetworkPercent `json:"network"`
	Load     *load.AvgStat   `json:"load"`
	BootTime uint64          `json:"boot_time"`
	Uptime   uint64          `json:"uptime"`
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
	Use    uint64 `json:"use"`
}

type DiskPercent struct {
	Path string  `json:"path"`
	Use  float64 `json:"use"`
}

type DiskIO struct {
	Device     string `json:"device"`
	ReadCount  uint64 `json:"read_count"`
	WriteCount uint64 `json:"write_count"`
	ReadBytes  uint64 `json:"read_bytes"`
	WriteBytes uint64 `json:"write_bytes"`
}

type NetworkPercent struct {
	ByteSent    uint64 `json:"byte_sent"`
	ByteRecv    uint64 `json:"byte_recv"`
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
	Errin       uint64 `json:"errin"`       // total number of errors while receiving
	Errout      uint64 `json:"errout"`      // total number of errors while sending
	Dropin      uint64 `json:"dropin"`      // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`     // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
	Fifoin      uint64 `json:"fifoin"`      // total number of FIFO buffers errors while receiving
	Fifoout     uint64 `json:"fifoout"`     // total number of FIFO buffers errors while sending
}

type InterfaceInfo struct {
	Addrs    []string `json:"addrs"`
	ByteSent uint64   `json:"byte_sent"`
	ByteRecv uint64   `json:"byte_recv"`
}
