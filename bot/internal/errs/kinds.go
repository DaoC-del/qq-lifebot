package errs

type Kind int

const (
	ErrBadRequest Kind = 400
	ErrNotFound   Kind = 404
	ErrInternal   Kind = 500
	ErrNull       Kind = 600
	ErrIndexOut   Kind = 601
)
