package mappers_utils

import "log"

func GetMappedPaymentType(lookup int) (paymentType string) {

	// 1. CASH
	// 2. CHEQUE
	// 3. CCARD
	// 4. EMONE

	switch lookup {
	case 1:
		paymentType = "CASH"
	case 2:
		paymentType = "CHEQUE"
	case 3:
		paymentType = "CCARD"
	case 4:
		paymentType = "EMONYE"

	default:
		//log.Panicln("No payment type for the supplied LOOKUP")
		log.Println("Defaulting to INVOICE payment")
		paymentType = "INVOICE"

	}

	return paymentType
}
