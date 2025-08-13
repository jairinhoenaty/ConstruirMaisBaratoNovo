package mail

type EmailService interface {
	SendEmail(to, subject, body string) error
}
