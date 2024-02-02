package raselper

import (
	"bytes"
	"os/exec"
	"runtime"
)

func RunCmd(cmd string) (error error, outStr string, errStr string) {
	var stdout, stderr bytes.Buffer
	var cmdExec *exec.Cmd

	_ = Log("运行命令" + cmd)
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
		_ = Log("当前操作系统未知")
	}
	cmdExec.Stdout = &stdout
	cmdExec.Stderr = &stderr
	err := cmdExec.Run()
	if err != nil {
		_ = Log(err.Error())
		return err, "", ""
	}
	if stdout.String() != "" {
		_ = Log(stdout.String())
	}
	if stderr.String() != "" {
		_ = Log(stderr.String())
	}
	return nil, stdout.String(), stderr.String()
}
