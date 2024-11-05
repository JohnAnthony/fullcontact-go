package fullcontact

type Household struct {
	HomeInfo     HomeInfo     `json:"homeInfo"`
	Presence     Presence     `json:"presence"`
	Finance      Finance      `json:"finance"`
	LocationInfo LocationInfo `json:"locationInfo"`
	FamilyInfo   FamilyInfo   `json:"familyInfo"`
}

type HomeInfo struct {
	HomeValueEstimate   int    `json:"homeValueEstimate"`
	LoanToValueEstimate int    `json:"loanToValueEstimate"`
	YearsInHome         int    `json:"yearsInHome"`
	DwellingType        string `json:"dwellingType"`
}

type Presence struct {
	PresenceOfChildren        string `json:"presenceOfChildren"`
	MultigenerationalResident string `json:"multigenerationalResident"`
}

type Finance struct {
	HouseholdIncomeEstimate        string `json:"householdIncomeEstimate"`
	DiscretionarySpendingIncome    string `json:"discretionarySpendingIncome"`
	FirstMortgageAmountInThousands string `json:"firstMortgageAmountInThousands"`
	HomeMarketValueTaxRecord       string `json:"homeMarketValueTaxRecord"`
	ShortTermLiability             string `json:"shortTermLiability"`
	NetWorthRange                  string `json:"netWorthRange"`
	WealthResources                string `json:"wealthResources"`
	PaymentMethodCreditCard        string `json:"paymentMethodCreditCard"`
}

type LocationInfo struct {
	CarrierRoute             string `json:"carrierRoute"`
	DesignatedMarketArea     string `json:"designatedMarketArea"`
	CoreBasedStatisticalArea string `json:"coreBasedStatisticalArea"`
	NielsenCountySize        string `json:"nielsenCountySize"`
	CongressionalDistrict    string `json:"congressionalDistrict"`
	NumericCountyCode        int    `json:"numericCountyCode"`
	SeasonalAddress          bool   `json:"seasonalAddress"`
}

type FamilyInfo struct {
	TotalAdults            int    `json:"totalAdults"`
	NumberOfChildren       string `json:"numberOfChildren"`
	TotalPeopleInHousehold int    `json:"totalPeopleInHousehold"`
}
