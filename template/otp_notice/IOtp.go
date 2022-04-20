package otp_notice

import (
	"fmt"
	"math/rand"
	"time"
)

type IOtp interface {
	GeneratOtp()
	CacheOtp()
	OtpNotice()
	OtpLog()
}

type Otp struct {
	email string
	cache string
}

func NewOtp(email string, cache string) *Otp {
	return &Otp{email, cache}
}

func (t Otp) OtpLog() {
}

func (t Otp) OtpNotice() {
}

func (t Otp) GeneratOtp() {
	fmt.Println("Generator otp: " + fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)))
}
func (t Otp) CacheOtp() {
	//fmt.Println("Cache otp")
}

func GenAndSendOTP(e IOtp, c IOtp) {
	e.GeneratOtp()
	c.CacheOtp()
	e.OtpNotice()
	e.OtpLog()
}
