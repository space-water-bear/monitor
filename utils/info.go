package utils

import (
	"clients/model"
	scpu "github.com/shirou/gopsutil/cpu"
	sdisk "github.com/shirou/gopsutil/disk"
	shost "github.com/shirou/gopsutil/host"
	sload "github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"runtime"
	"time"
)

func SystemInfo() *models.Server {
	// create server info
	info := new(models.Server)

	// CPU
	cpu, _ := scpu.Info()
	cpuPercent, _ := scpu.Percent(time.Second, false)
	//info.CPU = make([]models.CPUInfo, len(cpu))
	info.Percent.CPU = cpuPercent[0]
	//fmt.Println(cpu)
	// 默认是linux
	info.CPU.Cores = int32(len(cpu))
	info.CPU.ModelName = cpu[0].ModelName
	// 开发者本地为OS X
	if runtime.GOOS == "darwin" {
		info.CPU.Cores = cpu[0].Cores
		info.CPU.ModelName = cpu[0].ModelName
	}

	// 综合衡量
	load, _ := sload.Avg()
	host, _ := shost.Info()
	info.Load = load
	info.Uptime = host.Uptime
	info.BootTime = host.BootTime

	// 内存
	memory, _ := mem.VirtualMemory()
	//fmt.Println(memory)
	info.Mem.Total = memory.Total
	info.Mem.Available = memory.Available
	info.Mem.Used = memory.Used
	info.Percent.Mem = memory.UsedPercent

	// 交换分区
	swap, _ := mem.SwapMemory()
	info.Swap.Total = swap.Total
	info.Swap.Available = swap.Free
	info.Swap.Used = swap.Used
	info.Percent.Swap = swap.UsedPercent

	// 硬盘 TODO 案例
	allDisk, _ := sdisk.Partitions(false)
	aDisk := make([]*models.DiskInfo, 0)
	pDisk := make([]*models.DiskPercent, 0)
	info.Disk = make([]*models.DiskInfo, len(allDisk))
	info.Percent.Disk = make([]*models.DiskPercent, len(allDisk))
	for _, dValue := range allDisk {
		//fmt.Println(dValue)
		disk, err := sdisk.Usage(dValue.Mountpoint)
		if err != nil {
			continue
		}
		//fmt.Println(disk)
		aDisk = append(aDisk, &models.DiskInfo{
			User:   disk.Used,
			Free:   disk.Free,
			Path:   disk.Path,
			FsType: disk.Fstype,
			Total:  disk.Total,
		})
		pDisk = append(pDisk, &models.DiskPercent{
			Path: dValue.Mountpoint,
			User: disk.UsedPercent,
		})
	}
	info.Disk = aDisk
	info.Percent.Disk = pDisk

	allDiskIO := make([]*models.DiskIO, 0)
	diskIOs, _ := sdisk.IOCounters()
	for iok, iov := range diskIOs {
		allDiskIO = append(allDiskIO, &models.DiskIO{
			Device:     iok,
			ReadCount:  iov.ReadCount,
			WriteCount: iov.WriteCount,
			ReadBytes:  iov.ReadBytes,
			WriteBytes: iov.WriteBytes,
		})
	}
	info.Percent.DiskIO = allDiskIO

	// 网络
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
