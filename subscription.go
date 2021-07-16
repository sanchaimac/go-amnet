package fundconnext

import (
	"encoding/json"
	"fmt"
)

type SubscriptionOrder struct {
	SaOrderReferenceNo    string  `json:"saOrderReferenceNo"`
	TransactionDateTime   string  `json:"transactionDateTime"`
	SACode                string  `json:"saCode"`
	AccountId             string  `json:"accountId"`
	UnitholderId          string  `json:"unitholderId"`
	OverrideRiskProfile   string  `json:"overrideRiskProfile"`
	OverrideFxRisk        string  `json:"overrideFxRisk"`
	FundCode              string  `json:"fundCode"`
	Amount                float32 `json:"amount"`
	EffectiveDate         string  `json:"effectiveDate"`
	PaymentType           string  `json:"paymentType"`
	CreditCardNo          *string `json:"creditCardNo"`
	IssuedBy              *string `json:"issuedBy"`
	BankCode              *string `json:"bankCode"`
	BankAccount           *string `json:"bankAccount"`
	PayChequeNo           *string `json:"payChequeNo"`
	PayChequeDate         *string `json:"payChequeDate"`
	Channel               string  `json:"channel"`
	ICLicense             string  `json:"icLicense"`
	BranchNo              *string `json:"branchNo"`
	LTFCondition          *string `json:"ltfCondition"`
	AutoRedeemFundCode    *string `json:"autoRedeemFundCode"`
	ForceEntry            string  `json:"forceEntry"`
	SettlementBankCode    *string `json:"settlementBankCode"`
	SettlementBankAccount *string `json:"settlementBankAccount"`
	ChqBranch             *string `json:"chqBranch"`
	Status                *string `json:"status"`
	CollateralAccount     *string `json:"collateralAccount"`
}

type SubscriptionResponse struct {
	TransactionId      string `json:"transactionId"`
	SaOrderReferenceNo string `json:"saOrderReferenceNo"`
	UnitholderId       string `json:"unitholderId"`
}

type BasketOrder struct {
	SaOrderReferenceNo    string  `json:"saOrderReferenceNo"`
	UnitholderId          string  `json:"unitholderId"`
	OverrideRiskProfile   string  `json:"overrideRiskProfile"`
	OverrideFxRisk        string  `json:"overrideFxRisk"`
	FundCode              string  `json:"fundCode"`
	Amount                float32 `json:"amount"`
	AutoRedeemFundCode    *string `json:"autoRedeemFundCode"`
	LTFCondition          *string `json:"ltfCondition"`
	SettlementBankCode    *string `json:"settlementBankCode"`
	SettlementBankAccount *string `json:"settlementBankAccount"`
}

type SubscriptionBasketOrder struct {
	TransactionDateTime string        `json:"transactionDateTime"`
	SACode              string        `json:"saCode"`
	AccountId           string        `json:"accountId"`
	EffectiveDate       string        `json:"effectiveDate"`
	PaymentType         string        `json:"paymentType"`
	BankCode            *string       `json:"bankCode"`
	BankAccount         *string       `json:"bankAccount"`
	Channel             string        `json:"channel"`
	ICLicense           string        `json:"icLicense"`
	BranchNo            *string       `json:"branchNo"`
	ForceEntry          string        `json:"forceEntry"`
	BasketOrders        []BasketOrder `json:"basketOrders"`
}

type BasketOrderResponse struct {
	TransactionId      string `json:"transactionId"`
	SaOrderReferenceNo string `json:"saOrderReferenceNo"`
}

type SubscriptionBasketResponse struct {
	BasketTransactionId string                `json:"basketTransactionId"`
	BasketOrders        []BasketOrderResponse `json:"basketOrders"`
}

type TransactionIDResponse struct {
	TransactionId string `json:"transactionId"`
}

func (f *FundConnext) CreateSubscription(subOrder SubscriptionOrder) (*SubscriptionResponse, error) {
	cfg := MakeAPICallerConfig(f)
	url := "/api/subscriptions"
	body, err := json.Marshal(subOrder)
	if err != nil {
		return nil, err
	}
	resp, err := CallFCAPI(f.token, "POST", url, body, cfg)
	if err != nil {
		return nil, err
	}
	var results *SubscriptionResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) CreateSubscriptionV2(subOrder SubscriptionOrder) (*SubscriptionResponse, error) {
	cfg := MakeAPICallerConfig(f)
	url := "/api/subscriptions/v2"
	body, err := json.Marshal(subOrder)
	if err != nil {
		return nil, err
	}
	resp, err := CallFCAPI(f.token, "POST", url, body, cfg)
	if err != nil {
		return nil, err
	}
	var results *SubscriptionResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) CreateSubscriptionBasketOrder(subOrder SubscriptionBasketOrder) (*SubscriptionBasketResponse, error) {
	cfg := MakeAPICallerConfig(f)
	url := "/api/subscriptions/basket"
	body, err := json.Marshal(subOrder)
	if err != nil {
		return nil, err
	}
	resp, err := CallFCAPI(f.token, "POST", url, body, cfg)
	if err != nil {
		return nil, err
	}
	var results *SubscriptionBasketResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) CancelSubscription(transactionId, force string) (*TransactionIDResponse, error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/subscriptions/%s", transactionId)
	body, err := json.Marshal(map[string]string{
		"force": force,
	})
	if err != nil {
		return nil, err
	}
	resp, err := CallFCAPI(f.token, "DELETE", url, body, cfg)
	if err != nil {
		return nil, err
	}
	var results *TransactionIDResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) ApproveSubscription(transactionId, status string) (*TransactionIDResponse, error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/subscriptions/%s", transactionId)
	body, err := json.Marshal(map[string]string{
		"status": status,
	})
	if err != nil {
		return nil, err
	}
	resp, err := CallFCAPI(f.token, "PATCH", url, body, cfg)
	if err != nil {
		return nil, err
	}
	var results *TransactionIDResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
