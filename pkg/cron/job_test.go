package cron

import (
	"github.com/robfig/cron"
	"testing"
	"time"
)

func TestSystemInfo(t *testing.T) {
	t.Log(`start SystemInfo`)
	crontab = cron.New()
	spec := "*/3 * * * * ?"
	err := crontab.AddJob(spec, &SystemInfo{})
	if err != nil {
		t.Fatal(`Failed, `, err)
	}

	crontab.Start()
	defer crontab.Stop()
	//select {}
	time.Sleep(time.Minute * 1)
	t.Log(`success`)
}

func TestSystemMonitor(t *testing.T) {
	t.Log(`start SystemMonitor`)
	crontab = cron.New()
	spec := "*/3 * * * * ?"

	crontab.Start()
	defer crontab.Stop()
	//select {}
	err := crontab.AddJob(spec, &SystemMonitor{})
	if err != nil {
		t.Fatal(`Failed, `, err)
	}

	time.Sleep(time.Minute * 1)
	t.Log(`success`)
}