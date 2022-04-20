package otp_notice

import "fmt"

type OtpEmail struct {
	Otp
}

func NewOtpEmail() *OtpEmail {
	return &OtpEmail{}
}

func (t OtpEmail) OtpNotice() {
	fmt.Println("Email otp notice")
}

func (t OtpEmail) OtpLog() {
	fmt.Println("Email otp log")
}