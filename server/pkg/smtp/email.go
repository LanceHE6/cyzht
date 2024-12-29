package smtp

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"server/pkg/logger"
	"server/pkg/smtp/email_template"
)

var SMTPService *SMTP

type SMTP struct {
	UserName string // 发送消息的用户名
	D        *gomail.Dialer
}

// EmailType 邮件类型
type EmailType int

const (
	RegisterEmail      EmailType = 1 // 注册验证码邮件
	LoginEmail         EmailType = 2 // 登录验证码邮件
	ResetPasswordEmail EmailType = 3 // 重置密码验证码邮件
)

// InitSMTPService
//
//	@Description: 初始化 SMTP 服务
//	@param host SMTP 服务器地址
//	@param port SMTP 服务器端口
//	@param userName SMTP 服务器用户名
//	@param password SMTP 服务器密码
//	@return *smtpService SMTP 服务
func InitSMTPService(host string, port int, userName, password string) {
	SMTPService = &SMTP{
		UserName: userName,
		D:        gomail.NewDialer(host, port, userName, password),
	}
}

// SendVerifyCodeEmail
//
//	@Description: 发送邮箱邮件
//	@param targetEmail  目标邮箱
//	@param account 目标名称
//	@param code 验证码
//	@param emailType 邮件类别
//	@return error 错误信息
func (s *SMTP) SendVerifyCodeEmail(targetEmail string, account string, code string, emailType EmailType) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "NetChat"+"<"+s.UserName+">")
	m.SetHeader("To", targetEmail)

	if emailType == RegisterEmail {
		message := email_template.GetVerifyEmailHTML(targetEmail, code)

		m.SetHeader("Subject", "注册验证码："+code)
		m.SetBody("text/html", message)

		s.D.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := s.D.DialAndSend(m); err != nil {
			logger.Logger.Error("Failed to send register email:" + err.Error())
			return err
		} else {
			logger.Logger.Debug("Email sent successfully")
			return nil
		}
	}
	if emailType == LoginEmail {
		message := email_template.GetLoginVerifyCodeEmailHTML(targetEmail, code)

		m.SetHeader("Subject", "登录验证码："+code)
		m.SetBody("text/html", message)

		s.D.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := s.D.DialAndSend(m); err != nil {
			logger.Logger.Error("Failed to send register email:" + err.Error())
			return err
		} else {
			logger.Logger.Debug("Email sent successfully")
			return nil
		}
	}
	if emailType == ResetPasswordEmail {
		message := email_template.GetResetPasswordEmailHTML(account, code)

		m.SetHeader("Subject", "重置账号密码验证码："+code)
		m.SetBody("text/html", message)

		s.D.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := s.D.DialAndSend(m); err != nil {
			logger.Logger.Error("Failed to send reset psw email:" + err.Error())
			return err
		} else {
			logger.Logger.Debug("Email sent successfully")
			return nil
		}
	}
	return nil
}

// SendTemporaryPswEmail
//
//	@Description: 注册成功发送临时密码邮件
//	@param targetEmail 目标邮箱
//	@param psw 临时密码
//	@return error 错误信息
func (s *SMTP) SendTemporaryPswEmail(targetEmail string, psw string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "NetChat"+"<"+s.UserName+">")
	m.SetHeader("To", targetEmail)

	message := email_template.GetTempPswEmailHTML(targetEmail, psw)

	m.SetHeader("Subject", "注册成功")
	m.SetBody("text/html", message)

	s.D.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := s.D.DialAndSend(m); err != nil {
		logger.Logger.Error("Failed to send register email:" + err.Error())
		return err
	} else {
		logger.Logger.Debug("Email sent successfully")
		return nil
	}
}
