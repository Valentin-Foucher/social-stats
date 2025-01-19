package common

type ErrorKind int

const (
	NotFound ErrorKind = iota
	InternalError
	InvalidParameters
	Unauthorized
)

type AppError struct {
	inner       error
	description string
	kind        ErrorKind
}

func (err *AppError) Error() string {
	return err.inner.Error()
}
func (err *AppError) Kind() ErrorKind {
	return err.kind
}
func (err *AppError) Description() string {
	return err.description
}

func newAppError(err error, descr string, kind ErrorKind) *AppError {
	return &AppError{
		inner:       err,
		description: descr,
		kind:        kind,
	}
}

func NewInvalidParametersError(err error, descr string) *AppError {
	return newAppError(err, descr, InvalidParameters)
}

func NewNotFoundError(err error, descr string) *AppError {
	return newAppError(err, descr, NotFound)
}

func NewInternalError(err error, descr string) *AppError {
	return newAppError(err, descr, InternalError)
}

func NewDefaultInternalError(err error) *AppError {
	return NewInternalError(err, "Internal Server Error")
}

func NewUnauthorizedError(err error) *AppError {
	return newAppError(err, "Unauthorized", Unauthorized)
}
