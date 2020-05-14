package utils

import (
	"clients/model"
	"clients/pkg/errno"
	"github.com/lexkong/log"
	scpu "github.com/shirou/gopsutil/cpu"
	sdisk "github.com/shirou/gopsutil/disk"
	shost "github.com/shirou/gopsutil/host"
	sload "github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func SystemMonitor() *models.ServerPercent {
	// create server info
	var info models.ServerPercent

	// CPU
	cpuPercent, _ := scpu.Percent(time.Second, false)
	info.CPU = cpuPercent[0]

	// 内存
	memory, _ := mem.VirtualMemory()
	info.Mem = memory.UsedPercent

	// 交换分区
	swap, _ := mem.SwapMemory()
	info.Swap = swap.UsedPercent

	// 综合衡量
	load, _ := sload.Avg()
	host, _ := shost.Info()
	info.Load = load
	info.Uptime = host.Uptime

	// 硬盘
	allDisk, _ := sdisk.Partitions(false)
	pDisk := make([]*models.DiskPercent, 0)
	info.Disk = make([]*models.DiskPercent, len(allDisk))
	for _, dValue := range allDisk {
		disk, err := sdisk.Usage(dValue.Mountpoint)
		if err != nil {
			continue
		}
		pDisk = append(pDisk, &models.DiskPercent{
			Path: dValue.Mountpoint,
			Use:  disk.UsedPercent,
		})
	}
	//fmt.Println(pDisk)
	info.Disk = pDisk

	// 硬盘IO
	allDiskIO := make([]*models.DiskIO, 0)
	diskIOs, _ := sdisk.IOCounters()
	for iok, iov := range diskIOs {
		//fmt.Println(iov)
		allDiskIO = append(allDiskIO, &models.DiskIO{
			Device:     iok,
			ReadCount:  iov.ReadCount,
			WriteCount: iov.WriteCount,
			ReadBytes:  iov.ReadBytes,
			WriteBytes: iov.WriteBytes,
		})
	}
	info.DiskIO = allDiskIO

	// 网络
	// 特殊处理，需要延时1秒做差值
	oldNetwork, _ := net.IOCounters(false)
	time.Sleep(1 * time.Second)
	network, _ := net.IOCounters(false)
	//fmt.Println(network)
	np := models.NetworkPercent{
		ByteSent:    network[0].BytesSent - oldNetwork[0].BytesSent,
		ByteRecv:    network[0].BytesRecv - oldNetwork[0].BytesRecv,
		PacketsSent: network[0].PacketsSent - oldNetwork[0].PacketsSent,
		PacketsRecv: network[0].PacketsRecv - oldNetwork[0].PacketsRecv,
		Errin:       network[0].Errin,
		Errout:      network[0].Errout,
		Dropin:      network[0].Dropin,
		Fifoin:      network[0].Fifoin,
		Fifoout:     network[0].Fifoout,
	}
	//fmt.Println(np)
	info.Network = &np
	return &info
}

func SendMonitor() error {
	data := SystemMonitor()
	if data == nil {
		log.Errorf(errno.ErrScheduledTasks, `SystemMonitor`)
	}
	res := StructToMap(data)
	err := pushData(res, "/api/host/monitor/update")
	if err != nil {
		return err
	}
	return nil
}
