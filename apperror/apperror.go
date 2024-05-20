package apperror

type AppError struct {
	Code int
	Msg  string
}

func (e AppError) Error() string {
	return e.Msg
}

func (e AppError) Is(target error) bool {
	_, ok := target.(AppError)
	return ok
}

func New(code int, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

func NewByError(code int, err error) AppError {
	return AppError{Code: code, Msg: err.Error()}
}
