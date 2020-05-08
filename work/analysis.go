package work

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

//通过管道同步获取日志的函数
func syncLog(reader io.ReadCloser) {
	f, _ := os.OpenFile("file.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	buf := make([]byte, 1024, 1024)
	for {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			f.WriteString(string(outputByte))
		}
		if err != nil {
			//读到结尾
			if err == io.EOF || strings.Contains(err.Error(), "file already closed") {
				err = nil
			}
		}
	}
}

func Analysis() {
	cmdStr := `
		#!/bin/bash
		for var in {1..20}
		do
			sleep 1
			 echo "Hello, Welcome ${var} times "
		done`
	cmd := exec.Command("bash", "-c", cmdStr)

	cmdStdoutPipe, _ := cmd.StdoutPipe()
	cmdStderrPipe, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	go syncLog(cmdStdoutPipe)
	go syncLog(cmdStderrPipe)

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
