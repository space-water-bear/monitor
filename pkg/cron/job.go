package cron

import (
	"clients/utils"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type SystemInfo struct{}

type SystemMonitor struct{}

// 实现 cron interface
func (s *SystemInfo) Run() {
	fmt.Println(`SystemInfo`)
	//log.Info(`running SystemInfo !!!`)
}

// 实现 cron interface
func (s *SystemMonitor) Run() {
	err := utils.SendMonitor()
	if err != nil {
		log.Errorf(err, `计划任务失败: SystemMonitor`)
	}
}

func AddSystemInfoJob() error {
	spec := viper.GetString("monitor.host")
	//fmt.Println(spec)
	err := crontab.AddJob(spec, &SystemMonitor{})
	if err != nil {
		return err
	}
	//log.Info(`add system info job !!!`)
	return nil
}
