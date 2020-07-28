package httperror

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (e StatusError) Error() string {
	return e.Err.Error()
}

func (e StatusError) Status() int {
	return e.Code
}
