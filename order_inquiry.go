package fundconnext

import (
	"encoding/json"
	"fmt"
)

type OrderInquiry struct {
	TransactionID          string  `json:"transactionId"`
	SAOrderReferenceNo     string  `json:"saOrderReferenceNo"`
	OrderType              string  `json:"orderType"`
	AccountID              string  `json:"accountId"`
	UnitholderID           string  `json:"unitholderId"`
	FundCode               string  `json:"fundCode"`
	Unit                   *string `json:"unit"`
	Amount                 float64 `json:"amount"`
	Status                 string  `json:"status"`
	TransactionDateTime    string  `json:"transactionDateTime"`
	TransactionLastUpdated string  `json:"transactionLastUpdated"`
	EffectiveDate          string  `json:"effectiveDate"`
	SettlementDate         string  `json:"settlementDate"`
	AMCOrderReferenceNo    *string `json:"amcOrderReferenceNo"`
	AllottedUnit           float64 `json:"allottedUnit"`
	AllottedAmount         float64 `json:"allottedAmount"`
	AllottedNAV            float64 `json:"allottedNAV"`
	Fee                    float64 `json:"fee"`
	SellAllUnitFlag        *string `json:"sellAllUnitFlag"`
	AllotmentDate          *string `json:"allotmentDate"`
	PaymentType            *string `json:"paymentType"`
	BankCode               *string `json:"bankCode"`
	BankAccount            *string `json:"bankAccount"`
	Channel                *string `json:"channel"`
	ICLicense              string  `json:"icLicense"`
	BranchNo               *string `json:"branchNo"`
	ForceEntry             string  `json:"forceEntry"`
	SettlementBankCode     *string `json:"settlementBankCode"`
	SettlementBankAccount  *string `json:"settlementBankAccount"`
	RejectReason           *string `json:"rejectReason"`
	NAVDate                *string `json:"navDate"`
	CollateralAccount      *string `json:"collateralAccount"`
	AccountType            *string `json:"accountType"`
	RecurringOrderID       *string `json:"recurringOrderId"`
	PaymentStatus          *string `json:"paymentStatus"`
	PaymentProcessingType  *string `json:"paymentProcessingType"`
}

type OrderInquiryResults struct {
	Result []OrderInquiry `json:"result"`
}

func (f *FundConnext) OrderInquiryByAccountNo(accountNo, begEffectiveDate, endEffectiveDate string) (*OrderInquiryResults, error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/account/fundOrders?accountNo=%s&begEffectiveDate=%s&endEffectiveDate=%s", accountNo, begEffectiveDate, endEffectiveDate)
	out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg)
	if err != nil {
		return nil, err
	}
	var results *OrderInquiryResults
	json.Unmarshal(out, &results)
	return results, nil
}

func (f *FundConnext) OrderInquiryByEffectiveDate(effectiveDate string, status *string, channel *string, recuringFlag *string) (*OrderInquiryResults, error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/fundOrders?effectiveDate=%s", effectiveDate)
	if status != nil {
		url += fmt.Sprintf("&status=%s", *status)
	}
	if channel != nil {
		url += fmt.Sprintf("&channel=%s", *channel)
	}
	if recuringFlag != nil {
		url += fmt.Sprintf("&recuringFlag%s", *recuringFlag)
	}
	out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg)
	if err != nil {
		return nil, err
	}
	var results *OrderInquiryResults
	fmt.Println(string(out))
	json.Unmarshal(out, &results)
	return results, nil
}

func (f *FundConnext) OrderInquiryBySAReferenceNo(saRefNo string) (*OrderInquiryResults, error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/fundOrders/saOrderReferenceNo?saOrderReferenceNo=%s", saRefNo)
	out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg)
	if err != nil {
		return nil, err
	}
	var results *OrderInquiryResults
	json.Unmarshal(out, &results)
	return results, nil
}
