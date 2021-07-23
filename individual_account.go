package fundconnext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Verified V4
type BankAccount struct {
	BankCode         string  `json:"bankCode"`
	BankBranchCode   *string `json:"bankBranchCode"`
	BankAccountNo    string  `json:"bankAccountNo"`
	Default          bool    `json:"default"`
	FinnetCustomerNo *string `json:"finnetCustomerNo"`
}

type IndividualAccountDocument struct {
	IdentificationCardType   string        `json:"identificationCardType"`
	PassportCountry          *string       `json:"passportCountry"`
	CardNumber               string        `json:"cardNumber"`
	AccountId                string        `json:"accountId"`
	IcLicense                string        `json:"icLicense"`
	AccountOpenDate          string        `json:"accountOpenDate"`
	MailingAddressSameAsFlag *string       `json:"mailingAddressSameAsFlag"`
	Mailing                  *Address      `json:"mailing"`
	MailingMethod            string        `json:"mailingMethod"`
	InvestmentObjective      string        `json:"investmentObjective"`
	InvestmentObjectiveOther *string       `json:"investmentObjectiveOther"`
	RedemptionBankAccounts   []BankAccount `json:"redemptionBankAccounts"`
	SubscriptionBankAccounts []BankAccount `json:"subscriptionBankAccounts"`
	Approved                 bool          `json:"approved"`
	OpenOmnibusFormFlag      *bool         `json:"openOmnibusFormFlag"`
}

type IndividualAccountFile struct {
	IdentificationCardType string  `json:"identificationCardType"`
	PassportCountry        *string `json:"passportCountry"`
	CardNumber             string  `json:"cardNumber"`
	AccountId              string  `json:"accountId"`
	Approved               bool    `json:"approved"`
}

func (f *FundConnext) CreateIndividualAccount(individualAccDoc IndividualAccountDocument) error {
	cfg := MakeAPICallerConfig(f)
	url := "/api/customer/individual/account/v4"

	body, err := json.Marshal(individualAccDoc)

	if err != nil {
		fmt.Println("[Func CreateIndividualAccount] Error json.Marshal ::", err)
		return err
	}

	header := Headers{
		ContentType: "application/json",
		XAuthToken:  f.token,
	}

	// _, err = CallFCAPI(f.token, "POST", url, body, cfg)
	_, _, err = CallToFundConnext(cfg, "POST", url, header, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("[Func CreateIndividualAccount] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

func (f *FundConnext) UpdateIndividualAccount(individualAccDoc IndividualAccountDocument) error {
	cfg := MakeAPICallerConfig(f)
	url := "/api/customer/individual/account/v4"

	body, err := json.Marshal(individualAccDoc)

	if err != nil {
		fmt.Println("[Func UpdateIndividualAccount] Error json.Marshal ::", err)
		return err
	}

	header := Headers{
		ContentType: "application/json",
		XAuthToken:  f.token,
	}

	// _, err = CallFCAPI(f.token, "PUT", url, body, cfg)
	_, _, err = CallToFundConnext(cfg, "PUT", url, header, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("[Func UpdateIndividualAccount] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

// TODO:
// func (f *FundConnext) UploadIndividualAccountFile(fileType string, individualAccountFile IndividualAccountFile) error {
// 	cfg := MakeAPICallerConfig(f)
// 	url := fmt.Sprintf("/api/customer/individual/account/%s/upload", fileType)
// 	body, err := json.Marshal(individualAccountFile)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = CallFCAPI(f.token, "PUT", url, body, cfg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (f *FundConnext) UploadIndividualAccountFile(fileType string, header Headers, body io.Reader) error {
	var err error

	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/customer/individual/account/%s/upload", fileType)
	header.XAuthToken = f.token
	// body, err := json.Marshal(individualCustomerFile)  /api/customer/individual/account/:fileType/upload

	// _, err = CallFCAPIV2(f.token, "POST", url, body, cfg, contentType)
	_, _, err = CallToFundConnext(cfg, "POST", url, header, body)

	if err != nil {
		return err
	}

	return nil
}
