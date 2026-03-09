package notify

/* 通知渠道（Notification Channels） */
type NotificationChannel string

const (
	// 邮件
	ChannelEmail NotificationChannel = "email"
	// 短信
	ChannelSMS NotificationChannel = "sms"
	// 企业微信群组会话机器人
	ChannelWecomChatBot NotificationChannel = "wecom_chat_bot"
	// 飞书机群组会话器人
	ChannelFeishuChatBot NotificationChannel = "feishu_chat_bot"
)

type NotificationMessage struct {
	Message string
}

type NotificationManager struct {
}

func (m *NotificationManager) Send(message NotificationMessage, channels ...NotificationChannel) {

}
