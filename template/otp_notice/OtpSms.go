package otp_notice

import "fmt"

type OtpSms struct {
	Otp
}

func NewOtpSms() *OtpSms {
	return &OtpSms{}
}

func (t OtpSms) OtpNotice() {
	fmt.Println("SMS otp notice")
}

func (t OtpSms) OtpLog() {
	fmt.Println("SMS otp log")
}