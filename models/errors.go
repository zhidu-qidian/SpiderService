package models

type NoSuchObjectError struct {
	Info string
}

func (e NoSuchObjectError) Error() string {
	return e.Info
}

func NewNoSuchObjectError(s string) NoSuchObjectError {
	return NoSuchObjectError{Info: s}
}
