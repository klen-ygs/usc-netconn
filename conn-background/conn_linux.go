package main

import (
	"context"
	"os/exec"
	"time"
)

func ConnUsc() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	cmd := exec.CommandContext(ctx, "microsoft-edge", "--new-window", "http://210.43.112.9/srun_portal_pc?ac_id=5&theme=basic")
	return cmd.Run()
}

func CreateNotify(msg string) {
	command := exec.Command("notify-send", "-a", "USCConn", msg)
	command.Run()
}
