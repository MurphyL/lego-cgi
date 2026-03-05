package dingtalk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// https://oapi.dingtalk.com/robot/send?access_token=ACCESS_TOKEN

const (
	DingtalkChatGroupRobotWebhook = "https://oapi.dingtalk.com/robot/send"
	DefaultTimeout                = 10 * time.Second
	DefaultMaxRetries             = 3
)

// MessageType 消息类型
type MessageType string

// 消息类型常量
const (
	MessageTypeText       MessageType = "text"       // 文本消息
	MessageTypeMarkdown   MessageType = "markdown"   // Markdown消息
	MessageTypeLink       MessageType = "link"       // 链接消息
	MessageTypeActionCard MessageType = "actionCard" // 行动卡片消息
	MessageTypeFeedCard   MessageType = "feedCard"   //  Feed卡片消息
)

// TextMessage 文本消息
type TextMessage struct {
	Content string `json:"content"` // 消息内容
	At      struct {
		AtMobiles []string `json:"atMobiles,omitempty"` // 被@的手机号
		AtUserIds []string `json:"atUserIds,omitempty"` // 被@的用户ID
		IsAtAll   bool     `json:"isAtAll,omitempty"`   // 是否@所有人
	} `json:"at,omitempty"`
}

// MarkdownMessage Markdown消息
type MarkdownMessage struct {
	Title string `json:"title"` // 标题
	Text  string `json:"text"`  // Markdown内容
	At    struct {
		AtMobiles []string `json:"atMobiles,omitempty"` // 被@的手机号
		AtUserIds []string `json:"atUserIds,omitempty"` // 被@的用户ID
		IsAtAll   bool     `json:"isAtAll,omitempty"`   // 是否@所有人
	} `json:"at,omitempty"`
}

// LinkMessage 链接消息
type LinkMessage struct {
	Title      string `json:"title"`      // 标题
	Text       string `json:"text"`       // 内容
	MessageURL string `json:"messageUrl"` // 消息链接
	PicURL     string `json:"picUrl"`     // 图片链接
}

// ActionCardMessage 行动卡片消息
type ActionCardMessage struct {
	Title          string `json:"title"`          // 标题
	Text           string `json:"text"`           // 内容
	BtnOrientation string `json:"btnOrientation"` // 按钮排列方向（0：垂直，1：水平）
	Btns           []struct {
		Title     string `json:"title"`     // 按钮标题
		ActionURL string `json:"actionURL"` // 按钮链接
	} `json:"btns"` // 按钮列表
}

// FeedCardMessage Feed卡片消息
type FeedCardMessage struct {
	Links []struct {
		Title      string `json:"title"`      // 标题
		MessageURL string `json:"messageURL"` // 消息链接
		PicURL     string `json:"picURL"`     // 图片链接
	} `json:"links"` // 链接列表
}

// Message 消息结构
type Message struct {
	MsgType    MessageType        `json:"msgtype"`              // 消息类型
	Text       *TextMessage       `json:"text,omitempty"`       // 文本消息
	Markdown   *MarkdownMessage   `json:"markdown,omitempty"`   // Markdown消息
	Link       *LinkMessage       `json:"link,omitempty"`       // 链接消息
	ActionCard *ActionCardMessage `json:"actionCard,omitempty"` // 行动卡片消息
	FeedCard   *FeedCardMessage   `json:"feedCard,omitempty"`   // Feed卡片消息
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
func NewChatBot(accessToken string, secret string) (ChatBot, error) {
	robotWebhook, err := url.Parse(DingtalkChatGroupRobotWebhook)
	if err != nil {
		return nil, fmt.Errorf("解析URL出错：%v", err.Error())
	}

	// 构建查询参数
	params := url.Values{}
	params.Add("access_token", accessToken)

	// 如果有密钥，需要生成签名
	if secret != "" {
		timestamp := time.Now().UnixMilli()
		signature := generateSignature(secret, timestamp)
		params.Add("timestamp", fmt.Sprintf("%d", timestamp))
		params.Add("sign", signature)
	}

	robotWebhook.RawQuery = params.Encode()
	robotEndpoint := robotWebhook.String()

	return &chatGroupRobot{
		endpoint: robotEndpoint,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		maxRetries: DefaultMaxRetries,
	}, nil
}

// generateSignature 生成签名
func generateSignature(secret string, timestamp int64) string {
	// 构建签名字符串
	signStr := fmt.Sprintf("%d\n%s", timestamp, secret)

	// 使用HMAC-SHA256生成签名
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
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
	var atMobiles []string
	var atUserIds []string
	var isAtAll bool

	if opts, ok := options["at"]; ok {
		if atOpts, ok := opts.(map[string]interface{}); ok {
			if list, ok := atOpts["atMobiles"].([]string); ok {
				atMobiles = list
			}
			if userIds, ok := atOpts["atUserIds"].([]string); ok {
				atUserIds = userIds
			}
			if atAll, ok := atOpts["isAtAll"].(bool); ok {
				isAtAll = atAll
			}
		}
	}

	message := Message{
		MsgType: MessageTypeText,
		Text: &TextMessage{
			Content: content,
		},
	}

	message.Text.At.AtMobiles = atMobiles
	message.Text.At.AtUserIds = atUserIds
	message.Text.At.IsAtAll = isAtAll

	return r.SendMessage(ctx, message)
}

// SendMarkdown 发送Markdown消息
func (r *chatGroupRobot) SendMarkdown(ctx context.Context, content string, options map[string]interface{}) error {
	var title string
	var atMobiles []string
	var atUserIds []string
	var isAtAll bool

	if t, ok := options["title"].(string); ok {
		title = t
	}

	if opts, ok := options["at"]; ok {
		if atOpts, ok := opts.(map[string]interface{}); ok {
			if list, ok := atOpts["atMobiles"].([]string); ok {
				atMobiles = list
			}
			if userIds, ok := atOpts["atUserIds"].([]string); ok {
				atUserIds = userIds
			}
			if atAll, ok := atOpts["isAtAll"].(bool); ok {
				isAtAll = atAll
			}
		}
	}

	message := Message{
		MsgType: MessageTypeMarkdown,
		Markdown: &MarkdownMessage{
			Title: title,
			Text:  content,
		},
	}

	message.Markdown.At.AtMobiles = atMobiles
	message.Markdown.At.AtUserIds = atUserIds
	message.Markdown.At.IsAtAll = isAtAll

	return r.SendMessage(ctx, message)
}

// SendImage 发送图片消息
func (r *chatGroupRobot) SendImage(ctx context.Context, imageData map[string]string) error {
	// 钉钉机器人需要先上传图片获取media_id，这里简化处理
	return nil
}

// SendNews 发送图文消息
func (r *chatGroupRobot) SendNews(ctx context.Context, articles []map[string]string) error {
	// 钉钉没有直接的图文消息类型，使用链接消息或Feed卡片消息代替
	if len(articles) == 1 {
		article := articles[0]
		return r.SendLink(ctx, article)
	} else {
		return r.SendFeedCard(ctx, articles)
	}
}

// SendLink 发送链接消息
func (r *chatGroupRobot) SendLink(ctx context.Context, linkData map[string]string) error {
	message := Message{
		MsgType: MessageTypeLink,
		Link: &LinkMessage{
			Title:      linkData["title"],
			Text:       linkData["text"],
			MessageURL: linkData["messageUrl"],
			PicURL:     linkData["picUrl"],
		},
	}

	return r.SendMessage(ctx, message)
}

// SendActionCard 发送行动卡片消息
func (r *chatGroupRobot) SendActionCard(ctx context.Context, cardData map[string]interface{}) error {
	var title string
	var text string
	var btnOrientation string
	var btns []map[string]string

	if t, ok := cardData["title"].(string); ok {
		title = t
	}

	if t, ok := cardData["text"].(string); ok {
		text = t
	}

	if orientation, ok := cardData["btnOrientation"].(string); ok {
		btnOrientation = orientation
	}

	if buttons, ok := cardData["btns"].([]map[string]string); ok {
		btns = buttons
	}

	actionCard := &ActionCardMessage{
		Title:          title,
		Text:           text,
		BtnOrientation: btnOrientation,
		Btns: make([]struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		}, len(btns)),
	}

	for i, btn := range btns {
		actionCard.Btns[i].Title = btn["title"]
		actionCard.Btns[i].ActionURL = btn["actionURL"]
	}

	message := Message{
		MsgType:    MessageTypeActionCard,
		ActionCard: actionCard,
	}

	return r.SendMessage(ctx, message)
}

// SendFeedCard 发送Feed卡片消息
func (r *chatGroupRobot) SendFeedCard(ctx context.Context, links []map[string]string) error {
	feedCard := &FeedCardMessage{
		Links: make([]struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		}, len(links)),
	}

	for i, link := range links {
		feedCard.Links[i].Title = link["title"]
		feedCard.Links[i].MessageURL = link["messageUrl"]
		feedCard.Links[i].PicURL = link["picUrl"]
	}

	message := Message{
		MsgType:  MessageTypeFeedCard,
		FeedCard: feedCard,
	}

	return r.SendMessage(ctx, message)
}
