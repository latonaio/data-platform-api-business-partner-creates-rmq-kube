package requests

type PersonInstagramAuth struct {
	BusinessPartner					int		`json:"BusinessPartner"`
	BusinessPartnerType				string	`json:"BusinessPartnerType"`
	InstagramID						string	`json:"InstagramID"`
	InstagramUserName        		string  `json:"InstagramUserName"`
	InstagramHasProfilePricture 	bool    `json:"InstagramHasProfilePricture"`
	InstagramProfilePrictureURI		*string `json:"InstagramProfilePrictureURI"`
	InstagramIsPrivate 				bool    `json:"InstagramIsPrivate"`
	InstagramIsPublished 			bool    `json:"InstagramIsPublished"`
	InstagramIsAuthenticated		bool	`json:"InstagramIsAuthenticated"`
	CreationDate					string	`json:"CreationDate"`
	LastChangeDate					string	`json:"LastChangeDate"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
