package main

import "template/otp_notice"

func main() {
	//otpSms := otp_notice.OtpSms{}

	//otp := otp_notice.Otp{IOtp: otpSms}

	var otpEmail otp_notice.IOtp
	otpEmail = otp_notice.NewOtpEmail()

	var otpSms otp_notice.IOtp
	otpSms = otp_notice.NewOtpSms()

	var memCache otp_notice.IOtp
	memCache = otp_notice.NewMemCache()

	var redisCache otp_notice.IOtp
	redisCache = otp_notice.NewRedisCache()

	otp_notice.GenAndSendOTP(otpEmail, memCache)
	otp_notice.GenAndSendOTP(otpSms, redisCache)
}
