package email_service

import (
	"log"
	"strconv"

	"github.com/matcornic/hermes/v2"
)

const (
	ConfirmationEmail = "confirmation"
	ResetEmail        = "reset"
	ValidateEmail     = "validate"
	ActivationEmail   = "activation"
	TokenEmail        = "token"
	SignupEmail       = "signup"
	OTPEmail          = "otp"
	OrderEmail        = "order"
	PurchaseEmail     = "purchase"
	InviteEmail       = "invite"
	MarketingEmail    = "marketing"
	PromotionEmail    = "promotion"
	CustomerEmail     = "customer"
	ProductEmail      = "product"
	BillEmail         = "bill"
	GreetingEmail     = "greeting"
	ReceiptEmail      = "receipt"
	NotificationEmail = "notification"
	WelcomeEmail      = "welcome"
	NewsletterEmail   = "newsletter"
	MaintenanceEmail  = "maintenance"
	LoginEmail        = "login"
	APIEmail          = "api"
	SubscribeEmail    = "subscribe"
	SubscriptionEmail = "subscription"
	DefaultEmail      = "default"
)

type confirmation struct {
	Intent                  string
	CC                      string
	Company                 string
	Date                    string
	HasTableData            bool
	ProductLink             string
	CompanyLogoLink         string
	RecepientName           string
	Copyright               string
	TroubelText             string
	CompanyName      string
	VerificationLink string
	OutrosMessages          []string
	IntrosMessages          []string
	Signature               string
	Title                   string
	Greeting                string
	IntroBody               string
	WebsiteLink             string
	MailIntent              string
	MailSubject             string
	ValiationLink           string
	ConfirmationLink        string
	OTPData                 string
	InstructionActionText   string
	InviteCode              string
	DashboardButtonLink     string
	DashboardButtonText     string
	DashboardButtonColor    string
	DashboardButtonTexColor string
	CallForAction           string
	EmailSubject            string
	SupportEmail            string
	SenderEmail             string
	CopyrightLink           string
	TermsAndConditionLink   string
	MaxTableRows            int
	MaxTableColumns         int
	KeyValue                map[string]string
	KeyValueTableEntry      []KeyValueTableEntry
	KeyValueTableEntryMap   map[string][]KeyValueTableEntry
}

type KeyValueTableEntry struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

func (r *confirmation) BuildHermesEmailTableEntryKeys() (hm [][]hermes.Entry) {

	var mailHm [][]hermes.Entry
	hm = make([][]hermes.Entry, 0)
	hmRow1 := make([][]hermes.Entry, 0)
	hmRow2 := make([][]hermes.Entry, 0)
	hmRow3 := make([][]hermes.Entry, 0)
	hmRow4 := make([][]hermes.Entry, 0)

	if r.HasTableData {

		log.Println(r.KeyValueTableEntryMap)

		for k, vmap := range r.KeyValueTableEntryMap {

			switch k {
			case "row1":
				for _, kvmap := range vmap {
					log.Println("Appending ...  { key :" + kvmap.Key + ", Value :" + kvmap.Value + "}")
					hmRow1 = append(hmRow1, []hermes.Entry{
						{
							Key: kvmap.Key, Value: kvmap.Value,
						},
					})

				}

			case "row2":
				for _, kvmap := range vmap {
					log.Println("Appending ...  { key :" + kvmap.Key + ", Value :" + kvmap.Value + "}")
					hmRow2 = append(hmRow2, []hermes.Entry{
						{
							Key: kvmap.Key, Value: kvmap.Value,
						},
					})

				}

			case "row3":
				for _, kvmap := range vmap {
					log.Println("Appending ...  { key :" + kvmap.Key + ", Value :" + kvmap.Value + "}")
					hmRow3 = append(hmRow3, []hermes.Entry{
						{
							Key: kvmap.Key, Value: kvmap.Value,
						},
					})

				}

			case "row4":
				for _, kvmap := range vmap {
					log.Println("Appending ...  { key :" + kvmap.Key + ", Value :" + kvmap.Value + "}")
					hmRow4 = append(hmRow4, []hermes.Entry{
						{
							Key: kvmap.Key, Value: kvmap.Value,
						},
					})

				}
			}

		}

		if r.MaxTableRows == 3 && r.MaxTableColumns == 1 {

			log.Println("The length of ROW1:" + strconv.Itoa(len(hmRow1)))
			mailHm = [][]hermes.Entry{
				{
					{Key: hmRow1[0][0].Key, Value: hmRow1[0][0].Value},
					{Key: hmRow1[1][0].Key, Value: hmRow1[1][0].Value},
					{Key: hmRow1[2][0].Key, Value: hmRow1[2][0].Value},
				},
			}

			log.Println(">>>>>>>>>>")
			log.Println(hmRow1[0][0].Key)
			log.Println(hmRow1[1][0].Key)
			log.Println(hmRow1[2][0].Key)
			log.Println(">>>>>>>>>>")

		}

		log.Println(hmRow1)
		if r.MaxTableRows == 3 && r.MaxTableColumns == 2 {

			log.Println("The length of ROW1:" + strconv.Itoa(len(hmRow1)))
			log.Println("The length of ROW2:" + strconv.Itoa(len(hmRow2)))

			//mailHm=make([][]hermes.Entry, 0)
			mailHm = [][]hermes.Entry{
				{
					{Key: hmRow1[0][0].Key, Value: hmRow1[0][0].Value},
					{Key: hmRow1[1][0].Key, Value: hmRow1[1][0].Value},
					{Key: hmRow1[2][0].Key, Value: hmRow1[2][0].Value},
				},
				{
					{Key: hmRow2[0][0].Key, Value: hmRow2[0][0].Value},
					{Key: hmRow2[1][0].Key, Value: hmRow2[1][0].Value},
					{Key: hmRow2[2][0].Key, Value: hmRow2[2][0].Value},
				},
			}

			log.Println(">>>>>>>>>>")
			log.Println(hmRow1[0][0].Key)
			log.Println(hmRow1[1][0].Key)
			log.Println(hmRow1[2][0].Key)
			log.Println("=============")
			log.Println(hmRow2[0][0].Key)
			log.Println(hmRow2[1][0].Key)
			log.Println(hmRow2[2][0].Key)
			log.Println(">>>>>>>>>>")

		}

		if r.MaxTableRows == 3 && r.MaxTableColumns == 3 {

			log.Println("The length of ROW1:" + strconv.Itoa(len(hmRow1)))
			log.Println("The length of ROW2:" + strconv.Itoa(len(hmRow2)))
			log.Println("The length of ROW3:" + strconv.Itoa(len(hmRow3)))

			mailHm = [][]hermes.Entry{
				{
					{Key: hmRow1[0][0].Key, Value: hmRow1[0][0].Value},
					{Key: hmRow1[1][0].Key, Value: hmRow1[1][0].Value},
					{Key: hmRow1[2][0].Key, Value: hmRow1[2][0].Value},
				},
				{
					{Key: hmRow2[0][0].Key, Value: hmRow2[0][0].Value},
					{Key: hmRow2[1][0].Key, Value: hmRow2[1][0].Value},
					{Key: hmRow2[2][0].Key, Value: hmRow2[2][0].Value},
				}, {
					{Key: hmRow3[0][0].Key, Value: hmRow3[0][0].Value},
					{Key: hmRow3[1][0].Key, Value: hmRow3[1][0].Value},
					{Key: hmRow3[2][0].Key, Value: hmRow3[2][0].Value},
				},
			}

			log.Println(">>>>>>>>>>")
			log.Println(hmRow1[0][0].Key)
			log.Println(hmRow1[1][0].Key)
			log.Println(hmRow1[2][0].Key)
			log.Println("=============")
			log.Println(hmRow2[0][0].Key)
			log.Println(hmRow2[1][0].Key)
			log.Println(hmRow2[2][0].Key)
			log.Println("=============")
			log.Println(hmRow3[0][0].Key)
			log.Println(hmRow3[1][0].Key)
			log.Println(hmRow3[2][0].Key)
			log.Println(">>>>>>>>>>")

		}

		if r.MaxTableRows == 3 && r.MaxTableColumns == 4 {

			log.Println("The length of ROW1:" + strconv.Itoa(len(hmRow1)))
			log.Println("The length of ROW2:" + strconv.Itoa(len(hmRow2)))
			log.Println("The length of ROW3:" + strconv.Itoa(len(hmRow3)))
			log.Println("The length of ROW4:" + strconv.Itoa(len(hmRow4)))

			mailHm = [][]hermes.Entry{
				{
					{Key: hmRow1[0][0].Key, Value: hmRow1[0][0].Value},
					{Key: hmRow1[1][0].Key, Value: hmRow1[1][0].Value},
					{Key: hmRow1[2][0].Key, Value: hmRow1[2][0].Value},
				},
				{
					{Key: hmRow2[0][0].Key, Value: hmRow2[0][0].Value},
					{Key: hmRow2[1][0].Key, Value: hmRow2[1][0].Value},
					{Key: hmRow2[2][0].Key, Value: hmRow2[2][0].Value},
				}, {
					{Key: hmRow3[0][0].Key, Value: hmRow3[0][0].Value},
					{Key: hmRow3[1][0].Key, Value: hmRow3[1][0].Value},
					{Key: hmRow3[2][0].Key, Value: hmRow3[2][0].Value},
				}, {
					{Key: hmRow4[0][0].Key, Value: hmRow4[0][0].Value},
					{Key: hmRow4[1][0].Key, Value: hmRow4[1][0].Value},
					{Key: hmRow4[2][0].Key, Value: hmRow4[2][0].Value},
				},
			}

			log.Println(">>>>>>>>>>")
			log.Println(hmRow1[0][0].Key)
			log.Println(hmRow1[1][0].Key)
			log.Println(hmRow1[2][0].Key)
			log.Println("=============")
			log.Println(hmRow2[0][0].Key)
			log.Println(hmRow2[1][0].Key)
			log.Println(hmRow2[2][0].Key)
			log.Println("=============")
			log.Println(hmRow3[0][0].Key)
			log.Println(hmRow3[1][0].Key)
			log.Println(hmRow3[2][0].Key)
			log.Println("=============")
			log.Println(hmRow4[0][0].Key)
			log.Println(hmRow4[1][0].Key)
			log.Println(hmRow4[2][0].Key)
			log.Println(">>>>>>>>>>")

		}
	} else {
		log.Println("No table data ")
	}

	for i, v := range hm {
		log.Printf("Index=%v Valute %v", i, v)
	}

	return mailHm
}

func (r *confirmation) CopyWith(confirmationData confirmation) confirmation {
	r.Intent = confirmationData.Intent
	r.ValiationLink=confirmationData.ValiationLink
	r.CompanyLogoLink=confirmationData.CompanyLogoLink
	r.ConfirmationLink=confirmationData.ConfirmationLink
	r.ProductLink=confirmationData.ProductLink
	r.Greeting = confirmationData.Greeting
	r.Title = confirmationData.Title
	r.MailIntent = confirmationData.MailIntent
	r.MailSubject = confirmationData.MailSubject
	r.Signature = confirmationData.Signature
	r.EmailSubject = confirmationData.EmailSubject
	r.HasTableData = confirmationData.HasTableData
	r.MaxTableRows = confirmationData.MaxTableRows
	r.MaxTableColumns = confirmationData.MaxTableColumns
	r.KeyValueTableEntry = confirmationData.KeyValueTableEntry
	r.KeyValueTableEntryMap = confirmationData.KeyValueTableEntryMap
	r.RecepientName = confirmationData.RecepientName
	r.EmailSubject = confirmationData.EmailSubject
	r.OTPData = confirmationData.OTPData
	r.WebsiteLink = confirmationData.WebsiteLink
	r.InviteCode = confirmationData.InviteCode
	r.CopyrightLink = confirmationData.CopyrightLink
	r.TermsAndConditionLink = confirmationData.TermsAndConditionLink
	r.SupportEmail = confirmationData.SupportEmail
	r.DashboardButtonLink = confirmationData.DashboardButtonLink
	r.DashboardButtonText = confirmationData.DashboardButtonText
	r.KeyValue = confirmationData.KeyValue
	r.ConfirmationLink = confirmationData.ConfirmationLink
	r.ValiationLink = confirmationData.ValiationLink
	r.CC = confirmationData.CC
	r.OutrosMessages = confirmationData.OutrosMessages
	r.DashboardButtonColor = confirmationData.DashboardButtonColor
	r.DashboardButtonTexColor = confirmationData.DashboardButtonTexColor
	r.Company = confirmationData.Company
	r.IntrosMessages = confirmationData.IntrosMessages
	r.InstructionActionText = confirmationData.InstructionActionText
	r.CallForAction = confirmationData.CallForAction

	return *r
}

func (r *confirmation) Name() string {

	switch r.Intent {
	case ConfirmationEmail:
		return ConfirmationEmail

	case InviteEmail:
		return InviteEmail

	case SubscribeEmail:

		return SubscribeEmail

	case ReceiptEmail:
		return ReceiptEmail

	case ResetEmail:
		return ResetEmail

	case WelcomeEmail:

		return WelcomeEmail

	case LoginEmail:
		return LoginEmail

	case MaintenanceEmail:
		return MaintenanceEmail

	case SignupEmail:

		return SignupEmail

	case SubscriptionEmail:

		return SubscriptionEmail

	case DefaultEmail:
		return DefaultEmail

	case MarketingEmail:

		return MarketingEmail

	case ProductEmail:
		return ProductEmail

	case PromotionEmail:
		return PromotionEmail

	case APIEmail:
		return APIEmail
	case CustomerEmail:
		return CustomerEmail

	case OrderEmail:
		return OrderEmail

	case PurchaseEmail:
		return PurchaseEmail

	case OTPEmail:
		return OTPEmail

	case ActivationEmail:

		return ActivationEmail

	case TokenEmail:

		return TokenEmail

	case ValidateEmail:

		return ValidateEmail

	}

	return "confirmation"
}

func (r *confirmation) Email() hermes.Email {
	log.Println(r)

	if len(r.RecepientName) == 0 {
		log.Panicln("No receipient name provided")
	}

	//var  mailTable hermes.Table

	var mailTableDataEntries [][]hermes.Entry

	var mailTableColumn hermes.Columns

	mailTableDataEntries = make([][]hermes.Entry, 0)

	for key, value := range r.KeyValue {

		mailTableDataEntries = append(mailTableDataEntries, []hermes.Entry{
			{
				Key: key, Value: value,
			}})

	}

	mailTableColumn = hermes.Columns{
		// CustomWidth: map[string]string{
		// 	"Item":  "20%",
		// 	"Price": "15%",
		// },
		// CustomAlignment: map[string]string{
		// 	"Price": "right",
		// },
	}

	// Populate Table entries

	if len(r.KeyValueTableEntry) == 0 {
		log.Panicln("No Table Entry map proviede")
	}

	confirmationTable := hermes.Table{
		Data:    r.BuildHermesEmailTableEntryKeys(),
		Columns: mailTableColumn,
	}

	return hermes.Email{
		Body: hermes.Body{
			Title:     r.Title,
			Outros:    r.OutrosMessages,
			Greeting:  r.Greeting,
			Name:      r.RecepientName,
			Signature: r.Signature,
			Intros:    r.IntrosMessages,
			Table:     confirmationTable,
			Actions: []hermes.Action{
				{
					InviteCode:   r.InviteCode,
					Instructions: r.InstructionActionText,
					Button: hermes.Button{
						Text:      r.DashboardButtonText,
						Link:      r.DashboardButtonLink,
						Color:     r.DashboardButtonColor,
						TextColor: r.DashboardButtonTexColor,
					},
				},
			},
		},
	}
}
