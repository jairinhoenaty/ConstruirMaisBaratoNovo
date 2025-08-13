package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
	"log"
	"net/smtp"
)

type UserSendEmailUC struct {
	Service pkguser.UserService
	Email   string
}

type UserSendEmailUCParams struct {
	Service pkguser.UserService
}

func NewSendEmailUC(params UserSendEmailUCParams) UserSendEmailUC {
	return UserSendEmailUC{
		Service: params.Service,
	}
}

func (uc *UserSendEmailUC) Execute() error {

	//pesquisar o usuário pelo email
	params := FindByEmailUCParams{
		Service: uc.Service,
	}
	userUC := NewFindByEmailUC(params)
	userUC.Email = &uc.Email
	user, err := userUC.Execute()
	if err != nil {
		return err
	}

	//criptografar email
	encrypted := EncryptValue(user.Email)

	// 3. Enviar um email com o link de redefinição
	resetLink := fmt.Sprintf("https://construirmaisbarato.com.br/confirmar-senha/%s", encrypted)
	emailBody := fmt.Sprintf("A sua solicitação de redefinição de senha será confirmada pelo link abaixo.\n\nClique no link para recuperar sua senha:\n%s\n\n Se você não solicitou esta redefinição, por favor, ignore este e-mail.\n\nAtenciosamente \n\n Construir Mais Barato", resetLink)

	config := SMTPConfig{
		Host:     "smtp.zoho.com",
		Port:     "587",
		Username: "atendimento@construirmaisbarato.com.br",
		Password: "Jh230472*",
	}

	data := EmailData{
		From:    "atendimento@construirmaisbarato.com.br",
		To:      user.Email,
		Subject: "Recuperar senha",
		Text:    emailBody,
	}

	err = SendEmail(config, data)
	if err != nil {
		log.Fatalf("Erro ao enviar e-mail: %v", err)
	} else {
		log.Println("E-mail enviado com sucesso!")
		return err
	}

	return nil
}

// EmailData contém as informações necessárias para enviar um e-mail.
type EmailData struct {
	From    string
	To      string
	Subject string
	Text    string
}

// SMTPConfig contém a configuração do servidor SMTP.
type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

// SendEmail envia um e-mail com base nas configurações SMTP fornecidas e nos dados do e-mail.
func SendEmail(config SMTPConfig, data EmailData) error {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	// Montar o corpo do e-mail
	msg := []byte("From: " + data.From + "\r\n" +
		"To: " + data.To + "\r\n" +
		"Subject: " + data.Subject + "\r\n" +
		"\r\n" +
		data.Text + "\r\n")

	// Enviar o e-mail
	err := smtp.SendMail(config.Host+":"+config.Port, auth, data.From, []string{data.To}, msg)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nil
}
