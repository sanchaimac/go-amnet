package fundconnext

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/machinebox/graphql"
)

type HTTPSetting struct {
	RequestType  string
	URL          string
	Proxy        string
	IsEnableCert bool
	CertPath     string
	Timeout      int64
	Headers      map[string]string
	Parameters   map[string]string
	Body         io.Reader
}

func setProxy(setting *HTTPSetting) (*http.Transport, error) {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	if setting.Proxy != "" {
		path, err := url.Parse(setting.Proxy)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(path)
	}

	return transport, nil
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

////CallToFundConnext ...
//func CallToFundConnext(cfg *APICallerConfig, method, uri string, header Headers, body io.Reader) (int, []byte, error) {
//	client := http.Client{}
//	if cfg.Timeout == nil {
//		client.Timeout = time.Second * 10
//	} else {
//		client.Timeout = *cfg.Timeout
//	}
//
//	req, err := http.NewRequest(method, (urls["stage"] + uri), body)
//
//	if err != nil {
//		return 500, nil, err
//	}
//
//	if method == "GET" || method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE" {
//		req.Header.Add("Content-Type", header.ContentType)
//	}
//
//	req.Header.Add("X-Auth-Token", header.XAuthToken)
//
//	resp, err := client.Do(req)
//
//	if err != nil {
//		return 500, nil, err
//	}
//
//	defer resp.Body.Close()
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//		return 500, nil, err
//	}
//
//	if resp.StatusCode != 200 {
//		var errMsg FCError
//		json.Unmarshal(respBody, &errMsg)
//
//		return resp.StatusCode, nil, &errMsg
//	}
//
//	return resp.StatusCode, respBody, nil
//}

func CallFCAPI(env, token, method, fp string, body interface{}, cfg *APICallerConfig) ([]byte, error) {
	cfg.Logger.Infoln("[Funconnext:CallFCAPI] Called...")
	client := http.Client{}
	if cfg.Timeout == nil {
		client.Timeout = time.Second * 10
	} else {
		client.Timeout = *cfg.Timeout
	}
	if cfg.Proxy != "" {
		transport, err := setProxy(&HTTPSetting{
			Proxy: cfg.Proxy,
		})
		if err != nil {
			cfg.Logger.Error("[Func CallFundconnextAPI] Error proxy failed", err)
			return nil, err
		}
		client.Transport = transport
		cfg.Logger.Info("[Func CallFundconnextAPI] I am on proxy")
	}
	url := fmt.Sprintf("%s%s", env, fp)
	var reqReader io.Reader
	switch body.(type) {
	case []byte:
		reqReader = bytes.NewBuffer(body.([]byte))
	case io.Reader:
		reqReader = body.(io.Reader)
	case nil:
		reqReader = nil
	default:
		return nil, errors.New("invalid body type")
	}

	req, err := http.NewRequest(method, url, reqReader)
	if err != nil {
		cfg.Logger.Error("[Func CallFundconnextAPI] Error create new request failed", err)
		return nil, err
	}

	contentType := "application/json"
	if cfg.ContentType != "" {
		contentType = cfg.ContentType
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		cfg.Logger.Error("[Func CallFundconnextAPI] Error request failed", err)
		if strings.Contains(err.Error(), "EOF") {
			retire := 3
			for i := 1; i <= retire; i++ {
				resp, err = client.Do(req)
				if err == nil {
					break
				}
				time.Sleep(time.Duration(i) * time.Second)
			}
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	defer resp.Body.Close()

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		cfg.Logger.Warning("[Func CallFundconnextAPI] Error ioutil.ReadAll Request Body", err)
	}
	reqHeadersBytes, err := json.Marshal(req.Header)
	if err != nil {
		cfg.Logger.Warning("[Func CallFundconnextAPI] Error Marshal Request Header", err)
	}
	cfg.Logger.Debugf("[Funconnect:CallFundconnextAPI] Debug call Method: %s\n, Url: %s\n, Header: %s\n, Body: %s", method, url, string(reqHeadersBytes), string(reqBody))
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cfg.Logger.Info("[Func CallFundconnextAPI] Response: ", string(respBody), " StatusCode: ", resp.StatusCode)
	if resp.StatusCode != 200 {
		var errMsg FCError
		if err := json.Unmarshal(respBody, &errMsg); err != nil {
			cfg.Logger.Error("[Func CallFundconnextAPI] Unmarshal Error: ", err)
			return nil, err
		}
		cfg.Logger.Error("[Func CallFundconnextAPI] Error request failed: ", errMsg)
		return nil, &errMsg
	}
	return respBody, nil
}

func CallAmnetApi(ctx context.Context, req *graphql.Request, env, token, endPoint string, cfg *APICallerConfig) (responseData interface{}, err error) {
	u, err := url.Parse(env)
	if err != nil {
		cfg.Logger.Infoln("[Funconnext:CallAmnetApi] Parse url error: ", err)
		return nil, err
	}
	u.Path = path.Join(endPoint)
	client := graphql.NewClient(u.Path)

	contentType := "application/json"
	if cfg.ContentType != "" {
		contentType = cfg.ContentType
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Add("X-Auth-Token", token)
	// Perform the request
	if err := client.Run(ctx, req, responseData); err != nil {
		cfg.Logger.Infoln("[Funconnext:CallAmnetApi] Graphql client request error: ", err)
		return nil, err
	}

	return
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
