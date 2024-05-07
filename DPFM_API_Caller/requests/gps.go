package requests

type GPS struct {
	BusinessPartner     int		`json:"BusinessPartner"`
	BusinessPartnerType	string	`json:"BusinessPartnerType"`
	XCoordinate     	float32	`json:"XCoordinate"`
	YCoordinate     	float32	`json:"YCoordinate"`
	ZCoordinate     	float32	`json:"ZCoordinate"`
	LocalSubRegion  	string	`json:"LocalSubRegion"`
	LocalRegion     	string	`json:"LocalRegion"`
	Country         	string	`json:"Country"`
	CreationDate        string	`json:"CreationDate"`
	CreationTime        string	`json:"CreationTime"`
	LastChangeDate      string	`json:"LastChangeDate"`
	LastChangeTime      string	`json:"LastChangeTime"`
	IsMarkedForDeletion *bool	`json:"IsMarkedForDeletion"`
}
