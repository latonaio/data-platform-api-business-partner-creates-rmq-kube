package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		CreationTime:                  data.CreationTime,
		Industry:                      data.Industry,
		LegalEntityRegistration:       data.LegalEntityRegistration,
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
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func ConvertToRoleUpdates(role dpfm_api_input_reader.Role) *RoleUpdates {
	data := role

	return &RoleUpdates{
		ValidityStartDate: data.ValidityStartDate,
	}
}

func ConvertToFinInstUpdates(finInst dpfm_api_input_reader.FinInst) *FinInstUpdates {
	data := finInst

	return &FinInstUpdates{
		InternalFinInstAccountID: data.InternalFinInstAccountID,
		FinInstControlKey:        data.FinInstControlKey,
		FinInstAccountName:       data.FinInstAccountName,
		FinInstAccount:           data.FinInstAccount,
		HouseBank:                data.HouseBank,
		HouseBankAccount:         data.HouseBankAccount,
		IsMarkedForDeletion:      data.IsMarkedForDeletion,
	}
}

func ConvertToAccountingUpdates(accounting dpfm_api_input_reader.Accounting) *AccountingUpdates {
	data := accounting

	return &AccountingUpdates{
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
