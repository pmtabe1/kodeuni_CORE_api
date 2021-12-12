package conversion_utils

import (
	"fmt"
	"log"
)

func GetDecimalPositionFromFloat64(data float64, decimalPositions int) (rounded string) {

	//roundedMakeup := fmt.Sprintf("%.v%v", decimalPositions,"f")

	log.Println(rounded)

	//s := fmt.Sprintf("%.2f", 12.3456) // s == "12.35"

	roundedMakeup:="%.2f"

	rounded = fmt.Sprintf(roundedMakeup, data) // s == "12.35"

	return rounded

}
