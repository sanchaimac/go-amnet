package fundconnext

type AccountDocument struct {
	IdentificationCardType   string        `json:"identificationCardType"`
	PassportCountry          *string       `json:"passportCountry"`
	CardNumber               string        `json:"cardNumber"`
	AccountId                string        `json:"accountId"`
	IcLicense                string        `json:"icLicense"`
	AccountOpenDate          string        `json:"accountOpenDate"`
	MailingAddressSameAsFlag *string       `json:"mailingAddressSameAsFlag"`
	Mailing                  *Address      `json:"mailing"`
	MailingMethod            string        `json:"mailingMethod"`
	InvestmentObjective      string        `json:"investmentObjective"`
	InvestmentObjectiveOther *string       `json:"investmentObjectiveOther"`
	RedemptionBankAccounts   []BankAccount `json:"redemptionBankAccounts"`
	SubscriptionBankAccounts []BankAccount `json:"subscriptionBankAccounts"`
	Approved                 bool          `json:"approved"`
	OpenOmnibusFormFlag      *bool         `json:"openOmnibusFormFlag"`
}

func (f *FundConnext) RetrieveIndividualCustomerProfileAndAccount() error {
	return nil
}
