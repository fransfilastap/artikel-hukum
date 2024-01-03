package utils

type Email struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(email Email) error {
	return nil
}
