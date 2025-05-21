package handler

import (
	"fmt"
	"qq-liftbot/bot/model"
	"strings"
)

// DispatchCommand 路由指令到对应处理逻辑
func DispatchCommand(userID int64, content string) {
    fields := strings.Fields(content)
    if len(fields) == 0 {
        return
    }

    cmd := strings.ToLower(fields[0])

    switch cmd {
    case "/运动":
        if len(fields) >= 2 {
            fmt.Println(model.NewActivity(userID, "运动", "运动记录", map[string]interface{}{"detail": fields[1]}))
        }
    case "/记账":
        if len(fields) >= 3 {
            fmt.Println(model.NewLogEvent(userID, "finance", "记账", map[string]interface{}{"item": fields[1], "cost": fields[2]}))
        }
    case "/就医":
        if len(fields) >= 4 {
            fmt.Println(model.NewActivity(userID, "就医", "门诊记录", map[string]interface{}{"hospital": fields[1], "package": fields[2], "cost": fields[3]}))
        }
    default:
        fmt.Println("❓ 未知指令")
    }
}
