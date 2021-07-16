package models

// FundConnextType Enum
type FundConnextType uint32

// Enum of data Type (only for fund connect type)
// Ex. Nav, FundProfile, AllottedTransactions
const (
	// GenerateFile state
	Nav FundConnextType = iota
	FundProfile
	FundMapping
)

// ModelType get struct type
func (f FundConnextType) ModelType() interface{} {
	return [...]interface{}{FundNAV{}}[f]
}

func (f FundConnextType) String() string {
	return [...]string{"Nav"}[f]
}

// FundConnextFileType mapping filename and type
var FundConnextFileType = map[string]FundConnextType{
	"NAV": Nav,
}
