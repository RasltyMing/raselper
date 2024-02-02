package raselper

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func RunCmd(cmd string) (error error, outStr string, errStr string) {
	var stdout, stderr bytes.Buffer
	var cmdExec *exec.Cmd

	fmt.Println("运行命令", cmd)
	switch runtime.GOOS {
	case "windows":
		//fmt.Println("当前操作系统是windows")
		cmdExec = exec.Command("cmd", "/c", cmd)
	case "darwin":
		//fmt.Println("当前操作系统是macOS")
	case "linux":
		//fmt.Println("当前操作系统是linux")
		cmdExec = exec.Command("sh", "-c", cmd)
	default:
		fmt.Println("当前操作系统未知")
	}
	cmdExec.Stdout = &stdout
	cmdExec.Stderr = &stderr
	err := cmdExec.Run()
	if err != nil {
		fmt.Println(err)
		return err, "", ""
	}
	if stdout.String() != "" {
		fmt.Println(stdout.String())
	}
	if stderr.String() != "" {
		fmt.Println(stderr.String())
	}
	return nil, stdout.String(), stderr.String()
}
