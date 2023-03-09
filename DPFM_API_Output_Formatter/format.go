package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-business-partner-creates-rmq-kube/sub_func_complementer"
)

func ConvertToGeneralCreates(subfuncSDC *sub_func_complementer.SDC) *General {
	data := subfuncSDC.Message.General

	general := &General{
		BusinessPartner:               data.BusinessPartner,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerName:           data.BusinessPartnerName,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		Industry:                      data.Industry,
		LegalEntityRegistration:       data.LegalEntityRegistration,
		Country:                       data.Country,
		Language:                      data.Language,
		Currency:                      data.Currency,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
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
		AddressID:                     data.AddressID,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return general
}

func ConvertToGeneralUpdates(generalUpdates *dpfm_api_processing_formatter.GeneralUpdates) *General {
	data := generalUpdates

	general := &General{
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

	return general
}

func ConvertToRoleCreates(subfuncSDC *sub_func_complementer.SDC) *[]Role {
	var role []Role

	for _, data := range subfuncSDC.Message.Role {
		role = append(role, Role{
			BusinessPartner:     data.BusinessPartner,
			BusinessPartnerRole: data.BusinessPartnerRole,
			ValidityEndDate:     data.ValidityEndDate,
			ValidityStartDate:   data.ValidityStartDate,
		})
	}

	return &role
}

func ConvertToRoleUpdates(roleUpdates *dpfm_api_processing_formatter.RoleUpdates) *Role {
	data := roleUpdates

	role := &Role{
		ValidityStartDate: data.ValidityStartDate,
	}

	return role
}

func ConvertToFinInstCreates(subfuncSDC *sub_func_complementer.SDC) *[]FinInst {
	var finInst []FinInst

	for _, data := range subfuncSDC.Message.FinInst {
		finInst = append(finInst, FinInst{
			BusinessPartner:           data.BusinessPartner,
			FinInstIdentification:     data.FinInstIdentification,
			ValidityEndDate:           data.ValidityEndDate,
			ValidityStartDate:         data.ValidityStartDate,
			FinInstCountry:            data.FinInstCountry,
			FinInstCode:               data.FinInstCode,
			FinInstBranchCode:         data.FinInstBranchCode,
			FinInstFullCode:           data.FinInstFullCode,
			FinInstName:               data.FinInstName,
			FinInstBranchName:         data.FinInstBranchName,
			SWIFTCode:                 data.SWIFTCode,
			InternalFinInstCustomerID: data.InternalFinInstCustomerID,
			InternalFinInstAccountID:  data.InternalFinInstAccountID,
			FinInstControlKey:         data.FinInstControlKey,
			FinInstAccountName:        data.FinInstAccountName,
			FinInstAccount:            data.FinInstAccount,
			HouseBank:                 data.HouseBank,
			HouseBankAccount:          data.HouseBankAccount,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}

	return &finInst
}

func ConvertToFinInstUpdates(finInstUpdates *dpfm_api_processing_formatter.FinInstUpdates) *FinInst {
	data := finInstUpdates

	finInst := &FinInst{
		InternalFinInstAccountID: data.InternalFinInstAccountID,
		FinInstControlKey:        data.FinInstControlKey,
		FinInstAccountName:       data.FinInstAccountName,
		FinInstAccount:           data.FinInstAccount,
		HouseBank:                data.HouseBank,
		HouseBankAccount:         data.HouseBankAccount,
		IsMarkedForDeletion:      data.IsMarkedForDeletion,
	}

	return finInst
}

func ConvertToAccountingCreates(subfuncSDC *sub_func_complementer.SDC) *[]Accounting {
	var accounting []Accounting

	for _, data := range subfuncSDC.Message.Accounting {
		accounting = append(accounting, Accounting{
			BusinessPartner:     data.BusinessPartner,
			ChartOfAccounts:     data.ChartOfAccounts,
			FiscalYearVariant:   data.FiscalYearVariant,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &accounting
}

func ConvertToAccountingUpdates(accountingUpdates *dpfm_api_processing_formatter.AccountingUpdates) *Accounting {
	data := accountingUpdates

	accounting := &Accounting{
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return accounting
}
