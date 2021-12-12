package mappers_utils

import "log"

func GetMappedTaxCode(taxRateLookup float64) (taxCode string) {

	// Identifier of the Tax rate
	// A= Standard Rate (18%) (For VAT items)
	// B= Special Rate
	// C= Zero rated (For Non-VAT items) D= Special Relief
	// E= Exempt

	switch taxRateLookup {
	case float64(18):
		taxCode = "A"
	case float64(1):
		taxCode = "B"
	case float64(2):
		taxCode = "B"
	case float64(3):
		taxCode = "E"
	case float64(0):
		taxCode = "C"

	default:
		log.Panicln("No payment type for the supplied LOOKUP")
	}

	return taxCode
}
