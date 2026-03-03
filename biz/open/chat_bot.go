package open

import (
	"context"
	"net/http"
)

// ChatBot 群机器人统一接口
type ChatBot interface {
	// CreateRequest 创建HTTP请求
	CreateRequest(ctx context.Context, data any) (*http.Request, error)
	// SendMessage 发送消息
	SendMessage(ctx context.Context, message Message) error
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

// Message 消息结构
type Message struct {
	MsgType        string `json:"msgtype"`                  // 消息类型
	Content        any    `json:"content,omitempty"`        // 消息内容
	At             any    `json:"at,omitempty"`             // @信息
	Title          string `json:"title,omitempty"`          // 标题
	Text           string `json:"text,omitempty"`           // 文本
	MessageURL     string `json:"messageUrl,omitempty"`     // 消息链接
	PicURL         string `json:"picUrl,omitempty"`         // 图片链接
	Base64         string `json:"base64,omitempty"`         // 图片base64编码
	MD5            string `json:"md5,omitempty"`            // 图片md5值
	MediaID        string `json:"media_id,omitempty"`       // 文件媒体ID
	Articles       any    `json:"articles,omitempty"`       // 文章列表
	BtnOrientation string `json:"btnOrientation,omitempty"` // 按钮排列方向
	Btns           any    `json:"btns,omitempty"`           // 按钮列表
	Links          any    `json:"links,omitempty"`          // 链接列表
}

// BotType 机器人类型
type BotType string

// 机器人类型常量
const (
	BotTypeWework   BotType = "wework"   // 企业微信
	BotTypeDingtalk BotType = "dingtalk" // 钉钉
)

// BotConfig 机器人配置
type BotConfig struct {
	Type        BotType // 机器人类型
	AccessToken string  // 访问令牌
	Secret      string  // 密钥
	RobotKey    string  // 机器人密钥
}

// NewChatBot 创建群机器人
func NewChatBot(config BotConfig) (ChatBot, error) {
	factory := GetChatBotFactory(config.Type)
	if factory == nil {
		return nil, nil
	}
	return factory.CreateBot(config)
}

// ChatBotFactory 机器人工厂接口
type ChatBotFactory interface {
	CreateBot(config BotConfig) (ChatBot, error)
}

// RegisterChatBotFactory 注册机器人工厂
var chatBotFactories = make(map[BotType]ChatBotFactory)

// RegisterChatBotFactory 注册机器人工厂
func RegisterChatBotFactory(botType BotType, factory ChatBotFactory) {
	chatBotFactories[botType] = factory
}

// GetChatBotFactory 获取机器人工厂
func GetChatBotFactory(botType BotType) ChatBotFactory {
	return chatBotFactories[botType]
}
