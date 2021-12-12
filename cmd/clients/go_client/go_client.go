package go_client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GoClient struct {
}

type IGoClient interface {
	Post(payload string, url string, headers map[string]interface{}) (response string)
	Patch(payload string, url string, headers map[string]interface{}) (response string)
	Put(payload string, url string, headers map[string]interface{}) (response string)
	Delete(payload *string, url string, headers map[string]interface{}) (response string)
	Get(params map[string]interface{}, url string, headers map[string]interface{}) (response string)
	PostUrlFormEncode(formData url.Values, url string, headers map[string]interface{}) (response string)
}

// func Get() {
// 	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	req.Header.Set("Accept", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	defer resp.Body.Close()

// 	b, err := io.ReadAll(resp.Body)
// 	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	fmt.Println(string(b))
// }

func New() *GoClient {

	return &GoClient{}
}

func (h *GoClient) Post(payload string, url string, headers map[string]interface{}) (response string) {

	log.Println("GO Payload:" + payload)

	req, err := http.NewRequest("POST", url, bytes.NewReader(([]byte(payload))))

	if err != nil {
		log.Panicln(err)
	}

	//		req.Header.Set("Accept", "application/xml")

	for key, element := range headers {
		fmt.Println("Key:", key, "=>", "Element:", element.(string))
		req.Header.Set(key, element.(string))

		// check if it is url form encoded

	}

	// 'Content-type: application/xml',
	// 'Cert-Serial: ' . base64_encode($certSerial),
	// 'Client: WEBAPI',
	// 'Content-Length'.strval($cl)

	log.Println(req.Header)

	client := &http.Client{
		Jar: http.DefaultClient.Jar,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	if headers["Authorization"] != nil {
		if len(headers["Authorization"].(string)) > 0 {

			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				for key, val := range via[3].Header {
					req.Header[key] = val
				}
				return err
			}
		}

	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Panicln(err)
	}

	response = string(b)

	fmt.Println("API CALL RESPONSE: " + response)

	return response
}
func (h *GoClient) PostUrlFormEncode(formData url.Values, url string, headers map[string]interface{}) (response string) {

	log.Printf("GO UrlEncoded Payload: %+v", formData)

	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode())) // URL-encoded payload
	if err != nil {
		log.Panicln(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	if err != nil {
		log.Panicln(err)
	}

	//		req.Header.Set("Accept", "application/xml")

	// for key, element := range headers {
	// 	//fmt.Println("Key:", key, "=>", "Element:", element.(string))
	// 	r.Header.Set(key, element.(string))

	// 	// check if it is url form encoded

	// }

	// 'Content-type: application/xml',
	// 'Cert-Serial: ' . base64_encode($certSerial),
	// 'Client: WEBAPI',
	// 'Content-Length'.strval($cl)

	log.Println(r.Header)

	resp, err := client.Do(r)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Panicln(err)
	}

	response = string(b)

	fmt.Println("API CALL RESPONSE: " + response)

	return response
}

func (h *GoClient) Get(params map[string]interface{}, url string, headers map[string]interface{}) (response string) {

	log.Printf("GO Params:%v", params)

	// Use Params to build proper Uri

	req, err := http.NewRequest("GET", url, nil)

	//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Panicln(err)
	}

	if err != nil {
		log.Panicln(err)
	}

	//		req.Header.Set("Accept", "application/xml")

	for key, element := range headers {
		//fmt.Println("Key:", key, "=>", "Element:", element.(string))
		req.Header.Set(key, element.(string))

	}

	// 'Content-type: application/xml',
	// 'Cert-Serial: ' . base64_encode($certSerial),
	// 'Client: WEBAPI',
	// 'Content-Length'.strval($cl)

	log.Println(req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Panicln(err)
	}

	response = string(b)

	fmt.Println("API CALL RESPONSE: " + response)

	return response
}

func (h *GoClient) Delete(payload *string, params map[string]interface{}, url string, headers map[string]interface{}) (response string) {

	log.Printf("GO Params:%v", params)

	// Use Params to build proper Uri

	//req, err := http.NewRequest("GET", url, nil)

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Panicln(err)
	}

	if err != nil {
		log.Panicln(err)
	}

	//		req.Header.Set("Accept", "application/xml")

	for key, element := range headers {
		//fmt.Println("Key:", key, "=>", "Element:", element.(string))
		req.Header.Set(key, element.(string))

	}

	// 'Content-type: application/xml',
	// 'Cert-Serial: ' . base64_encode($certSerial),
	// 'Client: WEBAPI',
	// 'Content-Length'.strval($cl)

	log.Println(req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Panicln(err)
	}

	response = string(b)

	fmt.Println("API CALL RESPONSE: " + response)

	return response
}

func (h *GoClient) Put(payload string, params map[string]interface{}, url string, headers map[string]interface{}) (response string) {

	log.Printf("GO Params:%v", params)

	// Use Params to build proper Uri

	//req, err := http.NewRequest("GET", url, nil)
	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}

	//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Panicln(err)
	}

	if err != nil {
		log.Panicln(err)
	}

	//		req.Header.Set("Accept", "application/xml")

	for key, element := range headers {
		//fmt.Println("Key:", key, "=>", "Element:", element.(string))
		req.Header.Set(key, element.(string))

	}

	// 'Content-type: application/xml',
	// 'Cert-Serial: ' . base64_encode($certSerial),
	// 'Client: WEBAPI',
	// 'Content-Length'.strval($cl)

	log.Println(req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Panicln(err)
	}

	response = string(b)

	fmt.Println("API CALL RESPONSE: " + response)

	return response
}

func (h *GoClient) Patch(payload string, params map[string]interface{}, url string, headers map[string]interface{}) (response string) {

	log.Printf("GO Params:%v", params)

	// Use Params to build proper Uri

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}

	//req, err := http.NewRequest("GET", url, nil)

	//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Panicln(err)
	}

	if err != nil {
		log.Panicln(err)
	}

	//		req.Header.Set("Accept", "application/xml")

	for key, element := range headers {
		//fmt.Println("Key:", key, "=>", "Element:", element.(string))
		req.Header.Set(key, element.(string))

	}

	// 'Content-type: application/xml',
	// 'Cert-Serial: ' . base64_encode($certSerial),
	// 'Client: WEBAPI',
	// 'Content-Length'.strval($cl)

	log.Println(req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Panicln(err)
	}

	response = string(b)

	fmt.Println("API CALL RESPONSE: " + response)

	return response
}
