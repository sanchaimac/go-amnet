package data

// DividendNews ...
type DividendNews struct {
	FundTaxID         string
	FundCode          string
	AnnounceDate      string
	AnnounceType      string
	BookClosingDate   *string
	PaymentDate       *string
	DividendRate      *float64
	BegSuspendDateSUB *string
	EndSuspendDateSUB *string
	BegSuspendDateRED *string
	EndSuspendDateRED *string
	BegSuspendDateSWI *string
	EndSuspendDateSWI *string
	BegSuspendDateSWO *string
	EndSuspendDateSWO *string
	AcctPeriodFrom    *string
	AcctPeriodTo      *string
	CancelDate        *string
	UpdatedDate       *string
}

// Dividend ...
type DividendUHID struct {
	FundCode                  string
	BookClosedDate            string
	AMCCode                   string
	AccountID                 *string
	UnitholderID              string
	Unit                      float64
	DividendAmount            float64
	WitholdingTax             float64
	DividendAmountNet         float64
	PaymentType               *string
	BankCode                  string
	BankAccount               *string
	ChequeNo                  *string
	ReinvestFundCode          *string
	ReinvestAMCOrderReference *string
	AgentPayFlag              string
	FundTaxID                 string
	PaymentDate               string
	DividendRate              float64
}
