package main

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	var data = make([]byte, 1024)
	var chanel = make(chan struct{})
	go func() {
		for {
			chanel <- struct{}{}
			time.Sleep(time.Second * 10)
		}

	}()

	for {
		time.Sleep(time.Second * 3)
		resp, err := http.Get("http://bilibili.com")
		if err != nil {
			slog.Warn("网络连接断开", "err", err)
			continue
		}
		_, err = io.ReadFull(resp.Body, data)
		if err != nil {
			slog.Warn("网络连接断开", "err", err)
			continue
		}
		if bytes.Contains(data, []byte("深澜软件")) {
			<-chanel
			slog.Info("登录usc")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
			cmd := exec.CommandContext(ctx, "microsoft-edge", "--new-window", "http://210.43.112.9/srun_portal_pc?ac_id=5&theme=basic")
			err := cmd.Run()
			if err != nil {
				slog.Warn("打开浏览器失败", "err", err)
				cancel()
				continue
			}
		}

	}
}
