package existence_conf

type Returns struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	Filepath          string          `json:"filepath"`
	ServiceLabel      string          `json:"service_label"`
	BPGeneralReturn   BPGeneralReturn `json:"BusinessPartnerGeneral"`
	IndustryReturn    IndustryReturn  `json:"Industry"`
	CountryReturn     *CountryReturn  `json:"SupplyChainRelationshipGeneral,omitempty"`
	LanguageReturn    *LanguageReturn `json:"Language,omitempty"`
	CurrencyReturn    CurrencyReturn  `json:"Currency"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	Deleted           bool            `json:"deleted"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type IndustryReturn struct {
	Industry string `json:"Industry"`
}

type CountryReturn struct {
	Country      string `json:"Country"`
	GlobalRegion string `json:"GlobalRegion"`
}

type LanguageReturn struct {
	Language string `json:"Language"`
}

type CurrencyReturn struct {
	Currency string `json:"Currency"`
}
