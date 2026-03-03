package wework

import (
	"context"
	"testing"
)

func TestNewChatGroupRobot(t *testing.T) {
	robot, err := NewChatGroupRobot("813e2278-7858-432f-8ecb-1ac0c1a41bac")
	if err != nil {
		t.Log("创建机器人出错:", err)
		return
	}

	// 测试发送文本消息
	err = robot.SendText(context.Background(), "测试文本消息", nil, nil)
	if err != nil {
		t.Log("发送文本消息出错:", err)
	} else {
		t.Log("发送文本消息成功")
	}

	// 测试发送Markdown消息
	markdownContent := "# 测试Markdown消息\n## 二级标题\n- 列表项1\n- 列表项2"
	err = robot.SendMarkdown(context.Background(), markdownContent)
	if err != nil {
		t.Log("发送Markdown消息出错:", err)
	} else {
		t.Log("发送Markdown消息成功")
	}

	// 测试发送图文消息
	articles := []NewsArticle{
		{
			Title:       "测试图文消息",
			Description: "这是一条测试图文消息",
			URL:         "https://www.example.com",
			PicURL:      "https://www.example.com/image.jpg",
		},
	}
	err = robot.SendNews(context.Background(), articles)
	if err != nil {
		t.Log("发送图文消息出错:", err)
	} else {
		t.Log("发送图文消息成功")
	}
}

func TestCreateRequest(t *testing.T) {
	robot, err := NewChatGroupRobot("813e2278-7858-432f-8ecb-1ac0c1a41bac")
	if err != nil {
		t.Log("创建机器人出错:", err)
		return
	}

	// 测试创建文本消息请求
	message := Message{
		MsgType: MessageTypeText,
		Text: &TextMessage{
			Content: "测试文本消息",
		},
	}

	req, err := robot.CreateRequest(context.Background(), message)
	if err != nil {
		t.Log("创建请求出错:", err)
		return
	}

	t.Log("请求地址:", req.URL)
	t.Log("请求方法:", req.Method)
	t.Log("Content-Type:", req.Header.Get("Content-Type"))
}
