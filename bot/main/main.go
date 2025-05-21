package main

import (
    "qq-liftbot/bot/handler"
    "qq-liftbot/bot/simulator"
    "fmt"
)

func main() {
    fmt.Println("🧪 NapCat Bot 模拟器启动")

    for _, raw := range simulator.MockMessages {
        handler.HandleMessage(raw)
    }
}
