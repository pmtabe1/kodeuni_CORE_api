package go_client

import (
	"fmt"
	"testing"

	"github.com/paulmsegeya/pos/utils"
	"github.com/stretchr/testify/require"
)

func TestPostUpdateSetup(t *testing.T) {

	headers := make(map[string]interface{}, 0)
	headers["Content-Type"] = "application/json"
	payload := `{
		"ID": 112318380,
		"CreatedAt": "0001-01-01T00:00:00Z",
		"UpdatedAt": "2021-10-15T23:43:02.967+03:00",
		"DeletedAt": null,
		"SyncToken": 0,
		"QRCode": null,
		"BarCode": null,
		"Attachment": null,
		"Version": 0,
		"Enabled": 1,
		"Locale": "en",
		"Name": "HHH",
		"Type": "",
		"Stage": "updated",
		"Maker": "",
		"Checker": "",
		"Description": "",
		"Status": 0,
		"WorkflowLevels": 0,
		"CompanyID": 26,
		"CertName": "duxte",
		"CertFile": null,
		"CertSerial": "32213952251598420917675373467124718226",
		"CertFilePath": "/Users/touchbar/ssl/app/etc/certs/duxte.pfx",
		"CertTin": "112318380",
		"CertKey": "10TZ100666",
		"CertPass": "uxte2!",
		"RegID": "1121212",
		"Username": "",
		"Password": "",
		"GrantType": "password",
		"GC": 0,
		"RCTNUM_IS_GC": 0,
		"ClientName": "",
		"ZNUM_IS_YYYYMMDD": "",
		"RCTVNUM_IS_GC_PLUS_RCTVCODE": "",
		"EFDSerial": "",
		"DC": 0,
		"ReceiptEODCounter": 0,
		"RealmID": "",
		"AuthCode": "",
		"ClientID": "",
		"ClientSecret": "jhgghghghg",
		"RefreshToken": "vvvvvvvv",
		"AccessToken": "vvvvvvv"
	}`

	uri := "http://localhost:9001/tra/api/v1/setup/update/112318380"

	response := New().Post(payload, uri, headers)
	require.Nilf(t, response, "Expected not nill but go %v", response)

}
func TestPost(t *testing.T) {

	headers := make(map[string]interface{}, 0)

	headers["Cert-Serial"] = "MzIyMTM5NTIyNTE1OTg0MjA5MTc2NzUzNzM0NjcxMjQ3MTgyMjY="
	headers["Client"] = "WEBAPI"
	headers["Content-Type"] = "application/xml"

	payload := `<EFDMS><REGDATA><TIN>112318380</TIN><CERTKEY>10TZ100666</CERTKEY></REGDATA><EFDMSSIGNATURE>V6hqmr-uHrvJ7JesEyF0IeQjEzA35v1oiO3MHyVcycC4-hS22eVPdpTDZ9F8CzPkLqfKlSpo4UffdZVihMhMpoo1wBSIoS3rG_187LzQxGJzaL_62yLG2Nuddlh1mumcg9bJgy7GnHcZJRuWLApgrO5n4qa1JvzMicG2-jwv4TBFNAOnLd-1LSinsmcw4QOgkjyAC0ZV6-42mV_wXJjmRz4uv9b8PpeTLfZeFC20PvI9V6aaOo1VIzaK1daPc5YxavoXcAJaV_-xMj73m7MmzYI-aIrXg_uv0cvsZK2vM2jhC9TnV4-hutV1g1A0xuJmuKcwBNDpFdR9ZIShH06xdg</EFDMSSIGNATURE></EFDMS>`

	url := "https://virtual.tra.go.tz/efdmsRctApi/api/vfdRegReq"
	response := New().Post(payload, url, headers)

	require.Nilf(t, response, "Expected not nill but go %v", response)

}

func TestPostRegistrationFoward(t *testing.T) {

	headers := make(map[string]interface{}, 0)

	headers["Cert-Serial"] = "MzIyMTM5NTIyNTE1OTg0MjA5MTc2NzUzNzM0NjcxMjQ3MTgyMjY="
	headers["Client"] = "WEBAPI"
	headers["Content-Type"] = "application/xml"

	temp := `{"header":{"Cert-Serial":"%v","Client":"%v","Content-Type":"%v"},"payload":"%v"}`

	// payload:=`<?xml version='1.0' encoding='UTF-8'?>
	// 		<EFDMS>
	// 			<REGDATA>
	// 				<TIN>112318380</TIN>
	// 				<CERTKEY>10TZ100666</CERTKEY>
	// 			</REGDATA>
	// 			<EFDMSSIGNATURE>
	// 				V6hqmr+uHrvJ7JesEyF0IeQjEzA35v1oiO3MHyVcycC4-hS22eVPdpTDZ9F8CzPkLqfKlSpo4UffdZVihMhMpoo1wBSIoS3rG/187LzQxGJzaL_62yLG2Nuddlh1mumcg9bJgy7GnHcZJRuWLApgrO5n4qa1JvzMicG2/jwv4TBFNAOnLd-1LSinsmcw4QOgkjyAC0ZV6/42mV_wXJjmRz4uv9b8PpeTLfZeFC20PvI9V6aaOo1VIzaK1daPc5YxavoXcAJaV/+xMj73m7MmzYI+aIrXg/uv0cvsZK2vM2jhC9TnV4+hutV1g1A0xuJmuKcwBNDpFdR9ZIShH06xdg==
	// 			</EFDMSSIGNATURE>
	// 		</EFDMS>`

	payload := `<?xml version='1.0' encoding='UTF-8'?>
	<EFDMS>
		<REGDATA>
			<TIN>112318380</TIN>
			<CERTKEY>10TZ100666</CERTKEY>
		</REGDATA>
		<EFDMSSIGNATURE>
			V6hqmr+uHrvJ7JesEyF0IeQjEzA35v1oiO3MHyVcycC4+hS22eVPdpTDZ9F8CzPkLqfKlSpo4UffdZVihMhMpoo1wBSIoS3rG/187LzQxGJzaL/62yLG2Nuddlh1mumcg9bJgy7GnHcZJRuWLApgrO5n4qa1JvzMicG2+jwv4TBFNAOnLd+1LSinsmcw4QOgkjyAC0ZV6+42mV/wXJjmRz4uv9b8PpeTLfZeFC20PvI9V6aaOo1VIzaK1daPc5YxavoXcAJaV/+xMj73m7MmzYI+aIrXg/uv0cvsZK2vM2jhC9TnV4+hutV1g1A0xuJmuKcwBNDpFdR9ZIShH06xdg==
		</EFDMSSIGNATURE>
	</EFDMS>`
	url := "http://localhost:80/tra/regfoward.php"

	payloadFoward := fmt.Sprintf(temp, headers["Cert-Serial"], headers["Client"], headers["Content-Type"], utils.StringToBase64Encoded(payload))
	response := New().Post(payloadFoward, url, headers)
	require.Nilf(t, response, "Expected not nill but go %v", response)

}
