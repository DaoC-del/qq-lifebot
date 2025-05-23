package handler

import (
	"strconv"
	"strings"

	"qq-liftbot/bot/internal/errs"
	"qq-liftbot/bot/internal/result"
	"qq-liftbot/bot/model"
)

// DispatchCommand 是主命令处理器，返回结构化结果
func DispatchCommand(userID int64, content string) result.Result[string] {
	fields := strings.Fields(content)
	if len(fields) == 0 {
		return result.Err[string](errs.New(errs.ErrBadRequest, "空命令"))
	}

	cmd := strings.ToLower(fields[0])

	switch cmd {
	case "/运动":
		return handleExercise(userID, fields)
	case "/记账":
		return handleFinance(userID, fields)
	case "/就医":
		return handleMedical(userID, fields)
	default:
		return result.Err[string](errs.New(errs.ErrBadRequest, "未知指令"))
	}
}

func handleExercise(userID int64, fields []string) result.Result[string] {
	if len(fields) < 2 {
		return result.Err[string](errs.New(errs.ErrBadRequest, "运动命令参数不足"))
	}
	return model.NewActivity(userID, "运动", "运动记录", map[string]interface{}{
		"detail": fields[1],
	})
}

func handleFinance(userID int64, fields []string) result.Result[string] {
	if len(fields) < 3 {
		return result.Err[string](errs.New(errs.ErrBadRequest, "记账命令参数不足"))
	}
	cost, err := strconv.Atoi(fields[2])
	if err != nil {
		return result.Err[string](errs.New(errs.ErrBadRequest, "金额必须为整数"))
	}
	return model.NewLogEvent(userID, "finance", "记账", map[string]interface{}{
		"item": fields[1],
		"cost": cost,
	})
}

func handleMedical(userID int64, fields []string) result.Result[string] {
	if len(fields) < 4 {
		return result.Err[string](errs.New(errs.ErrBadRequest, "就医命令参数不足"))
	}
	return model.NewActivity(userID, "就医", "门诊记录", map[string]interface{}{
		"hospital": fields[1],
		"package":  fields[2],
		"cost":     fields[3],
	})
}
