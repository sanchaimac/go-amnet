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
	SellSettlementDay      int
	SwitchingSettlementDay *int
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
	LagAllocationDay       int
	SettlementHolidayFlag  string
	HealthInsurrance       string
	PreviousFundCode       *string
	InvestorAlert          *string
	ISIN                   *string
	LowBalCondition        *string
}

// SwitchingData ...
type SwitchingData struct {
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
