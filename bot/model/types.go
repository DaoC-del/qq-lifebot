package model

import (
	"encoding/json"
	"time"

	"qq-liftbot/bot/internal/result"
)

type Activity struct {
	UserID    int64
	Type      string
	Summary   string
	Payload   map[string]interface{}
	Timestamp time.Time
}

type LogEvent struct {
	UserID    int64
	Type      string
	Summary   string
	Payload   map[string]interface{}
	CreatedAt time.Time
}

// 返回 Result[string]
func NewActivity(userID int64, actType, summary string, payload map[string]interface{}) result.Result[string] {
	a := Activity{
		UserID:    userID,
		Type:      actType,
		Summary:   summary,
		Payload:   payload,
		Timestamp: time.Now(),
	}
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return result.Err[string](err)
	}
	return result.Ok("📘 Activity:\n" + string(b))
}

func NewLogEvent(userID int64, typ, summary string, payload map[string]interface{}) result.Result[string] {
	e := LogEvent{
		UserID:    userID,
		Type:      typ,
		Summary:   summary,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return result.Err[string](err)
	}
	return result.Ok("📝 LogEvent:\n" + string(b))
}
