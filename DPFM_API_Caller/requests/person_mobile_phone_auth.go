package requests

type PersonMobilePhoneAuth struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	BusinessPartnerType			string	`json:"BusinessPartnerType"`
	MobilePhoneNumber			string	`json:"MobilePhoneNumber"`
	MobilePhoneIsAuthenticated	bool	`json:"MobilePhoneIsAuthenticated"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
