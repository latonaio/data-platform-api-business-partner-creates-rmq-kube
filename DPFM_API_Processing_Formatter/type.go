package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	BusinessPartnerFullName       *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName		      string  `json:"BusinessPartnerName"`
	Industry                      *string `json:"Industry"`
	LegalEntityRegistration       *string `json:"LegalEntityRegistration"`
	Country                       string  `json:"Country"`
	Language                      string  `json:"Language"`
	Currency                      *string `json:"Currency"`
	Representative           	  *string `json:"Representative"`
	PhoneNumber           		  *string `json:"PhoneNumber"`
	OrganizationBPName1           *string `json:"OrganizationBPName1"`
	OrganizationBPName2           *string `json:"OrganizationBPName2"`
	OrganizationBPName3           *string `json:"OrganizationBPName3"`
	OrganizationBPName4           *string `json:"OrganizationBPName4"`
	Tag1                          *string `json:"Tag1"`
	Tag2                          *string `json:"Tag2"`
	Tag3                          *string `json:"Tag3"`
	Tag4                          *string `json:"Tag4"`
	OrganizationFoundationDate    *string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string `json:"GroupBusinessPartnerName2"`
	BusinessPartnerIDByExtSystem  *string `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
	LastChangeDate				  string  `json:"LastChangeDate"`
}

type RoleUpdates struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
}

type PersonUpdates struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	BusinessPartnerType			string	`json:"BusinessPartnerType"`
	FirstName					*string	`json:"FirstName"`
	LastName					*string	`json:"LastName"`
	FullName					*string	`json:"FullName"`
	MiddleName					*string	`json:"MiddleName"`
	NickName					string	`json:"NickName"`
	Gender						string	`json:"Gender"`
	Language					string	`json:"Language"`
	CorrespondenceLanguage		*string	`json:"CorrespondenceLanguage"`
	BirthDate					*string	`json:"BirthDate"`
	Nationality					string	`json:"Nationality"`
	EmailAddress				*string	`json:"EmailAddress"`
	MobilePhoneNumber			*string	`json:"MobilePhoneNumber"`
	ProfileComment				*string	`json:"ProfileComment"`
	PreferableLocalSubRegion	string `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion		string `json:"PreferableLocalRegion"`
	PreferableCountry			string `json:"PreferableCountry"`
	ActPurpose					string  `json:"ActPurpose"`
	LastChangeDate				string	`json:"LastChangeDate"`
}

type AddressUpdates struct {
	BusinessPartner     int 	`json:"BusinessPartner"`
	AddressID   		int     `json:"AddressID"`
	PostalCode  		string 	`json:"PostalCode"`
	LocalSubRegion 		string 	`json:"LocalSubRegion"`
	LocalRegion 		string 	`json:"LocalRegion"`
	Country     		string 	`json:"Country"`
	GlobalRegion   		string 	`json:"GlobalRegion"`
	TimeZone   			string 	`json:"TimeZone"`
	District    		*string `json:"District"`
	StreetName  		*string `json:"StreetName"`
	CityName    		*string `json:"CityName"`
	Building    		*string `json:"Building"`
	Floor       		*int	`json:"Floor"`
	Room        		*int	`json:"Room"`
	XCoordinate 		*float32 `json:"XCoordinate"`
	YCoordinate 		*float32 `json:"YCoordinate"`
	ZCoordinate 		*float32 `json:"ZCoordinate"`
	Site				*int	`json:"Site"`
}

type SNSUpdates struct {
	BusinessPartner     int		`json:"BusinessPartner"`
	XURL  				*string	`json:"XURL"`
	InstagramURL     	*string	`json:"InstagramURL"`
	TikTokURL         	*string	`json:"TikTokURL"`
	LastChangeDate      string	`json:"LastChangeDate"`
}

type GPSUpdates struct {
	BusinessPartner     int		`json:"BusinessPartner"`
	XCoordinate     	float32	`json:"XCoordinate"`
	YCoordinate     	float32	`json:"YCoordinate"`
	ZCoordinate     	float32	`json:"ZCoordinate"`
	LocalSubRegion  	string	`json:"LocalSubRegion"`
	LocalRegion     	string	`json:"LocalRegion"`
	Country         	string	`json:"Country"`
	LastChangeDate      string	`json:"LastChangeDate"`
	LastChangeTime      string	`json:"LastChangeTime"`
}

type RankUpdates struct {
	BusinessPartner     int 	`json:"BusinessPartner"`
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate     string	`json:"ValidityEndDate"`
	LastChangeDate		string  `json:"LastChangeDate"`
}

type PersonOrganizationUpdates struct {
	BusinessPartner        		int     `json:"BusinessPartner"`
}

type PersonMobilePhoneAuthUpdates struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	MobilePhoneIsAuthenticated	bool	`json:"MobilePhoneIsAuthenticated"`
	LastChangeDate				string	`json:"LastChangeDate"`
}

type PersonGoogleAccountAuthUpdates struct {
	BusinessPartner					int		`json:"BusinessPartner"`
	GoogleAccountIsAuthenticated	bool	`json:"GoogleAccountIsAuthenticated"`
	LastChangeDate					string	`json:"LastChangeDate"`
}

type PersonInstagramAuthUpdates struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	InstagramID					string	`json:"InstagramID"`
	InstagramUserName        	string  `json:"InstagramUserName"`
	InstagramHasProfilePricture bool    `json:"InstagramHasProfilePricture"`
	InstagramProfilePrictureURI	*string `json:"InstagramProfilePrictureURI"`
	InstagramIsPrivate 			bool    `json:"InstagramIsPrivate"`
	InstagramIsPublished 		bool    `json:"InstagramIsPublished"`
	InstagramIsAuthenticated	bool	`json:"InstagramIsAuthenticated"`
	LastChangeDate				string	`json:"LastChangeDate"`
}

type FinInstUpdates struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	FinInstCountry            string  `json:"FinInstCountry"`
	FinInstCode               string  `json:"FinInstCode"`
	FinInstBranchCode         string  `json:"FinInstBranchCode"`
	FinInstFullCode           string  `json:"FinInstFullCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         string  `json:"FinInstControlKey"`
	FinInstAccountName        string  `json:"FinInstAccountName"`
	FinInstAccount            string  `json:"FinInstAccount"`
	HouseBank                 *string `json:"HouseBank"`
	HouseBankAccount          *string `json:"HouseBankAccount"`
	LastChangeDate			  string  `json:"LastChangeDate"`
}

type AccountingUpdates struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     string  `json:"ChartOfAccounts"`
	FiscalYearVariant   string  `json:"FiscalYearVariant"`
	LastChangeDate		string	`json:"LastChangeDate"`
}
