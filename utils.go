package fundconnext

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var urls map[string]string = map[string]string{
	"demo":       "https://demo.fundconnext.com",
	"stage":      "https://stage.fundconnext.com",
	"production": "https://www.fundconnext.com",
}

type APICallerConfig struct {
	Timeout time.Duration
}

//Headers ...
type Headers struct {
	ContentType string
	XAuthToken  string
}

func ScanRowToStruct(data string, r interface{}) {
	elem := reflect.ValueOf(r).Elem()
	num := elem.Type().NumField()
	s := strings.Split(data, "|")
	fmt.Println(s)
	for i := 0; i < num; i++ {
		// := elem.Field(i)
		t := elem.Type().Field(i)
		if t.Type.Kind() == reflect.Ptr {
			//
		} else {
			fmt.Println(t.Type.Kind())
		}
	}
}

//CallToFundConnext ...
func CallToFundConnext(cfg *APICallerConfig, method, uri string, header Headers, body io.Reader) (int, []byte, error) {
	client := http.Client{
		Timeout: cfg.Timeout,
	}

	req, err := http.NewRequest(method, (urls["stage"] + uri), body)

	if err != nil {
		return 500, nil, err
	}

	if method == "GET" || method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE" {
		req.Header.Add("Content-Type", header.ContentType)
	}

	req.Header.Add("X-Auth-Token", header.XAuthToken)

	resp, err := client.Do(req)

	if err != nil {
		return 500, nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 500, nil, err
	}

	if resp.StatusCode != 200 {
		var errMsg FCError
		json.Unmarshal(respBody, &errMsg)

		return resp.StatusCode, nil, &errMsg
	}

	return resp.StatusCode, respBody, nil
}

func CallFCAPI(token, method, fp string, body []byte, cfg *APICallerConfig) ([]byte, error) {
	client := http.Client{
		Timeout: cfg.Timeout,
	}

	req, err := http.NewRequest(method, (urls["stage"] + fp), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE" {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// log.Println("resp.StatusCode ::", resp.StatusCode)
	// log.Println("resp.Body ::", string(respBody))
	// log.Println("resp.Header ::", resp.Header)

	if resp.StatusCode != 200 {
		var errMsg FCError
		json.Unmarshal(respBody, &errMsg)
		return nil, &errMsg
	}
	return respBody, nil
}

//CallFCAPIV2 ...
func CallFCAPIV2(token, method, fp string, body io.Reader, cfg *APICallerConfig, contentType string) ([]byte, error) {
	client := http.Client{
		Timeout: cfg.Timeout,
	}

	req, err := http.NewRequest(method, (urls["stage"] + fp), body)

	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE" {
		req.Header.Add("Content-Type", contentType)
		log.Println("Content-Type :: ", contentType)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	log.Println("resp.StatusCode ::", resp.StatusCode)
	log.Println("resp.Body ::", string(respBody))
	log.Println("resp.Header ::", resp.Header)

	if resp.StatusCode != 200 {
		var errMsg FCError
		json.Unmarshal(respBody, &errMsg)

		return nil, &errMsg
	}

	return respBody, nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CheckStruct print struct value
func CheckStruct(structValue interface{}) {

	v := reflect.ValueOf(structValue)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	reflectType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Ptr {
			if v.Field(i).IsNil() {
				fmt.Printf("Field: %s\tValue: %v\n", reflectType.Field(i).Name, nil)
			} else {
				fmt.Printf("Field: %s\tValue: %v\n", reflectType.Field(i).Name, v.Field(i).Elem())
			}
		} else {
			fmt.Printf("Field: %s\tValue: %v\n", reflectType.Field(i).Name, v.Field(i).Interface())

		}
	}
}

// StringToInt ...
func StringToInt(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		log.Println(err)
		return 0
	}
	return number
}

// MarshalCombo Copy value of similar struct from A -> B
func MarshalCombo(valueStruct interface{}, emptyStruct interface{}) error {
	arrayBytes, marsahlErr := json.Marshal(valueStruct)
	if marsahlErr != nil {
		return errors.New("Fail to perform Marshal")
	}
	unmarsahlErr := json.Unmarshal(arrayBytes, emptyStruct)
	if unmarsahlErr != nil {
		return errors.New("Fail to perform Unmarshal")
	}
	return nil
}
