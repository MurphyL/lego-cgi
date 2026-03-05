package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"murphyl.com/lego/fns/perform"
)

// 企业微信群聊机器人 - https://developer.work.weixin.qq.com/document/path/99110

const (
	WecomChatBotWebook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%v"
	DefaultTimeout     = 10 * time.Second
	DefaultMaxRetries  = 3
)

type MessageType string

type MessagePack map[string]any

// 文本消息配置项
type TextMessageOption func(MessagePack)

const (
	MessageTypeText     MessageType = "text"        // 文本消息
	MessageTypeMarkdown MessageType = "markdown_v2" // Markdown消息
	MessageTypeImage    MessageType = "image"       // 图片消息
	MessageTypeVoice    MessageType = "voice"       // 语音消息
	MessageTypeFile     MessageType = "file"        // 文件消息
	MessageTypeNews     MessageType = "news"        // 图文消息
)

type WecomChatBot struct {
	makeRequest perform.PerformAgent[*http.Request, []byte]
	httpClient  *http.Client
}

func NewWecomChatBot(robotKey string) WecomChatBot {
	robotEndpoint := fmt.Sprintf(WecomChatBotWebook, robotKey)
	return WecomChatBot{
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

func (bot *WecomChatBot) SendTextMessage(content string, opts ...TextMessageOption) error {
	body := newMessage(MessageTypeText, content)
	for _, opt := range opts {
		opt(body)
	}
	return bot.sendMessage(body)
}

func AtAll() TextMessageOption {
	return func(msg MessagePack) {
		msg["mentioned_list"] = []string{"@all"}
	}
}

func AtUserIds(userIds ...string) TextMessageOption {
	return func(msg MessagePack) {
		msg["mentioned_list"] = userIds
	}
}

func AtMobiles(userIds ...string) TextMessageOption {
	return func(msg MessagePack) {
		msg["mentioned_mobile_list"] = userIds
	}
}

func (bot *WecomChatBot) SendMarkdownMessage(content string) error {
	body := newMessage(MessageTypeMarkdown, content)
	return bot.sendMessage(body)
}

func (bot *WecomChatBot) sendMessage(body map[string]any) error {
	messageBytes, _ := json.Marshal(body)
	httpReq, err := bot.makeRequest(messageBytes)
	if err != nil {
		return fmt.Errorf("构造请求出错：%v", err.Error())
	}
	httpResp, err := bot.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("发送请求出错：%v", err.Error())
	}
	defer httpResp.Body.Close()
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

func newMessage(msgtype MessageType, content string) MessagePack {
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
