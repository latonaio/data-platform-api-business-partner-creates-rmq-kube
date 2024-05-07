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
		Representative:           	   data.Representative,
		PhoneNumber:           		   data.PhoneNumber,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		Tag1:                          data.Tag1,
		Tag2:                          data.Tag2,
		Tag3:                          data.Tag3,
		Tag4:                          data.Tag4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		LastChangeDate:				   data.LastChangeDate,
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
		LastChangeDate:		 data.LastChangeDate,
	}
}

func ConvertToPersonUpdates(general dpfm_api_input_reader.General, person dpfm_api_input_reader.Person) *PersonUpdates {
	dataGeneral := general
	data := person

	return &PersonUpdates{
			BusinessPartner:			data.BusinessPartner,
			BusinessPartnerType:		data.BusinessPartnerType,
			FirstName:					data.FirstName,
			LastName:					data.LastName,
			FullName:					data.FullName,
			MiddleName:					data.MiddleName,
			NickName:					data.NickName,
			Gender:						data.Gender,
			Language:					data.Language,
			CorrespondenceLanguage:		data.CorrespondenceLanguage,
			BirthDate:					data.BirthDate,
			Nationality:				data.Nationality,
			EmailAddress:				data.EmailAddress,
			MobilePhoneNumber:			data.MobilePhoneNumber,
			ProfileComment:				data.ProfileComment,
			PreferableLocalSubRegion:	data.PreferableLocalSubRegion,
			PreferableLocalRegion:		data.PreferableLocalRegion,
			PreferableCountry:			data.PreferableCountry,
			ActPurpose:					data.ActPurpose,
			LastChangeDate:		 		data.LastChangeDate,
	}
}

func ConvertToAddressUpdates(general dpfm_api_input_reader.General, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataGeneral := general
	data := address

	return &AddressUpdates{
		BusinessPartner:	data.BusinessPartner,
		AddressID:   		data.AddressID,
		PostalCode:  		data.PostalCode,
		LocalSubRegion: 	data.LocalSubRegion,
		LocalRegion: 		data.LocalRegion,
		Country:     		data.Country,
		GlobalRegion: 		data.GlobalRegion,
		TimeZone:	 		data.TimeZone,
		District:    		data.District,
		StreetName:  		data.StreetName,
		CityName:    		data.CityName,
		Building:    		data.Building,
		Floor:       		data.Floor,
		Room:        		data.Room,
		XCoordinate: 		data.XCoordinate,
		YCoordinate: 		data.YCoordinate,
		ZCoordinate: 		data.ZCoordinate,
		Site:		 		data.Site,
		LastChangeDate:		data.LastChangeDate,
	}
}

func ConvertToSNSUpdates(general dpfm_api_input_reader.General, role dpfm_api_input_reader.SNS) *SNSUpdates {
	dataGeneral := general
	data := sNS

	return &SNSUpdates{
		BusinessPartner:     	data.BusinessPartner,
		XURL:  					data.XURL,
		InstagramURL:     		data.InstagramURL,
		TikTokURL:         		data.TikTokURL,
		LastChangeDate:      	data.LastChangeDate,
	}
}

func ConvertToGPSUpdates(general dpfm_api_input_reader.General, role dpfm_api_input_reader.GPS) *GPSUpdates {
	dataGeneral := general
	data := gPS

	return &GPSUpdates{
		BusinessPartner:     	data.BusinessPartner,
		XCoordinate:     		data.XCoordinate,
		YCoordinate:     		data.YCoordinate,
		ZCoordinate:     		data.ZCoordinate,
		LocalSubRegion:  		data.LocalSubRegion,
		LocalRegion:     		data.LocalRegion,
		Country:         		data.Country,
		LastChangeDate:      	data.LastChangeDate,
		LastChangeTime:      	data.LastChangeTime,
	}
}

func ConvertToRankUpdates(general dpfm_api_input_reader.General, role dpfm_api_input_reader.Rank) *RankUpdates {
	dataGeneral := general
	data := rank

	return &RankUpdates{
		BusinessPartner:     dataGeneral.BusinessPartner,
		RankType:			 data.RankType,
		Rank:				 data.Rank,
		ValidityStartDate:   data.ValidityStartDate,
		ValidityEndDate:     data.ValidityEndDate,
		LastChangeDate:		 data.LastChangeDate,
	}
}

func ConvertToPersonOrganizationUpdates(general dpfm_api_input_reader.General, role dpfm_api_input_reader.PersonOrganization) *PersonOrganizationUpdates {
	dataGeneral := general
	data := personOrganization

	return &PersonOrganizationUpdates{
		BusinessPartner:     dataGeneral.BusinessPartner,
	}
}

func ConvertToPersonMobilePhoneAuthUpdates(general dpfm_api_input_reader.General, personMobilePhoneAuth dpfm_api_input_reader.PersonMobilePhoneAuth) *PersonMobilePhoneAuthUpdates {
	dataGeneral := general
	data := personMobilePhoneAuth

	return &PersonMobilePhoneAuthUpdates{
		BusinessPartner:			dataGeneral.BusinessPartner,
		MobilePhoneIsAuthenticated:	data.MobilePhoneIsAuthenticated,
		LastChangeDate:		 		data.LastChangeDate,
	}
}

func ConvertToPersonGoogleAccountAuthUpdates(general dpfm_api_input_reader.General, personGoogleAccountAuth dpfm_api_input_reader.PersonGoogleAccountAuth) *PersonGoogleAccountAuthUpdates {
	dataGeneral := general
	data := personGoogleAccountAuth

	return &PersonGoogleAccountAuthUpdates{
		BusinessPartner:				dataGeneral.BusinessPartner,
		GoogleAccountIsAuthenticated:	data.GoogleAccountIsAuthenticated,
		LastChangeDate:		 			data.LastChangeDate,
	}
}

func ConvertToPersonInstagramAuthUpdates(general dpfm_api_input_reader.General, personInstagramAuth dpfm_api_input_reader.PersonInstagramAuth) *PersonInstagramAuthUpdates {
	dataGeneral := general
	data := personInstagramAuth

	return &PersonInstagramAuthUpdates{
		BusinessPartner:				dataGeneral.BusinessPartner,
		InstagramID:					data.InstagramID,
		InstagramUserName:				data.InstagramUserName,
		InstagramHasProfilePricture:	data.InstagramHasProfilePricture,
		InstagramProfilePrictureURI:	data.InstagramProfilePrictureURI,
		InstagramIsPrivate:				data.InstagramIsPrivate,
		InstagramIsPublished:			data.InstagramIsPublished,
		InstagramIsAuthenticated:		data.InstagramIsAuthenticated,
		LastChangeDate:		 			data.LastChangeDate,
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
		LastChangeDate:		 		data.LastChangeDate,
	}
}

func ConvertToAccountingUpdates(general dpfm_api_input_reader.General, accounting dpfm_api_input_reader.Accounting) *AccountingUpdates {
	dataGeneral := general
	data := accounting

	return &AccountingUpdates{
		BusinessPartner:     dataGeneral.BusinessPartner,
		ChartOfAccounts:     data.ChartOfAccounts,
		FiscalYearVariant:   data.FiscalYearVariant,
		LastChangeDate:		 data.LastChangeDate,
	}
}
