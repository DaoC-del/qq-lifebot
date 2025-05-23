package errs

import "fmt"

type Error struct {
	Kind    Kind
	Message string
	Context string // 可选上下文（可用于扩展）
}

func (e Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Kind, e.Message)
}

func New(kind Kind, msg string) Error {
	return Error{Kind: kind, Message: msg}
}
