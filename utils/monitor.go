package utils

import (
	"clients/model"
	scpu "github.com/shirou/gopsutil/cpu"
	sdisk "github.com/shirou/gopsutil/disk"
	shost "github.com/shirou/gopsutil/host"
	sload "github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func SystemMonitor() *models.Server {
	// create server info
	info := new(models.Server)

	// CPU
	cpuPercent, _ := scpu.Percent(time.Second, false)
	info.Percent.CPU = cpuPercent[0]

	// 综合衡量
	load, _ := sload.Avg()
	host, _ := shost.Info()
	info.Load = load
	info.Uptime = host.Uptime

	// 内存
	memory, _ := mem.VirtualMemory()
	info.Mem.Available = memory.Available
	info.Mem.Used = memory.Used
	info.Percent.Mem = memory.UsedPercent

	// 交换分区
	swap, _ := mem.SwapMemory()
	info.Swap.Available = swap.Free
	info.Swap.Used = swap.Used
	info.Percent.Swap = swap.UsedPercent

	// 硬盘 TODO 案例
	allDisk, _ := sdisk.Partitions(true)
	//info.Disk = make([]models.DiskInfo, len(allDisk))
	//info.Percent.Disk = make([]models.DiskPercent, len(allDisk))
	//aDisk := make([]*models.DiskInfo, 0)
	pDisk := make([]*models.DiskPercent, 0)
	//fmt.Println(allDisk)
	for _, dValue := range allDisk {
		//fmt.Println(dValue)
		disk, err := sdisk.Usage(dValue.Device)
		if err != nil {
			//fmt.Println(err)
			continue
		}
		//aDisk = append(aDisk, &models.DiskInfo{
		//	User: disk.Used,
		//	Free: disk.Free,
		//})
		pDisk = append(pDisk, &models.DiskPercent{
			Path: dValue.Device,
			User: disk.UsedPercent,
		})
	}
	//info.Disk = aDisk
	info.Percent.Disk = pDisk

	//// 网络
	network, _ := net.IOCounters(true)
	networkInterfaces, _ := net.Interfaces()
	info.Network = make(map[string]models.InterfaceInfo)
	for _, networkV := range network {
		ii := models.InterfaceInfo{}
		ii.ByteSent = networkV.BytesSent
		ii.ByteRecv = networkV.BytesRecv
		info.Network[networkV.Name] = ii
	}
	for _, networkInterfacesV := range networkInterfaces {
		if nw, ok := info.Network[networkInterfacesV.Name]; ok {
			nw.Addrs = make([]string, len(networkInterfacesV.Addrs))
			for n, nnw := range networkInterfacesV.Addrs {
				nw.Addrs[n] = nnw.Addr
			}
			info.Network[networkInterfacesV.Name] = nw
		}
	}

	return info
}
