package dingtalk

import (
	"context"
	"testing"
)

func TestNewChatBot(t *testing.T) {
	// 请替换为实际的accessToken和secret
	accessToken := "your-access-token"
	secret := "your-secret"

	robot, err := NewChatBot(accessToken, secret)
	if err != nil {
		t.Log("创建机器人出错:", err)
		return
	}

	// 测试发送文本消息
	err = robot.SendText(context.Background(), "测试文本消息", nil)
	if err != nil {
		t.Log("发送文本消息出错:", err)
	} else {
		t.Log("发送文本消息成功")
	}

	// 测试发送Markdown消息
	markdownContent := "# 测试Markdown消息\n## 二级标题\n- 列表项1\n- 列表项2"
	err = robot.SendMarkdown(context.Background(), markdownContent, map[string]interface{}{
		"title": "测试Markdown消息",
	})
	if err != nil {
		t.Log("发送Markdown消息出错:", err)
	} else {
		t.Log("发送Markdown消息成功")
	}

	// 测试发送链接消息
	linkData := map[string]string{
		"title":      "测试链接消息",
		"text":       "这是一条测试链接消息",
		"messageUrl": "https://www.example.com",
		"picUrl":     "https://www.example.com/image.jpg",
	}
	err = robot.SendLink(context.Background(), linkData)
	if err != nil {
		t.Log("发送链接消息出错:", err)
	} else {
		t.Log("发送链接消息成功")
	}

	// 测试发送行动卡片消息
	cardData := map[string]interface{}{
		"title":          "测试行动卡片消息",
		"text":           "这是一条测试行动卡片消息",
		"btnOrientation": "1",
		"btns": []map[string]string{
			{"title": "按钮1", "actionURL": "https://www.example.com/1"},
			{"title": "按钮2", "actionURL": "https://www.example.com/2"},
		},
	}
	err = robot.SendActionCard(context.Background(), cardData)
	if err != nil {
		t.Log("发送行动卡片消息出错:", err)
	} else {
		t.Log("发送行动卡片消息成功")
	}

	// 测试发送Feed卡片消息
	links := []map[string]string{
		{"title": "测试Feed卡片1", "messageUrl": "https://www.example.com/1", "picUrl": "https://www.example.com/image1.jpg"},
		{"title": "测试Feed卡片2", "messageUrl": "https://www.example.com/2", "picUrl": "https://www.example.com/image2.jpg"},
	}
	err = robot.SendFeedCard(context.Background(), links)
	if err != nil {
		t.Log("发送Feed卡片消息出错:", err)
	} else {
		t.Log("发送Feed卡片消息成功")
	}
}

func TestCreateRequest(t *testing.T) {
	// 请替换为实际的accessToken和secret
	accessToken := "your-access-token"
	secret := "your-secret"

	robot, err := NewChatBot(accessToken, secret)
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
