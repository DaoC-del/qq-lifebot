package simulator

// 模拟 NapCat 推送的消息（OneBot 格式）
var MockMessages = []string{
    `{"message_type":"private","user_id":123456,"message":"/运动 30分钟"}`,
    `{"message_type":"private","user_id":123456,"message":"/记账 早餐 500"}`,
    `{"message_type":"private","user_id":123456,"message":"/就医 品川诊所 体检A 3500"}`,
    `{"message_type":"private","user_id":123456,"message":"/未知指令 测试"}`,
}
