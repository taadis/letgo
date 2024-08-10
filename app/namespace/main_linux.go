package main

import (
	"log/slog"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// fork出一个新进程,让新进程执行shell命令
	cmd := exec.Command("sh")
	// 设置系统调用参数，这里传入的 Cloneflags 要创建的对应资源的命名空间
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
		//Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNET,
	}
	// 重定标准输入和标准输出
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		slog.Error("cmd.Run() failed", "error", err)
	}
}
