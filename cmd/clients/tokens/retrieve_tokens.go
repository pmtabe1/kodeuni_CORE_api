package tokens

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"


// )

// // import (
// // 	"bitbucket.org/simbiligosi/microservices/rest_microservices/tra_integration_api/cmd/clients/go_client"
// // 	"bitbucket.org/simbiligosi/microservices/rest_microservices/tra_integration_api/constants"
// // )

// func RetrieveTRAToken(registration models.Registration) (tokenResponse token_vfd_post_res_ack.TokenAckResponse) {

// 	// var urlFormEncodedDataValues url.Values

// 	// urlFormEncodedDataValues = url.Values{}
// 	// urlFormEncodedDataValues.Set("Username", registration.Username)
// 	// urlFormEncodedDataValues.Set("Password", registration.Password)
// 	// urlFormEncodedDataValues.Set("grant_type", "password")

// 	// // urlFormEncodedDataValues.Set("grant_type", registration.GrantType)

// 	// // Do the request as the token is not valid anymore

// 	// log.Println("GETTING TOKEN ....")
// 	// log.Printf("Token_Payload :%+v", urlFormEncodedDataValues.Encode())

// 	// url := constants.GetTokenURI(registration.TokenPath)

// 	// client := &http.Client{}
// 	// r, err := http.NewRequest("POST", url, strings.NewReader(urlFormEncodedDataValues.Encode())) // URL-encoded payload
// 	// if err != nil {
// 	// 	log.Panicln(err)
// 	// }
// 	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	// r.Header.Add("Content-Length", strconv.Itoa(len(urlFormEncodedDataValues.Encode())))

// 	// res, err := client.Do(r)
// 	// if err != nil {
// 	// 	log.Panicln(err)
// 	// }
// 	// log.Println(res.Status)
// 	// defer res.Body.Close()
// 	// body, err := ioutil.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// //tokenPayloadJSONStringResponse :=fmt.Sprintf(string(body))

// 	// log.Println("token Response :>>>>" + string(body))
// 	// //var traPayloadModels models.TRAPayloadData

// 	// //	var tokenResponse token_vfd_post_res_ack.TokenAckResponse

// 	// merr := json.Unmarshal(body, &tokenResponse)

// 	// if merr != nil {
// 	// 	log.Panicln("ERROR while unmashalling TOKEN RESPONSE")
// 	// }
// 	// log.Println(string(body))

// 	// if len(tokenResponse.AccessToken) == 0 {
// 	// 	RetrieveTRAToken(registration)
// 	// }

// 	return tokenResponse
// }
