package fundconnext

import (
	"encoding/json"
	"fmt"
)

type RetrievalIndividualCustomerProfileAndAccount struct {
	IdentificationCardType         string              `json:"identificationCardType"`
	PassportCountry                *string             `json:"passportCountry"`
	CardNumber                     string              `json:"cardNumber"`
	CardExpiryDate                 string              `json:"cardExpiryDate"`
	AccompanyingDocument           *string             `json:"accompanyingDocument"`
	Title                          string              `json:"title"`
	TitleOther                     *string             `json:"titleOther"`
	EnFirstName                    string              `json:"enFirstName"`
	EnLastName                     string              `json:"EnLastName"`
	ThFirstName                    string              `json:"thFirstName"`
	ThLastName                     string              `json:"thLastName"`
	BirthDate                      string              `json:"birthDate"`
	Nationality                    string              `json:"nationality"`
	MobileNumber                   string              `json:"mobileNumber"`
	Email                          *string             `json:"email"`
	Phone                          *string             `json:"phone"`
	Fax                            *string             `json:"fax"`
	MaritalStatus                  string              `json:"maritalStatus"`
	Spouse                         interface{}         `json:"spouse"`
	OccupationId                   uint64              `json:"occupationId"`
	OccupationOther                *string             `json:"occupationOther"`
	BusinessTypeId                 *uint64             `json:"businessTypeId"`
	BusinessTypeOther              *string             `json:"businessTypeOther"`
	MonthlyIncomeLevel             string              `json:"monthlyIncomeLevel"`
	AssetValue                     *string             `json:"assetValue"`
	IncomeSource                   string              `json:"incomeSource"`
	IncomeSourceOther              *string             `json:"incomeSourceOther"`
	IdentificationDocument         Address             `json:"identificationDocument"`
	CurrentAddressSameAsFlag       *string             `json:"currentAddressSameAsFlag"`
	Current                        *Address            `json:"current"`
	CompanyName                    *string             `json:"companyName"`
	Work                           *Address            `json:"work"`
	WorkPosition                   *string             `json:"workPosition"`
	RelatedPoliticalPerson         *bool               `json:"relatedPoliticalPerson"`
	PoliticalRelatedPersonPosition *string             `json:"politicalRelatedPersonPosition"`
	CanAcceptFxRisk                bool                `json:"canAcceptFxRisk"`
	CanAcceptDerivativeInvestment  bool                `json:"canAcceptDerivativeInvestment"`
	SuitabilityRiskLevel           string              `json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate      string              `json:"suitabilityEvaluationDate"`
	Fatca                          bool                `json:"fatca"`
	FatcaDeclarationDate           string              `json:"fatcaDeclarationDate"`
	CddScore                       *string             `json:"cddScore"`
	CddDate                        *string             `json:"cddDate"`
	ReferralPerson                 *string             `json:"referralPerson"`
	ApplicationDate                string              `json:"applicationDate"`
	IncomeSourceCountry            string              `json:"incomeSourceCountry"`
	AcceptBy                       *string             `json:"acceptBy"`
	OpenFundConnextFormFlag        string              `json:"openFundConnextFormFlag"`
	VulnerableFlag                 *bool               `json:"vulnerableFlag"`
	VulnerableDetail               *string             `json:"vulnerableDetail"`
	NdidFlag                       *bool               `json:"ndidFlag"`
	NdidRequestId                  *string             `json:"ndidRequestId"`
	OpenChannel                    *string             `json:"openChannel"`
	SuitabilityForm                *SuitabilityForm    `json:"suitabilityForm"`
	InvestorClass                  *string             `json:"investorClass"`
	Accounts                       *[]RetrievalAccount `json:"accounts"`
	ApprovedDate                   *string             `json:"approvedDate"`
}

type RetrievalAccount struct {
	AccountId                string         `json:"accountId"`
	IcLicense                string         `json:"icLicense"`
	AccountOpenDate          string         `json:"accountOpenDate"`
	MailingAddressSameAsFlag *string        `json:"mailingAddressSameAsFlag"`
	Mailing                  *Address       `json:"mailing"`
	MailingMethod            string         `json:"mailingMethod"`
	InvestmentObjective      string         `json:"investmentObjective"`
	InvestmentObjectiveOther *string        `json:"investmentObjectiveOther"`
	RedemptionBankAccounts   []BankAccount  `json:"redemptionBankAccounts"`
	SubscriptionBankAccounts *[]BankAccount `json:"subscriptionBankAccounts"`
	ApprovedDate             *string        `json:"approvedDate"`
	OpenOmnibusFormFlag      *bool          `json:"openOmnibusFormFlag"`
}

func (f *FundConnext) RetrieveIndividualCustomerProfileAndAccount(cardNumber string) (*RetrievalIndividualCustomerProfileAndAccount, error) {
	url := fmt.Sprintf("/api/customer/individual/investor/profile/v4?cardNumber=%s", cardNumber)

	out, err := f.APICall("GET", url, make([]byte, 0))
	if err != nil {
		f.cfg.Logger.Fatalln("[Func RetrieveIndividualCustomerProfileAndAccount] Error CallToFundConnext ::", err)
		return nil, err
	}

	var results *RetrievalIndividualCustomerProfileAndAccount
	if err := json.Unmarshal(out, &results); err != nil {
		f.cfg.Logger.Fatalln("[Func RetrieveIndividualCustomerProfileAndAccount] Error json.Marshal ::", err)
		return nil, err
	}
	return results, nil
}
