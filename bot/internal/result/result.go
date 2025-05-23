package result

// Result is your generic result type.
type Result[T any] struct {
    Value T
    Err   error
}

func Ok[T any](v T) Result[T] {
    return Result[T]{Value: v, Err: nil}
}

func Err[T any](e error) Result[T] {
    var zero T
    return Result[T]{Value: zero, Err: e}
}

func (r Result[T]) IsOk() bool {
    return r.Err == nil
}

func (r Result[T]) IsErr() bool {
    return r.Err != nil
}

// Map applies f to the Value if no error, otherwise propagates the error.
func Map[T, U any](r Result[T], f func(T) U) Result[U] {
    if r.IsErr() {
        return Err[U](r.Err)
    }
    return Ok(f(r.Value))
}

// AndThen chains a second operation that also returns a Result.
func AndThen[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
    if r.IsErr() {
        return Err[U](r.Err)
    }
    return f(r.Value)
}
