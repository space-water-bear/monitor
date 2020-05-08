package cron

import (
	"fmt"
	"github.com/lexkong/log"
)

type SystemInfo struct {}

type SystemMonitor struct {}

// 实现 cron interface
func (s *SystemInfo) Run() {
	fmt.Println(`SystemInfo`)
	//log.Info(`running SystemInfo !!!`)
}

// 实现 cron interface
func (s *SystemMonitor) Run() {
	fmt.Println(`SystemMonitor`)
}

func AddSystemInfoJob() error {
	spec := "*/3 * * * * ?"
	err := crontab.AddJob(spec, &SystemMonitor{})
	if err != nil {
		return err
	}
	log.Info(`add system info job !!!`)
	return nil
}
