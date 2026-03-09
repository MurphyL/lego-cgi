package wecom

import (
	"testing"
)

func TestNewWecomChatBot(t *testing.T) {
	robot := NewChatBot("c080ed83-cf23-46cf-a8f7-fc3e1a28f9b5")
	t.Run("SendTextMessage", func(t *testing.T) {
		err := robot.SendTextMessage("发布消息", AtAll())
		if err != nil {
			t.Log("发从消息出错:", err)
			return
		}
	})
	t.Run("SendMarkdownMessage", func(t *testing.T) {
		err := robot.SendMarkdownMessage("发布消息")
		if err != nil {
			t.Log("发从消息出错:", err)
			return
		}
	})
}
