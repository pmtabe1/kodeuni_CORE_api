package sage_constants

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const OrderDetailsInsertParams string = "?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?"
const OrderHeadersInsertParams string = "?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?"

func ConvertOrderHeadersInsertParamsToMSSQLParams() (params string) {

	var buffer bytes.Buffer

	dataSlice := strings.Split(OrderHeadersInsertParams, ",")

	if len(dataSlice) == 0 {
		log.Panicln("The Order detail Params slice must never be nil")
	}

	for index, v := range dataSlice {
		log.Println(index)
		indexInc := index + 1

		if indexInc == len(dataSlice) {

			_, _ = buffer.WriteString(fmt.Sprintf("@p%v", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc))))

		} else {
			_, _ = buffer.WriteString(fmt.Sprintf("@p%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc))))

		}
	}
	log.Println(len(dataSlice))
	params = fmt.Sprintf("%v", buffer.String())

	return params
}

func ConvertOrderHeadersInsertParamsToMSSQLParamsWithCast() (params string) {

	var buffer bytes.Buffer

	dataSlice := strings.Split(OrderHeadersInsertParams, ",")

	if len(dataSlice) == 0 {
		log.Panicln("The Order detail Params slice must never be nil")
	}

	for index, v := range dataSlice {
		log.Println(index)
		indexInc := index + 1

		if indexInc == len(dataSlice) {
			_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS char(22)%v", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))

		} else {

			if indexInc == 1 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 2 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 3 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 4 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 5 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 6 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(22)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 7 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 8 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 9 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 10 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 11 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 12 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 13 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 14 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 15 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 16 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 17 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 18 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 19 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 20 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 21 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 22 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 23 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 24 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 25 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 26 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 27 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 28 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 29 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 30 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 31 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 32 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(30)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 33 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 34 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 35 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 36 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(22)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 37 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 38 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 39 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 40 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 41 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 42 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 43 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 44 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 45 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 46 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 47 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 48 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 49 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 50 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 51 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(22)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 52 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 53 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 54 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 55 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 56 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 57 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 58 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(250)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 59 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 60 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 61 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 62 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 63 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 64 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 65 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 66 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 67 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 68 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(2)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 69 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 70 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 71 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(15, 7)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 72 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(15, 7)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 73 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 74 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 75 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 76 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 77 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 78 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 79 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 80 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 81 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 82 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 83 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 84 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 85 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 86 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 87 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 88 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 89 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 90 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 91 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 92 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 93 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 94 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 95 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 96 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 97 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 98 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 99 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 100 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 101 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 102 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 103 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 104 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 105 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 106 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 107 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 108 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 109 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 110 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 111 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 112 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 113 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 114 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 115 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 116 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 117 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 118 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 119 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 120 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 121 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 122 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 123 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 124 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(20)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 125 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 126 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 127 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(22)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 128 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 129 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 130 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 131 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 132 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(5, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 133 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 134 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 135 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 136 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 137 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 138 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 139 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 140 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 141 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 142 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 143 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 144 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 145 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 146 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 147 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 148 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 149 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 150 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 151 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 152 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 153 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 154 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(2)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 155 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 156 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 157 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(15, 7)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 158 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(15, 7)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 159 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 160 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 161 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 162 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 163 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 164 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 165 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 166 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 167 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 168 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   decimal(19, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			// if indexInc == 169 {
			// 	_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v AS   char(22)", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)),")"))
			// }

		}
	}
	log.Println(len(dataSlice))
	params = fmt.Sprintf("%v", buffer.String())

	return params
}

func ConvertOrderDetailsInsertParamsToMSSQLParams() (params string) {

	var buffer bytes.Buffer

	dataSlice := strings.Split(OrderDetailsInsertParams, ",")

	if len(dataSlice) == 0 {
		log.Panicln("The Order detail Params slice must never be nil")
	}

	for index, v := range dataSlice {
		log.Println(index)
		indexInc := index + 1

		if indexInc == len(dataSlice) {
			_, _ = buffer.WriteString(fmt.Sprintf("@p%v", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc))))

		} else {
			_, _ = buffer.WriteString(fmt.Sprintf("@p%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc))))

		}
	}
	log.Println(len(dataSlice))
	params = fmt.Sprintf("%v", buffer.String())

	return params
}

func ConvertOrderDetailsInsertParamsToMSSQLParamsWithCast() (params string) {

	var buffer bytes.Buffer

	dataSlice := strings.Split(OrderDetailsInsertParams, ",")

	if len(dataSlice) == 0 {
		log.Panicln("The Order detail Params slice must never be nil")
	}

	for index, v := range dataSlice {
		log.Println(index)
		indexInc := index + 1

		if indexInc == len(dataSlice) {
			//_, _ = buffer.WriteString(fmt.Sprintf("@p%v", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)),")"))
			_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 0)%v", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))

		} else {
			//_, _ = buffer.WriteString(fmt.Sprintf("@p%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)),")"))

			// if indexInc == 152 {
			// 	_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 0)", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)),")"))
			// }

			if indexInc == 1 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 2 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 3 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 4 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 5 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 6 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 7 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 8 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(24)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 9 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 10 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 11 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 12 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 13 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 14 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 15 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 16 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 17 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 0)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 18 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 19 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 20 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 21 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 22 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 23 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 24 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 25 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 26 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 27 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 28 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 29 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 30 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 31 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 32 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 33 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 34 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 35 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 36 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 37 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 38 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 39 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 40 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 41 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 42 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 43 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 44 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 45 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 46 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 47 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 48 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 49 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 50 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 51 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 52 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 53 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 54 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 55 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 56 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 57 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 58 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 59 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 60 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 61 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(12)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 62 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 63 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 64 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 65 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 66 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 67 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 68 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 69 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 70 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 71 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 72 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 73 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 74 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 75 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 76 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 77 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 78 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 79 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 80 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 81 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 82 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(15, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 83 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(15, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 84 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(15, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 85 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(15, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 86 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(15, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 87 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 88 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 89 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(45)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 90 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 91 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}

			if indexInc == 92 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 93 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(22)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 94 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 95 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(36)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 96 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 97 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(60)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 98 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 99 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 100 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(24)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 101 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(24)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 102 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 103 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 104 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 105 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 106 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 107 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 108 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 109 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(22)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 110 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 111 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 112 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 113 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 114 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 115 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 116 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 117 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 6)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 118 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 119 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 120 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 121 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 122 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(8)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 123 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 124 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 125 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 126 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 127 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 128 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 129 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 130 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 131 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(16)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 132 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(16)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 133 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(16)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 134 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 135 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 136 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 137 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 138 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 139 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(45)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 140 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(45)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 141 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(9, 5)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 142 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 143 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 144 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(16)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 145 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  char(10)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 146 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 3)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 147 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 148 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 149 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  int %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 150 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v  AS  decimal(19, 4)%v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}
			if indexInc == 151 {
				_, _ = buffer.WriteString(fmt.Sprintf("CAST(@p%v   AS  smallint %v,", strings.ReplaceAll(v, "?", strconv.Itoa(indexInc)), ")"))
			}

		}
	}
	log.Println(len(dataSlice))
	params = fmt.Sprintf("%v", buffer.String())

	return params
}
