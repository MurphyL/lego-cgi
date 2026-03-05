package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"murphyl.com/lego/fns/perform"
)

// 企业微信群聊机器人 - https://developer.work.weixin.qq.com/document/path/99110

const (
	chatBotWebook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%v"
)

type messageType string

type messagePack map[string]any

// 发送消息的可选配置项
type messageOption func(messagePack)

const (
	MessageTypeText     messageType = "text"        // 文本消息
	MessageTypeMarkdown messageType = "markdown_v2" // Markdown消息
	MessageTypeImage    messageType = "image"       // 图片消息
	MessageTypeVoice    messageType = "voice"       // 语音消息
	MessageTypeFile     messageType = "file"        // 文件消息
	MessageTypeNews     messageType = "news"        // 图文消息
)

type ChatBot struct {
	makeRequest perform.PerformAgent[*http.Request, []byte]
	httpClient  *http.Client
}

// 新建企业微信群聊机器人
func NewChatBot(robotKey string) ChatBot {
	robotEndpoint := fmt.Sprintf(chatBotWebook, robotKey)
	return ChatBot{
		httpClient: http.DefaultClient,
		makeRequest: func(b []byte) (*http.Request, error) {
			httpReq, err := http.NewRequest(http.MethodPost, robotEndpoint, bytes.NewReader(b))
			if err != nil {
				return nil, fmt.Errorf("创建HTTP请求出错：%v", err.Error())
			}
			httpReq.Header.Set("Content-Type", "application/json")
			return httpReq, nil
		},
	}
}

// 发送文本消息
func (bot *ChatBot) SendTextMessage(content string, opts ...messageOption) error {
	body := newMessage(MessageTypeText, content)
	for _, opt := range opts {
		opt(body)
	}
	return bot.sendMessage(body)
}

// 发送Markdown消息
func (bot *ChatBot) SendMarkdownMessage(content string) error {
	body := newMessage(MessageTypeMarkdown, content)
	return bot.sendMessage(body)
}

// 发送消息
func (bot *ChatBot) sendMessage(body map[string]any) error {
	messageBytes, _ := json.Marshal(body)
	// 构造请求
	httpReq, err := bot.makeRequest(messageBytes)
	if err != nil {
		return fmt.Errorf("构造请求出错：%v", err.Error())
	}
	// 发送请求
	httpResp, err := bot.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("发送请求出错：%v", err.Error())
	}
	defer httpResp.Body.Close()
	// 验证请求响应
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("发送请求出错：%v", httpResp.Status)
	}
	if respBody, err := io.ReadAll(httpResp.Body); err == nil {
		ret := WeworkResp{}
		if err = json.Unmarshal(respBody, &ret); err == nil && ret.ErrCode == 0 {
			return nil
		}
		return fmt.Errorf("读取响应出错：%v", err.Error())
	} else {
		return fmt.Errorf("执行HTTP请求出错：%v", err.Error())
	}
}

// 构造请求
func newMessage(msgtype messageType, content string) messagePack {
	message := map[string]any{"msgtype": msgtype}
	switch msgtype {
	case MessageTypeText: // 文本消息，支持@用户
		message[string(msgtype)] = map[string]any{"content": content}
	case MessageTypeMarkdown: // Markdown消息
		message[string(msgtype)] = map[string]any{"content": content}
	case MessageTypeFile, MessageTypeVoice: // 文件消息、语音消息
		message[string(msgtype)] = map[string]any{"media_id": content}
	default:
		panic(fmt.Sprintf("不支持的消息类型：%v", msgtype))
	}
	return message
}
