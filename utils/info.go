package utils

import (
	"clients/model"
	"clients/pkg/errno"
	"github.com/lexkong/log"
	scpu "github.com/shirou/gopsutil/cpu"
	sdisk "github.com/shirou/gopsutil/disk"
	shost "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"runtime"
)

func SystemInfo() *models.Server {
	// create server info
	info := new(models.Server)

	// CPU
	cpu, _ := scpu.Info()
	// 默认是linux
	info.CPU.Cores = int32(len(cpu))
	info.CPU.ModelName = cpu[0].ModelName
	// 开发者本地为OS X
	if runtime.GOOS == "darwin" {
		info.CPU.Cores = cpu[0].Cores
		info.CPU.ModelName = cpu[0].ModelName
	}

	// 综合衡量
	host, _ := shost.Info()
	info.Kernel = host.KernelVersion
	info.OS = host.OS

	// 内存
	memory, _ := mem.VirtualMemory()
	info.Mem.Total = memory.Total
	info.Mem.Available = memory.Available
	info.Mem.Used = memory.Used

	// 交换分区
	swap, _ := mem.SwapMemory()
	info.Swap.Total = swap.Total
	info.Swap.Available = swap.Free
	info.Swap.Used = swap.Used

	// 硬盘 TODO 案例
	allDisk, _ := sdisk.Partitions(false)
	aDisk := make([]*models.DiskInfo, 0)
	info.Disk = make([]*models.DiskInfo, len(allDisk))
	for _, dValue := range allDisk {
		disk, err := sdisk.Usage(dValue.Mountpoint)
		if err != nil {
			continue
		}
		aDisk = append(aDisk, &models.DiskInfo{
			Use:    disk.Used,
			Free:   disk.Free,
			Path:   disk.Path,
			FsType: disk.Fstype,
			Total:  disk.Total,
		})
	}
	info.Disk = aDisk

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

func SendInfo() error {
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
