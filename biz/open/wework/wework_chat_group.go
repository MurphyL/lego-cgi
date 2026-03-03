package wework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"murphyl.com/lego/biz/open"
)

// 初始化时注册机器人工厂
func init() {
	open.RegisterChatBotFactory(open.BotTypeWework, &weworkChatBotFactory{})
}

// weworkChatBotFactory 企业微信机器人工厂
type weworkChatBotFactory struct{}

// CreateBot 创建企业微信机器人
func (f *weworkChatBotFactory) CreateBot(config open.BotConfig) (open.ChatBot, error) {
	bot, err := NewChatBot(config.RobotKey)
	if err != nil {
		return nil, err
	}
	return &weworkChatBotAdapter{bot: bot}, nil
}

// weworkChatBotAdapter 企业微信机器人适配器
type weworkChatBotAdapter struct {
	bot ChatBot
}

// CreateRequest 创建HTTP请求
func (a *weworkChatBotAdapter) CreateRequest(ctx context.Context, data any) (*http.Request, error) {
	return a.bot.CreateRequest(ctx, data)
}

// SendMessage 发送消息
func (a *weworkChatBotAdapter) SendMessage(ctx context.Context, message open.Message) error {
	return a.bot.SendMessage(ctx, message)
}

// SendText 发送文本消息
func (a *weworkChatBotAdapter) SendText(ctx context.Context, content string, options map[string]interface{}) error {
	return a.bot.SendText(ctx, content, options)
}

// SendMarkdown 发送Markdown消息
func (a *weworkChatBotAdapter) SendMarkdown(ctx context.Context, content string, options map[string]interface{}) error {
	return a.bot.SendMarkdown(ctx, content, options)
}

// SendImage 发送图片消息
func (a *weworkChatBotAdapter) SendImage(ctx context.Context, imageData map[string]string) error {
	return a.bot.SendImage(ctx, imageData)
}

// SendNews 发送图文消息
func (a *weworkChatBotAdapter) SendNews(ctx context.Context, articles []map[string]string) error {
	return a.bot.SendNews(ctx, articles)
}

// SendLink 发送链接消息
func (a *weworkChatBotAdapter) SendLink(ctx context.Context, linkData map[string]string) error {
	return a.bot.SendLink(ctx, linkData)
}

// SendActionCard 发送行动卡片消息
func (a *weworkChatBotAdapter) SendActionCard(ctx context.Context, cardData map[string]interface{}) error {
	return a.bot.SendActionCard(ctx, cardData)
}

// SendFeedCard 发送Feed卡片消息
func (a *weworkChatBotAdapter) SendFeedCard(ctx context.Context, links []map[string]string) error {
	return a.bot.SendFeedCard(ctx, links)
}

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

// ChatBot 群机器人接口
type ChatBot interface {
	// CreateRequest 创建HTTP请求
	CreateRequest(ctx context.Context, data any) (*http.Request, error)
	// SendMessage 发送消息
	SendMessage(ctx context.Context, message interface{}) error
	// SendText 发送文本消息
	SendText(ctx context.Context, content string, options map[string]interface{}) error
	// SendMarkdown 发送Markdown消息
	SendMarkdown(ctx context.Context, content string, options map[string]interface{}) error
	// SendImage 发送图片消息
	SendImage(ctx context.Context, imageData map[string]string) error
	// SendNews 发送图文消息
	SendNews(ctx context.Context, articles []map[string]string) error
	// SendLink 发送链接消息
	SendLink(ctx context.Context, linkData map[string]string) error
	// SendActionCard 发送行动卡片消息
	SendActionCard(ctx context.Context, cardData map[string]interface{}) error
	// SendFeedCard 发送Feed卡片消息
	SendFeedCard(ctx context.Context, links []map[string]string) error
}

// chatGroupRobot 群机器人实现
type chatGroupRobot struct {
	endpoint   string
	httpClient *http.Client
	maxRetries int
}

// NewChatBot 创建群机器人
func NewChatBot(robotKey string) (ChatBot, error) {
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
func (r *chatGroupRobot) SendMessage(ctx context.Context, message interface{}) error {
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
func (r *chatGroupRobot) SendText(ctx context.Context, content string, options map[string]interface{}) error {
	var mentionedList []string
	var mentionedMobileList []string

	if opts, ok := options["at"]; ok {
		if atOpts, ok := opts.(map[string]interface{}); ok {
			if list, ok := atOpts["mentionedList"].([]string); ok {
				mentionedList = list
			}
			if mobileList, ok := atOpts["mentionedMobileList"].([]string); ok {
				mentionedMobileList = mobileList
			}
		}
	}

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
func (r *chatGroupRobot) SendMarkdown(ctx context.Context, content string, options map[string]interface{}) error {
	message := Message{
		MsgType: MessageTypeMarkdown,
		Markdown: &MarkdownMessage{
			Content: content,
		},
	}
	return r.SendMessage(ctx, message)
}

// SendImage 发送图片消息
func (r *chatGroupRobot) SendImage(ctx context.Context, imageData map[string]string) error {
	base64, _ := imageData["base64"]
	md5, _ := imageData["md5"]

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
func (r *chatGroupRobot) SendNews(ctx context.Context, articles []map[string]string) error {
	var weworkArticles []NewsArticle

	for _, article := range articles {
		weworkArticles = append(weworkArticles, NewsArticle{
			Title:       article["title"],
			Description: article["description"],
			URL:         article["url"],
			PicURL:      article["picUrl"],
		})
	}

	message := Message{
		MsgType: MessageTypeNews,
		News: &NewsMessage{
			Articles: weworkArticles,
		},
	}
	return r.SendMessage(ctx, message)
}

// SendLink 发送链接消息
func (r *chatGroupRobot) SendLink(ctx context.Context, linkData map[string]string) error {
	// 企业微信没有直接的链接消息类型，使用图文消息代替
	articles := []map[string]string{
		{
			"title":       linkData["title"],
			"description": linkData["text"],
			"url":         linkData["messageUrl"],
			"picUrl":      linkData["picUrl"],
		},
	}

	return r.SendNews(ctx, articles)
}

// SendActionCard 发送行动卡片消息
func (r *chatGroupRobot) SendActionCard(ctx context.Context, cardData map[string]interface{}) error {
	// 企业微信没有直接的行动卡片消息类型，使用Markdown消息代替
	content := ""
	if c, ok := cardData["text"].(string); ok {
		content = c
	}

	return r.SendMarkdown(ctx, content, nil)
}

// SendFeedCard 发送Feed卡片消息
func (r *chatGroupRobot) SendFeedCard(ctx context.Context, links []map[string]string) error {
	// 企业微信没有直接的Feed卡片消息类型，使用图文消息代替
	return r.SendNews(ctx, links)
}
