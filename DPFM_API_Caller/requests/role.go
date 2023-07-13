package requests

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
