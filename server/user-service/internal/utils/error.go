package utils

type CustomError struct {
	Name    string
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewError(name, msg string) *CustomError {
	return &CustomError{
		Name:    name,
		Message: msg,
	}
}

func NewCodeError(name string) *CustomError {
	return &CustomError{
		Name: name,
	}
}
