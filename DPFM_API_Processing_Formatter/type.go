package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BusinessPartnerFullName       *string `json:"BusinessPartnerFullName"`
	CreationTime                  *string `json:"CreationTime"`
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
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
	GroupBusinessPartnerName1     *string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string `json:"GroupBusinessPartnerName2"`
	BusinessPartnerIDByExtSystem  *string `json:"BusinessPartnerIDByExtSystem"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type RoleUpdates struct {
	ValidityStartDate string `json:"ValidityStartDate"`
}

type FinInstUpdates struct {
	InternalFinInstAccountID *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey        *string `json:"FinInstControlKey"`
	FinInstAccountName       *string `json:"FinInstAccountName"`
	FinInstAccount           *string `json:"FinInstAccount"`
	HouseBank                *string `json:"HouseBank"`
	HouseBankAccount         *string `json:"HouseBankAccount"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type AccountingUpdates struct {
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
