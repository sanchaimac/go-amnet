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
		"AllottedTransaction",
		"DividendNews",
		"DividendTransactions",
	}[f]
}

// FundConnextFileType mapping filename and type
var FundConnextFileTypeMapping = map[string]FundConnextFileType{
	"FundMapping":          FundMapping,
	"FundProfile":          FundProfile,
	"SwitchingMatrix":      SwitchingMatrix,
	"FundHoliday":          FundHoliday,
	"TradeCalendar":        TradeCalendar,
	"Fee":                  Fee,
	"FundPerformance":      FundPerformance,
	"Nav":                  NAV,
	"UnitholderBalance":    UnitholderBalance,
	"LtfBalance":           UnitholderBalanceLTF,
	"AllottedTransaction":  AllottedTransaction,
	"DividendNews":         DividendNews,
	"DividendTransactions": DividendTransaction,
}
