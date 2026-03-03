package open

import (
	"context"
	"testing"
)

func TestNewChatBot(t *testing.T) {
	// 测试创建企业微信机器人
	weworkConfig := BotConfig{
		Type:     BotTypeWework,
		RobotKey: "813e2278-7858-432f-8ecb-1ac0c1a41bac",
	}

	weworkBot, err := NewChatBot(weworkConfig)
	if err != nil {
		t.Log("创建企业微信机器人出错:", err)
	} else {
		t.Log("创建企业微信机器人成功")

		// 测试发送文本消息
		err = weworkBot.SendText(context.Background(), "测试文本消息", nil)
		if err != nil {
			t.Log("发送文本消息出错:", err)
		} else {
			t.Log("发送文本消息成功")
		}

		// 测试发送Markdown消息
		err = weworkBot.SendMarkdown(context.Background(), "# 测试Markdown消息", nil)
		if err != nil {
			t.Log("发送Markdown消息出错:", err)
		} else {
			t.Log("发送Markdown消息成功")
		}
	}

	// 测试创建钉钉机器人
	dingtalkConfig := BotConfig{
		Type:        BotTypeDingtalk,
		AccessToken: "your-access-token",
		Secret:      "your-secret",
	}

	dingtalkBot, err := NewChatBot(dingtalkConfig)
	if err != nil {
		t.Log("创建钉钉机器人出错:", err)
	} else {
		t.Log("创建钉钉机器人成功")

		// 测试发送文本消息
		err = dingtalkBot.SendText(context.Background(), "测试文本消息", nil)
		if err != nil {
			t.Log("发送文本消息出错:", err)
		} else {
			t.Log("发送文本消息成功")
		}

		// 测试发送Markdown消息
		err = dingtalkBot.SendMarkdown(context.Background(), "# 测试Markdown消息", map[string]interface{}{
			"title": "测试Markdown消息",
		})
		if err != nil {
			t.Log("发送Markdown消息出错:", err)
		} else {
			t.Log("发送Markdown消息成功")
		}
	}
}

func TestSendLink(t *testing.T) {
	// 测试发送链接消息
	weworkConfig := BotConfig{
		Type:     BotTypeWework,
		RobotKey: "813e2278-7858-432f-8ecb-1ac0c1a41bac",
	}

	bot, err := NewChatBot(weworkConfig)
	if err != nil {
		t.Log("创建机器人出错:", err)
		return
	}

	linkData := map[string]string{
		"title":      "测试链接消息",
		"text":       "这是一条测试链接消息",
		"messageUrl": "https://www.example.com",
		"picUrl":     "https://www.example.com/image.jpg",
	}

	err = bot.SendLink(context.Background(), linkData)
	if err != nil {
		t.Log("发送链接消息出错:", err)
	} else {
		t.Log("发送链接消息成功")
	}
}

func TestSendNews(t *testing.T) {
	// 测试发送图文消息
	weworkConfig := BotConfig{
		Type:     BotTypeWework,
		RobotKey: "813e2278-7858-432f-8ecb-1ac0c1a41bac",
	}

	bot, err := NewChatBot(weworkConfig)
	if err != nil {
		t.Log("创建机器人出错:", err)
		return
	}

	articles := []map[string]string{
		{
			"title":       "测试图文消息",
			"description": "这是一条测试图文消息",
			"url":         "https://www.example.com",
			"picUrl":      "https://www.example.com/image.jpg",
		},
	}

	err = bot.SendNews(context.Background(), articles)
	if err != nil {
		t.Log("发送图文消息出错:", err)
	} else {
		t.Log("发送图文消息成功")
	}
}
