package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	BusinessPartnerFullName       *string `json:"BusinessPartnerFullName"`
	Industry                      *string `json:"Industry"`
	LegalEntityRegistration       *string `json:"LegalEntityRegistration"`
	OrganizationBPName1           *string `json:"OrganizationBPName1"`
	OrganizationBPName2           *string `json:"OrganizationBPName2"`
	OrganizationBPName3           *string `json:"OrganizationBPName3"`
	OrganizationBPName4           *string `json:"OrganizationBPName4"`
	BPTag1                        *string `json:"BPTag1"`
	BPTag2                        *string `json:"BPTag2"`
	BPTag3                        *string `json:"BPTag3"`
	BPTag4                        *string `json:"BPTag4"`
	OrganizationFoundationDate    *string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string `json:"GroupBusinessPartnerName2"`
	BusinessPartnerIDByExtSystem  *string `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
}

type RoleUpdates struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityStartDate   string `json:"ValidityStartDate"`
}

type FinInstUpdates struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstCode               *string `json:"FinInstCode"`
	FinInstBranchCode         *string `json:"FinInstBranchCode"`
	FinInstFullCode           *string `json:"FinInstFullCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         *string `json:"FinInstControlKey"`
	FinInstAccountName        *string `json:"FinInstAccountName"`
	FinInstAccount            *string `json:"FinInstAccount"`
	HouseBank                 *string `json:"HouseBank"`
	HouseBankAccount          *string `json:"HouseBankAccount"`
}

type AccountingUpdates struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
}
