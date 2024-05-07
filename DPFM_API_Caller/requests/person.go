package requests

type Person struct {
	BusinessPartner        		int     `json:"BusinessPartner"`
	BusinessPartnerType    		string  `json:"BusinessPartnerType"`
	FirstName              		*string `json:"FirstName"`
	LastName               		*string `json:"LastName"`
	FullName               		*string `json:"FullName"`
	MiddleName             		*string `json:"MiddleName"`
	NickName               		string  `json:"NickName"`
	Gender                 		string  `json:"Gender"`
	Language               		string  `json:"Language"`
	CorrespondenceLanguage 		*string `json:"CorrespondenceLanguage"`
	BirthDate              		*string `json:"BirthDate"`
	Nationality		            string  `json:"Nationality"`
	EmailAddress           		*string `json:"EmailAddress"`
	MobilePhoneNumber      		*string `json:"MobilePhoneNumber"`
	ProfileComment         		*string `json:"ProfileComment"`
	PreferableLocalSubRegion	string  `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion		string  `json:"PreferableLocalRegion"`
	PreferableCountry			string  `json:"PreferableCountry"`
	ActPurpose					string  `json:"ActPurpose"`
	CreationDate           		string  `json:"CreationDate"`
	LastChangeDate         		string  `json:"LastChangeDate"`
	IsMarkedForDeletion    		*bool   `json:"IsMarkedForDeletion"`
}
