package utils

func ErrorMessage(err error, message string) string {
	msg := ""
	if err != nil {
		msg = err.Error()
	} else {
		msg = message
	}
	return msg
}
