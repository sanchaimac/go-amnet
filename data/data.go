package data

// FundNAV ...
type FundMappingData struct {
	AmcCode  string
	FundCode string
}

// FundProfileData ...
type FundProfileData struct {
	FundCode               string
	AMCCode                string
	FundNameTH             string
	FundNameEN             string
	FundPolicy             string
	TaxType                *string
	FIFFlag                string
	DividendFlag           string
	RegistrationDate       string
	FundRiskLevel          string
	FXRiskFlag             string
	FATCAAllowFlag         string
	BuyCutOffTime          string
	FstLowbuyVal           *float64
	NxtLowbuyVal           *float64
	SellCutOffTime         string
	LowSellVal             *float64
	LowSellUnit            *float64
	LowBalVal              *float64
	LowBalUnit             *float64
	SellSettlementDay      uint64
	SwitchingSettlementDay *uint64
	SwitchOutFlag          string
	SwitchInFlag           string
	FundClass              *string
	BuyPeriodFlag          string
	SellPeriodFlag         string
	SwitchInPerioldFlag    *string
	SwitchOutPerioldFlag   *string
	BuyPreOrderDay         float64
	SellPreOrderDay        float64
	SwitchPreOrderDay      float64
	AutoRedeemFund         *string
	BegIPODate             *string
	EndIPODate             *string
	PlainComplexFund       string
	DerivativesFlag        string
	LagAllocationDay       uint64
	SettlementHolidayFlag  string
	HealthInsurrance       string
	PreviousFundCode       *string
	InvestorAlert          *string
	ISIN                   *string
	LowBalCondition        *string
}

// SwitchingMatrixData ...
type SwitchingMatrixData struct {
	FundCodeOut         string
	FundCodeIn          string
	SwitchSettlementDay int
	SwitchingType       *string
}

// FundHolidayData ...
type FundHolidayData struct {
	FundCode    string
	HolidayDate string
}

// TradeCalendarData ...
type TradeCalendarData struct {
	FundCode        string
	TransactionCode string
	TradeType       string
	TradeDate       string
}

// FeeData ...
type FeeData struct {
	EffectiveDate string
	FundCode      string
	FeeType       string
	FeeUnit       string
	MaximumFee    *float64
	ActualFee     *float64
	MinimumFee    *float64
	Remark        *string
	Maximum_Value *float64
}

// FundPerformanceData...
type FundPerformanceData struct {
	FundCode   string
	PYTDReturn *float64
	P3MReturn  *float64
	P6MReturn  *float64
	P1YReturn  *float64
	P3YReturn  *float64
	P5YReturn  *float64
	P10YReturn *float64
	PSIReturn  *float64
	P1YSD      *float64
	P3YSD      *float64
	P5YSD      *float64
	P10YSD     *float64
	NAVDate    string
}
