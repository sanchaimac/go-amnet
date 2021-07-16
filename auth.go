package fundconnext

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	SACode      string `json:"sa_code"`
}

type AuthClaims struct {
	IssuedAt         string `json:"iat"`
	ExpiredAt        string `json:"exp"`
	SellingAgentCode string `json:"sellingAgentCode"`
	SellingAgentId   string `json:"sellingAgentId"`
	Username         string `json:"username"`
}

func Login(username, password string) (*AuthResponse, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	reqBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://stage.fundconnext.com/api/auth", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, MakeInternalError(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 503 {
		return nil, errors.New("service is unavailable")
	}
	if resp.StatusCode != 200 {
		var errMsg FCError
		json.Unmarshal(body, &errMsg)
		return nil, &errMsg
	}

	var authResp AuthResponse
	json.Unmarshal(body, &authResp)
	return &authResp, nil
}
