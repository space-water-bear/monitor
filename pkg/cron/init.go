package cron

import (
	"github.com/lexkong/log"
	"github.com/robfig/cron"
)

// 全局
var crontab *cron.Cron

func Init() {
	log.Info(`init crontab`)
	crontab = cron.New()
	crontab.Start()

	// default load system monitor function

	// default load game monitor function

	select {}
}

func Close()  {
	crontab.Stop()
}