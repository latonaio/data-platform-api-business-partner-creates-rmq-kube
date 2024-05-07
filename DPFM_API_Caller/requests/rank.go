package requests

type Rank struct {
	BusinessPartner     int 	`json:"BusinessPartner"`
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate     string	`json:"ValidityEndDate"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string  `json:"LastChangeDate"`
	IsMarkedForDeletion	*bool   `json:"IsMarkedForDeletion"`
}
