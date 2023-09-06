package main

import (
	"context"
	"github.com/go-toast/toast"
	"os/exec"
	"syscall"
	"time"
)

func init() {
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return
	}
	defer syscall.FreeLibrary(kernel32)

	freeConsole, err := syscall.GetProcAddress(kernel32, "FreeConsole")
	if err != nil {
		return
	}
	syscall.SyscallN(freeConsole)
}

func ConnUsc() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	cmd := exec.CommandContext(ctx, "cmd", "/c", "start", "http://210.43.112.9/srun_portal_pc?ac_id=5&theme=basic")
	return cmd.Run()
}

func CreateNotify(msg string) {
	notify := toast.Notification{
		AppID:   "USCConn",
		Message: msg,
	}
	notify.Push()
}
