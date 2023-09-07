package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func TestInUSCWiFiButNotLogin() (bool, error) {
	data := make([]byte, 1000)
	resp, err := http.Get("http://bilibili.com")
	if err != nil {
		return false, fmt.Errorf("发起网络请求失败: %w", err)
	}
	_, err = io.ReadFull(resp.Body, data)
	if err != nil {
		return false, fmt.Errorf("发起网络请求失败: %w", err)
	}
	if bytes.Contains(data, []byte("深澜软件")) {
		return true, nil
	}
	return false, nil
}

func main() {

	for i := 0; i < 70; i++ {
		time.Sleep(time.Second * 3)

		inUSCButLogout, err := TestInUSCWiFiButNotLogin()
		if err != nil {
			slog.Warn("网络请求失败", "err", err)
			continue
		}
		if !inUSCButLogout {
			continue
		}

		ConnUsc()
		time.Sleep(time.Second * 3)
		inUSCButLogout, err = TestInUSCWiFiButNotLogin()
		if err != nil {
			if os.IsTimeout(err) {
				CreateNotify("网络连接已断开")
			}
			slog.Warn("网络请求失败", "err", err)
			time.Sleep(time.Second * 10)
			continue
		}
		if inUSCButLogout {
			slog.Warn("登录USC失败")
			CreateNotify("登录USC失败，稍后将重试")
			time.Sleep(time.Second * 10)
			continue
		}
		return
	}
}
