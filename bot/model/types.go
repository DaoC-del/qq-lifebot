package model

import (
    "encoding/json"
    "time"
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

func NewActivity(userID int64, actType, summary string, payload map[string]interface{}) string {
    a := Activity{
        UserID:    userID,
        Type:      actType,
        Summary:   summary,
        Payload:   payload,
        Timestamp: time.Now(),
    }
    b, _ := json.MarshalIndent(a, "", "  ")
    return "📘 Activity:" + string(b)
}

func NewLogEvent(userID int64, typ, summary string, payload map[string]interface{}) string {
    e := LogEvent{
        UserID:    userID,
        Type:      typ,
        Summary:   summary,
        Payload:   payload,
        CreatedAt: time.Now(),
    }
    b, _ := json.MarshalIndent(e, "", "  ")
    return "📝 LogEvent:" + string(b)
}
