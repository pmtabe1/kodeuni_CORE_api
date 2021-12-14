package constants

//date formatters
const  AppName="subscription"
const Golang_date_YYYYMMDD = "2006-01-02"
const TRA_receipt_date_format = "2006-01-02"
const TRA_Znum_format = "20060102"
const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)
const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	HeaderNoEncoding = `<?xml version="1.0" ?>` + "\n"
)

//Time formats

const TRA_receipt_time_format = "2006-01-02 15:04:05"

//const TRA_receipt_time_format = "HH:MM:SS"

const Certkey = "10TZ100666"
const Tin = "112318380"
const Password = "uxte2!"
const TRACertSerial = "32213952251598420917675373467124718226"
const QRCODE_URI = "https://api.qrserver.com/v1/create-qr-code/?data=https://verify.tra.go.tz/721C5840775_212435&amp;size=100x100"
const QRCODE_BASE_URI = "https://api.qrserver.com/v1/create-qr-code/?data=https://verify.tra.go.tz/"

const Predefined_schedule_daily = "daily"
const Predefined_schedule_weekly = "weekly"
const Predefined_schedule_yearly = "yearly"
const Predefined_schedule_montly = "monthly"
const Predefined_schedule_hourly = "hourly"

const TRA_INTENT_RECEIPT = "receipt"
const TRA_INTENT_REGISTRATION = "registration"
const TRA_INTENT_TOKEN = "token"
const TRA_INTENT_VALIDATION = "validation"
