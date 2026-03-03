package wework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=813e2278-7858-432f-8ecb-1ac0c1a41bac

const (
	WeworkChatGroupRobotWebook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	DefaultTimeout             = 10 * time.Second
	DefaultMaxRetries          = 3
)

// MessageType 消息类型
type MessageType string

// 消息类型常量
const (
	MessageTypeText     MessageType = "text"          // 文本消息
	MessageTypeMarkdown MessageType = "markdown"      // Markdown消息
	MessageTypeImage    MessageType = "image"         // 图片消息
	MessageTypeNews     MessageType = "news"          // 图文消息
	MessageTypeFile     MessageType = "file"          // 文件消息
	MessageTypeTemplate MessageType = "template_card" // 模板卡片消息
)

// TextMessage 文本消息
type TextMessage struct {
	Content             string   `json:"content"`                         // 消息内容
	MentionedList       []string `json:"mentioned_list,omitempty"`        // 提及的用户列表
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"` // 提及的手机号列表
}

// MarkdownMessage Markdown消息
type MarkdownMessage struct {
	Content string `json:"content"` // Markdown内容
}

// ImageMessage 图片消息
type ImageMessage struct {
	Base64 string `json:"base64"` // 图片base64编码
	MD5    string `json:"md5"`    // 图片md5值
}

// NewsArticle 图文消息文章
type NewsArticle struct {
	Title       string `json:"title"`       // 标题
	Description string `json:"description"` // 描述
	URL         string `json:"url"`         // 链接
	PicURL      string `json:"picurl"`      // 图片链接
}

// NewsMessage 图文消息
type NewsMessage struct {
	Articles []NewsArticle `json:"articles"` // 文章列表
}

// FileMessage 文件消息
type FileMessage struct {
	MediaID string `json:"media_id"` // 文件媒体ID
}

// Message 消息结构
type Message struct {
	MsgType  MessageType      `json:"msgtype"`            // 消息类型
	Text     *TextMessage     `json:"text,omitempty"`     // 文本消息
	Markdown *MarkdownMessage `json:"markdown,omitempty"` // Markdown消息
	Image    *ImageMessage    `json:"image,omitempty"`    // 图片消息
	News     *NewsMessage     `json:"news,omitempty"`     // 图文消息
	File     *FileMessage     `json:"file,omitempty"`     // 文件消息
}

// ChatGroupRobot 群机器人接口
type ChatGroupRobot interface {
	// CreateRequest 创建HTTP请求
	CreateRequest(ctx context.Context, data any) (*http.Request, error)
	// SendMessage 发送消息
	SendMessage(ctx context.Context, message Message) error
	// SendText 发送文本消息
	SendText(ctx context.Context, content string, mentionedList []string, mentionedMobileList []string) error
	// SendMarkdown 发送Markdown消息
	SendMarkdown(ctx context.Context, content string) error
	// SendImage 发送图片消息
	SendImage(ctx context.Context, base64 string, md5 string) error
	// SendNews 发送图文消息
	SendNews(ctx context.Context, articles []NewsArticle) error
	// SendFile 发送文件消息
	SendFile(ctx context.Context, mediaID string) error
}

// chatGroupRobot 群机器人实现
type chatGroupRobot struct {
	endpoint   string
	httpClient *http.Client
	maxRetries int
}

// NewChatGroupRobot 创建群机器人
func NewChatGroupRobot(robotKey string) (ChatGroupRobot, error) {
	robotWebhook, err := url.Parse(WeworkChatGroupRobotWebook)
	if err != nil {
		return nil, fmt.Errorf("解析URL出错：%v", err.Error())
	}
	robotWebhook.RawQuery = url.Values{"key": []string{robotKey}}.Encode()
	robotEndpoint := robotWebhook.String()

	return &chatGroupRobot{
		endpoint: robotEndpoint,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		maxRetries: DefaultMaxRetries,
	}, nil
}

// CreateRequest 创建HTTP请求
func (r *chatGroupRobot) CreateRequest(ctx context.Context, data any) (*http.Request, error) {
	var body []byte
	switch v := data.(type) {
	case []byte:
		body = v
	case string:
		body = []byte(v)
	default:
		tmp, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		body = tmp
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("构造报文出错：%v", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// SendMessage 发送消息
func (r *chatGroupRobot) SendMessage(ctx context.Context, message Message) error {
	var err error
	for i := 0; i < r.maxRetries; i++ {
		req, err := r.CreateRequest(ctx, message)
		if err != nil {
			return err
		}

		resp, err := r.httpClient.Do(req)
		if err != nil {
			// 网络错误，重试
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			// 检查响应内容
			var result struct {
				Errcode int    `json:"errcode"`
				Errmsg  string `json:"errmsg"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				return fmt.Errorf("解析响应出错：%v", err.Error())
			}
			if result.Errcode != 0 {
				return fmt.Errorf("发送消息失败：%s", result.Errmsg)
			}
			return nil
		}

		// 非200状态码，重试
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	return fmt.Errorf("发送消息失败：%v", err)
}

// SendText 发送文本消息
func (r *chatGroupRobot) SendText(ctx context.Context, content string, mentionedList []string, mentionedMobileList []string) error {
	message := Message{
		MsgType: MessageTypeText,
		Text: &TextMessage{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}
	return r.SendMessage(ctx, message)
}

// SendMarkdown 发送Markdown消息
func (r *chatGroupRobot) SendMarkdown(ctx context.Context, content string) error {
	message := Message{
		MsgType: MessageTypeMarkdown,
		Markdown: &MarkdownMessage{
			Content: content,
		},
	}
	return r.SendMessage(ctx, message)
}

// SendImage 发送图片消息
func (r *chatGroupRobot) SendImage(ctx context.Context, base64 string, md5 string) error {
	message := Message{
		MsgType: MessageTypeImage,
		Image: &ImageMessage{
			Base64: base64,
			MD5:    md5,
		},
	}
	return r.SendMessage(ctx, message)
}

// SendNews 发送图文消息
func (r *chatGroupRobot) SendNews(ctx context.Context, articles []NewsArticle) error {
	message := Message{
		MsgType: MessageTypeNews,
		News: &NewsMessage{
			Articles: articles,
		},
	}
	return r.SendMessage(ctx, message)
}

// SendFile 发送文件消息
func (r *chatGroupRobot) SendFile(ctx context.Context, mediaID string) error {
	message := Message{
		MsgType: MessageTypeFile,
		File: &FileMessage{
			MediaID: mediaID,
		},
	}
	return r.SendMessage(ctx, message)
}
