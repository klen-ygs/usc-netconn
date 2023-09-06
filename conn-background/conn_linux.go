package main

import (
	"context"
	"os/exec"
	"time"
)

func ConnUsc() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	cmd := exec.CommandContext(ctx, "microsoft-edge", "--new-window", "http://210.43.112.9/srun_portal_pc?ac_id=5&theme=basic")
	return cmd.Run()
}
