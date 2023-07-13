package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
		BusinessPartner:               data.BusinessPartner,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerName:		   data.BusinessPartnerName,
		Industry:                      data.Industry,
		LegalEntityRegistration:       data.LegalEntityRegistration,
		Country:                       data.Country,
		Language:                      data.Language,
		Currency:                      data.Currency,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		BPTag1:                        data.BPTag1,
		BPTag2:                        data.BPTag2,
		BPTag3:                        data.BPTag3,
		BPTag4:                        data.BPTag4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
	}
}

func ConvertToRoleUpdates(general dpfm_api_input_reader.General, role dpfm_api_input_reader.Role) *RoleUpdates {
	dataGeneral := general
	data := role

	return &RoleUpdates{
		BusinessPartner:     dataGeneral.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidityStartDate:   data.ValidityStartDate,
		ValidityEndDate:     data.ValidityEndDate,
	}
}

func ConvertToFinInstUpdates(general dpfm_api_input_reader.General, finInst dpfm_api_input_reader.FinInst) *FinInstUpdates {
	dataGeneral := general
	data := finInst

	return &FinInstUpdates{
		BusinessPartner:          	dataGeneral.BusinessPartner,
		FinInstIdentification:  	data.FinInstIdentification,
		FinInstCountry:				data.FinInstCountry,
		FinInstCode:				data.FinInstCode,
		FinInstBranchCode:			data.FinInstBranchCode,
		FinInstFullCode:			data.FinInstFullCode,
		InternalFinInstCustomerID:	data.InternalFinInstCustomerID,
		InternalFinInstAccountID:	data.InternalFinInstAccountID,
		FinInstControlKey:        	data.FinInstControlKey,
		FinInstAccountName:       	data.FinInstAccountName,
		FinInstAccount:           	data.FinInstAccount,
		HouseBank:                	data.HouseBank,
		HouseBankAccount:         	data.HouseBankAccount,
	}
}

func ConvertToAccountingUpdates(general dpfm_api_input_reader.General, accounting dpfm_api_input_reader.Accounting) *AccountingUpdates {
	dataGeneral := general
	data := accounting

	return &AccountingUpdates{
		BusinessPartner:     dataGeneral.BusinessPartner,
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
	}
}
