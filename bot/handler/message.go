package handler

import (
	"encoding/json"
	"fmt"
)

// OneBotMessage 定义模拟的 NapCat 消息格式
type OneBotMessage struct {
    MessageType string `json:"message_type"`
    UserID      int64  `json:"user_id"`
    Message     string `json:"message"`
}

// HandleMessage 解析并分发命令
func HandleMessage(raw string) {
    var msg OneBotMessage
    if err := json.Unmarshal([]byte(raw), &msg); err != nil {
        fmt.Println("❌ 解析失败:", err)
        return
    }

    fmt.Printf("📩 [%d] %s", msg.UserID, msg.Message)

    if len(msg.Message) > 0 && msg.Message[0] == '/' {
        DispatchCommand(msg.UserID, msg.Message)
    }
}
