package requests

type PersonGoogleAccountAuth struct {
	BusinessPartner					int		`json:"BusinessPartner"`
	BusinessPartnerType				string	`json:"BusinessPartnerType"`
	EmailAddress					string	`json:"EmailAddress"`
	GoogleID						string	`json:"GoogleID"`
	GoogleAccountIsAuthenticated	bool	`json:"GoogleAccountIsAuthenticated"`
	CreationDate					string	`json:"CreationDate"`
	LastChangeDate					string	`json:"LastChangeDate"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
