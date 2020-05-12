#!/usr/bin/env bash

path="/data/sofeware/ops/monitor"
PIDFILE="$path/monitor.pid"

pid=0
getPid()
{
  pid=`ps -ef|grep -v grep|grep $path|awk '{print $2}'`
}

start()
{
  getPid()
  if [ $pid != 0 ]; then
      echo "Process already exists !!! PID: $pid"
  fi
  #nohup $path/clients >/dev/null 2>&1 &
}

stop()
{

}

case "$1" in
start)	
	echo "Starting "
	;;
stop)	
	echo "Stoped"
  ;;
restart) 
	echo "Restartin"
  bash $0 stop
  bash $0 start
  ;;
status) 
  echo "Status..."
  ;;
*)	
  echo "Usage: /bin/bash service {start|stop|status|restart}"
  exit 2
  ;;
esac
exit 0