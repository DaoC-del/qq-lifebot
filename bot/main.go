package main

import (
    "qq-liftbot/bot/handler"
    "qq-liftbot/bot/simulator"
    "fmt"
)

func main() {
    fmt.Println("🧪 启动模拟 NapCat QQ 客户端...")

    for _, mock := range simulator.MockMessages {
        handler.HandleMessage(mock)
    }
}
