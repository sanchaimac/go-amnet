package fundconnext

import (
	"encoding/json"
	"fmt"
)

type BalanceInquiry struct {
	UnitholderId  string  `json:"unitholderId"`
	FundCode      string  `json:"fundCode"`
	Unit          float32 `json:"unit"`
	Amount        float32 `json:"amount"`
	RemainUnit    float32 `json:"remainUnit"`
	RemainAmount  float32 `json:"remainAmount"`
	PendingAmount float32 `json:"pendingAmount"`
	PendingUnit   float32 `json:"pendingUnit"`
	AVGCost       float32 `json:"avgCost"`
	NAV           float32 `json:"nav"`
	NAVDate       string  `json:"navDate"`
}

type BalanceInquiryResults struct {
	Result []BalanceInquiry `json:"result"`
}

func (f *FundConnext) BalanceInquiry(accountNo string) (*BalanceInquiryResults, error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/account/balances?accountNo=%s", accountNo)
	out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg)
	if err != nil {
		return nil, err
	}
	var results *BalanceInquiryResults
	json.Unmarshal(out, &results)
	return results, nil
}
