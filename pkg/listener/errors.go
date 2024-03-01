package listener

type _WrongMethodError struct {
	Message string
}

func (e _WrongMethodError) Error() string {
	return e.Message
}
