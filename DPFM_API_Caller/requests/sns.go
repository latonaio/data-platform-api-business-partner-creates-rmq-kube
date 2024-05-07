package requests

type SNS struct {
	BusinessPartner     int		`json:"BusinessPartner"`
	BusinessPartnerType	string	`json:"BusinessPartnerType"`
	XURL  				*string	`json:"XURL"`
	InstagramURL     	*string	`json:"InstagramURL"`
	TikTokURL         	*string	`json:"TikTokURL"`
	CreationDate        string	`json:"CreationDate"`
	LastChangeDate      string	`json:"LastChangeDate"`
	IsMarkedForDeletion *bool	`json:"IsMarkedForDeletion"`
}
