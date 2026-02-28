package wework

import (
	"context"
	"io"
	"net/http"
	"testing"
)

func TestNewChatGroupRobot(t *testing.T) {
	sendMessage, err := NewChatGroupRobot("813e2278-7858-432f-8ecb-1ac0c1a41bac")
	if err != nil {
		t.Log("生产机器人出错")
		return
	}
	req, err := sendMessage(context.Background(), "{\"msgtype\": \"text\",\"text\": {\"content\": \"hello world\"}}")
	if err != nil {
		t.Log("构造请求报文出从")
		return
	}
	t.Log("请求地址：", req.URL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log("发送请求出错")
		return
	}
	data, _ := io.ReadAll(resp.Body)
	t.Log("请求已发送：", string(data))
}
