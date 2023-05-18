package existence_conf

type Returns struct {
	ConnectionKey         string                `json:"connection_key"`
	Result                bool                  `json:"result"`
	RedisKey              string                `json:"redis_key"`
	RuntimeSessionID      string                `json:"runtime_session_id"`
	BusinessPartnerID     *int                  `json:"business_partner"`
	Filepath              string                `json:"filepath"`
	ServiceLabel          string                `json:"service_label"`
	BPGeneralReturn       BPGeneralReturn       `json:"BusinessPartnerGeneral"`
	PlantGeneralReturn    PlantGeneralReturn    `json:"PlantGeneral"`
	CurrencyReturn        CurrencyReturn        `json:"Currency"`
	CountryReturn         CountryReturn         `json:"Country"`
	IndustryReturn        IndustryReturn        `json:"Industry"`
	LanguageReturn        LanguageReturn        `json:"Language"`
	PaymentMethodReturn   PaymentMethodReturn   `json:"PaymentMethod"`
	AddressReturn         AddressReturn         `json:"Address"`
	BatchReturn           BatchReturn           `json:"Batch"`
	StorageLocationReturn StorageLocationReturn `json:"StorageLocation"`
	QuantityUnitReturn    QuantityUnitReturn    `json:"QuantityUnit"`
	ProjectReturn         ProjectReturn         `json:"Project"`
	APISchema             string                `json:"api_schema"`
	Accepter              []string              `json:"accepter"`
	Deleted               bool                  `json:"deleted"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type PlantGeneralReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type CurrencyReturn struct {
	Currency string `json:"Currency"`
}

type CountryReturn struct {
	Country string `json:"Country"`
}

type IndustryReturn struct {
	Industry string `json:"Industry"`
}

type LanguageReturn struct {
	Language string `json:"Language"`
}

type PaymentMethodReturn struct {
	PaymentMethod string `json:"PaymentMethod"`
}

type AddressReturn struct {
	AddressID       int    `json:"AddressID"`
	ValidityEndDate string `json:"ValidityEndDate"`
}

type BatchReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Product         string `json:"Product"`
	Plant           string `json:"Plant"`
	Batch           string `json:"Batch"`
}

type StorageLocationReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
}

type QuantityUnitReturn struct {
	QuantityUnit string `json:"QuantityUnit"`
}

type ProjectReturn struct {
	Project string `json:"Project"`
}
