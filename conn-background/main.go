package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	var data = make([]byte, 1024)

	for i := 0; i < 70; i++ {
		time.Sleep(time.Second * 5)
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
			ConnUsc()
		}

	}
}
