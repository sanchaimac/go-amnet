package models

// FundNAV ...
type FundNAV struct {
	AmcCode      *string
	FundCode     *string
	Aum          *float64
	Nav          *float64
	OfferNav     *float64
	BidNav       *float64
	SwitchOutNav *float64
	SwitchInNav  *float64
	NavDate      *string
	SaCode       *string
	TotalUnit    *float64
	TotalAum     *float64
}
