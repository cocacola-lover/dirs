package listener

type WrongMethodError struct {
	Message string
}

func (e WrongMethodError) Error() string {
	return e.Message
}
