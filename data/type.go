package data

// FundConnextType Enum
type FundConnextFileType uint32

// Enum of data Type (only for fund connect type)
// Ex. Nav, FundProfile, AllottedTransactions
const (
	// GenerateFile state
	FundMapping FundConnextFileType = iota
	FundProfile
	SwitchingMatrix
	FundHoliday
	TradeCalendar
	Fee
	FundPerformance
	NAV
	UnitholderBalance
	UnitholderBalanceLTF
	AllottedTransaction
	DividendNews
	DividendTransaction
	AccountProfile
	UnitholderMapping
	BankAccountUnitholder
)

// ModelType get struct type
func (f FundConnextFileType) ModelType() interface{} {
	return [...]interface{}{
		FundMappingData{},
		FundProfileData{},
		SwitchingMatrixData{},
		FundHolidayData{},
		TradeCalendarData{},
		FeeData{},
		FundPerformanceData{},
		NAVData{},
		UnitholderBalanceData{},
		UnitholderBalanceLTFData{},
		AllottedTransactionData{},
		DividendNewsData{},
		DividendUHIDConfirmation{},
		AccountProfileData{},
		UnitholderMappingData{},
		BankAccountUnitholderData{},
	}[f]
}

func (f FundConnextFileType) String() string {
	return [...]string{
		"FundMapping",
		"FundProfile",
		"SwitchingMatrix",
		"FundHoliday",
		"TradeCalendar",
		"Fee",
		"FundPerformance",
		"Nav",
		"UnitholderBalance",
		"LtfBalance",
		"AllottedTransactions",
		"DividendNews",
		"DividendTransactions",
		"AccountProfile",
		"UnitholderMapping",
		"BankAccountUnitholder",
	}[f]
}

type HeaderScheme struct {
	AsOfDate    int
	SACode      int
	TotalRecord int
	Version     int
}

func (f FundConnextFileType) Header() HeaderScheme {
	return [...]HeaderScheme{
		{0, 1, 2, -1},  // FundMapping
		{0, -1, 1, 2},  // FundProfile
		{0, -1, 1, 2},  // SwitchingMatrix
		{0, -1, 1, -1}, // FundHoliday
		{0, -1, 1, 2},  // TradeCalendar
		{0, -1, 1, 2},  // Fee
		{0, -1, 1, -1}, // FundPerformance
		{0, 1, 2, 3},   // Nav
		{0, 1, 2, -1},  // UnitholderBalance
		{0, 1, 2, 3},   // LtfBalance
		{0, 1, 2, 3},   // AllottedTransaction
		{0, 1, 2, 3},   // DividendNews
		{0, 1, 2, 3},   // DividendTransactions
		{0, 1, 2, -1},  // AccountProfile
		{0, 1, 2, -1},  // UnitholderMapping
		{0, 1, 2, -1},  // BankAccountUnitholder
	}[f]
}

// FundConnextFileType mapping filename and type
var FundConnextFileTypeMapping = map[string]FundConnextFileType{
	"FundMapping":           FundMapping,
	"FundProfile":           FundProfile,
	"SwitchingMatrix":       SwitchingMatrix,
	"FundHoliday":           FundHoliday,
	"TradeCalendar":         TradeCalendar,
	"Fee":                   Fee,
	"FundPerformance":       FundPerformance,
	"Nav":                   NAV,
	"UnitholderBalance":     UnitholderBalance,
	"LtfBalance":            UnitholderBalanceLTF,
	"AllottedTransactions":  AllottedTransaction,
	"DividendNews":          DividendNews,
	"DividendTransactions":  DividendTransaction,
	"AccountProfile":        AccountProfile,
	"UnitholderMapping":     UnitholderMapping,
	"BankAccountUnitholder": BankAccountUnitholder,
}
