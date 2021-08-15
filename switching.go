package fundconnext

import (
	"encoding/json"
	"fmt"
)

type SwitchingOrder struct {
	SaOrderReferenceNo  string   `json:"saOrderReferenceNo"`
	TransactionDateTime string   `json:"transactionDateTime"`
	SACode              string   `json:"saCode"`
	AccountId           string   `json:"accountId"`
	UnitholderId        string   `json:"unitholderId"`
	FundCode            string   `json:"fundCode"`
	RedemptionType      string   `json:"redemptionType"`
	Amount              *float64 `json:"amount"`
	Unit                *float64 `json:"unit"`
	CounterFundCode     string   `json:"counterFundCode"`
	EffectiveDate       string   `json:"effectiveDate"`
	OverrideRiskProfile string   `json:"overrideRiskProfile"`
	OverrideFxRisk      string   `json:"overrideFxRisk"`
	Channel             string   `json:"channel"`
	ICLicense           string   `json:"icLicense"`
	BranchNo            *string  `json:"branchNo"`
	ForceEntry          string   `json:"forceEntry"`
	SellAllUnitFlag     *string  `json:"sellAllUnitFlag"`
	AutoRedeemFundCode  *string  `json:"autoRedeemFundCode"`
	Status              *string  `json:"status"`
}

type SwitchingOrderResponse struct {
	TransactionId      string `json:"transactionId"`
	SaOrderReferenceNo string `json:"saOrderReferenceNo"`
}

func (f *FundConnext) CreateSwitching(switching SwitchingOrder) (*SwitchingOrderResponse, error) {
	url := "/api/switchings/"
	body, err := json.Marshal(switching)
	if err != nil {
		return nil, err
	}
	resp, err := f.APICall("POST", url, body)
	if err != nil {
		return nil, err
	}
	var results *SwitchingOrderResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) ApproveSwitching(transactionId, status string) (*TransactionIDResponse, error) {
	url := fmt.Sprintf("/api/switchings/%s", transactionId)
	body, err := json.Marshal(map[string]string{
		"status": status,
	})
	if err != nil {
		return nil, err
	}
	resp, err := f.APICall("PATCH", url, body)
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

func (f *FundConnext) CancelSwitching(transactionId, force string) (*TransactionIDResponse, error) {
	url := fmt.Sprintf("/api/switchings/%s", transactionId)
	body, err := json.Marshal(map[string]string{
		"force": force,
	})
	if err != nil {
		return nil, err
	}
	resp, err := f.APICall("DELETE", url, body)
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
