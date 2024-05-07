package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToGeneralCreates(sdc *dpfm_api_input_reader.SDC) (*General, error) {
	data := sdc.General

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}
	// general.CreationDate = *getSystemDatePtr()
	// general.CreationTime = *getSystemTimePtr()
	// general.LastChangeDate = getSystemDatePtr()
	// general.LastChangeTime = getSystemTimePtr()

	return general, nil
}

func ConvertToRoleCreates(sdc *dpfm_api_input_reader.SDC) (*[]Role, error) {
	roles := make([]Role, 0)

	for _, data := range sdc.General.Role {
		role, err := TypeConverter[*Role](data)
		if err != nil {
			return nil, err
		}

		roles = append(roles, *role)
	}

	return &roles, nil
}

func ConvertToPersonCreates(sdc *dpfm_api_input_reader.SDC) (*[]Person, error) {
	persons := make([]Person, 0)

	for _, data := range sdc.General.Person {
		person, err := TypeConverter[*Person](data)
		if err != nil {
			return nil, err
		}

		persons = append(persons, *person)
	}

	return &persons, nil
}

func ConvertToAddressCreates(sdc *dpfm_api_input_reader.SDC) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range sdc.General.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToSNSCreates(sdc *dpfm_api_input_reader.SDC) (*[]SNS, error) {
	sNSs := make([]SNS, 0)

	for _, data := range sdc.General.SNS {
		sNS, err := TypeConverter[*SNS](data)
		if err != nil {
			return nil, err
		}

		sNSs = append(sNSs, *sNS)
	}

	return &sNSs, nil
}

func ConvertToGPSCreates(sdc *dpfm_api_input_reader.SDC) (*[]GPS, error) {
	gPSs := make([]GPS, 0)

	for _, data := range sdc.General.GPS {
		gPS, err := TypeConverter[*GPS](data)
		if err != nil {
			return nil, err
		}

		gPSs = append(gPSs, *gPS)
	}

	return &gPSs, nil
}

func ConvertToRankCreates(sdc *dpfm_api_input_reader.SDC) (*[]Rank, error) {
	ranks := make([]Rank, 0)

	for _, data := range sdc.General.Rank {
		rank, err := TypeConverter[*Rank](data)
		if err != nil {
			return nil, err
		}

		ranks = append(ranks, *rank)
	}

	return &ranks, nil
}

func ConvertToPersonOrganizationCreates(sdc *dpfm_api_input_reader.SDC) (*[]PersonOrganization, error) {
	personOrganizations := make([]PersonOrganization, 0)

	for _, data := range sdc.General.PersonOrganization {
		personOrganization, err := TypeConverter[*PersonOrganization](data)
		if err != nil {
			return nil, err
		}

		personOrganizations = append(personOrganizations, *personOrganization)
	}

	return &personOrganizations, nil
}

func ConvertToPersonMobilePhoneAuthCreates(sdc *dpfm_api_input_reader.SDC) (*[]PersonMobilePhoneAuth, error) {
	personMobilePhoneAuths := make([]PersonMobilePhoneAuth, 0)

	for _, data := range sdc.General.PersonMobilePhoneAuth {
		personMobilePhoneAuth, err := TypeConverter[*PersonMobilePhoneAuth](data)
		if err != nil {
			return nil, err
		}

		personMobilePhoneAuths = append(personMobilePhoneAuths, *personMobilePhoneAuth)
	}

	return &personMobilePhoneAuths, nil
}

func ConvertToPersonGoogleAccountAuthCreates(sdc *dpfm_api_input_reader.SDC) (*[]PersonGoogleAccountAuth, error) {
	personGoogleAccountAuths := make([]PersonGoogleAccountAuth, 0)

	for _, data := range sdc.General.PersonGoogleAccountAuth {
		personGoogleAccountAuth, err := TypeConverter[*PersonGoogleAccountAuth](data)
		if err != nil {
			return nil, err
		}

		personGoogleAccountAuths = append(personGoogleAccountAuths, *personGoogleAccountAuth)
	}

	return &personGoogleAccountAuths, nil
}

func ConvertToFinInstCreates(sdc *dpfm_api_input_reader.SDC) (*[]FinInst, error) {
	finInsts := make([]FinInst, 0)

	for _, data := range sdc.General.FinInst {
		finInst, err := TypeConverter[*FinInst](data)
		if err != nil {
			return nil, err
		}

		finInsts = append(finInsts, *finInst)
	}

	return &finInsts, nil
}

func ConvertToAccountingCreates(sdc *dpfm_api_input_reader.SDC) (*[]Accounting, error) {
	accountings := make([]Accounting, 0)

	for _, data := range sdc.General.Accounting {
		accounting, err := TypeConverter[*Accounting](data)
		if err != nil {
			return nil, err
		}

		accountings = append(accountings, *accounting)
	}

	return &accountings, nil
}

func ConvertToGeneralUpdates(generalData dpfm_api_input_reader.General) (*General, error) {
	data := generalData

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}

	return general, nil
}

func ConvertToRoleUpdates(roleUpdates *[]dpfm_api_processing_formatter.RoleUpdates) (*[]Role, error) {
	roles := make([]Role, 0)

	for _, data := range *roleUpdates {
		role, err := TypeConverter[*Role](data)
		if err != nil {
			return nil, err
		}

		roles = append(roles, *role)
	}

	return &roles, nil
}

func ConvertToPersonUpdates(personUpdates *[]dpfm_api_processing_formatter.PersonUpdates) (*[]Person, error) {
	persons := make([]Person, 0)

	for _, data := range *personUpdates {
		person, err := TypeConverter[*Person](data)
		if err != nil {
			return nil, err
		}

		persons = append(persons, *person)
	}

	return &persons, nil
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToSNSUpdates(sNSUpdates *[]dpfm_api_processing_formatter.SNSUpdates) (*[]SNS, error) {
	sNSs := make([]SNS, 0)

	for _, data := range *sNSUpdates {
		sNS, err := TypeConverter[*SNS](data)
		if err != nil {
			return nil, err
		}

		sNSs = append(sNSs, *sNS)
	}

	return &sNSs, nil
}

func ConvertToGPSUpdates(gPSUpdates *[]dpfm_api_processing_formatter.GPSUpdates) (*[]GPS, error) {
	gPSs := make([]GPS, 0)

	for _, data := range *gPSUpdates {
		gPS, err := TypeConverter[*GPS](data)
		if err != nil {
			return nil, err
		}

		gPSs = append(gPSs, *gPS)
	}

	return &gPSs, nil
}

func ConvertToRankUpdates(rankUpdates *[]dpfm_api_processing_formatter.RankUpdates) (*[]Rank, error) {
	ranks := make([]Rank, 0)

	for _, data := range *rankUpdates {
		rank, err := TypeConverter[*Rank](data)
		if err != nil {
			return nil, err
		}

		ranks = append(ranks, *rank)
	}

	return &ranks, nil
}

func ConvertToPersonOrganizationUpdates(personOrganizationUpdates *[]dpfm_api_processing_formatter.PersonOrganizationUpdates) (*[]PersonOrganization, error) {
	personOrganizations := make([]PersonOrganization, 0)

	for _, data := range *personOrganizationUpdates {
		personOrganization, err := TypeConverter[*PersonOrganization](data)
		if err != nil {
			return nil, err
		}

		personOrganizations = append(personOrganizations, *personOrganization)
	}

	return &personOrganizations, nil
}

func ConvertToPersonMobilePhoneAuthUpdates(personMobilePhoneAuthUpdates *[]dpfm_api_processing_formatter.PersonMobilePhoneAuthUpdates) (*[]PersonMobilePhoneAuth, error) {
	personMobilePhoneAuths := make([]PersonMobilePhoneAuth, 0)

	for _, data := range *personMobilePhoneAuthUpdates {
		personMobilePhoneAuth, err := TypeConverter[*PersonMobilePhoneAuth](data)
		if err != nil {
			return nil, err
		}

		personMobilePhoneAuths = append(personMobilePhoneAuths, *personMobilePhoneAuth)
	}

	return &personMobilePhoneAuths, nil
}

func ConvertToPersonGoogleAccountAuthUpdates(personGoogleAccountAuthUpdates *[]dpfm_api_processing_formatter.PersonGoogleAccountAuthUpdates) (*[]PersonGoogleAccountAuth, error) {
	personGoogleAccountAuths := make([]PersonGoogleAccountAuth, 0)

	for _, data := range *personGoogleAccountAuthUpdates {
		personGoogleAccountAuth, err := TypeConverter[*PersonGoogleAccountAuth](data)
		if err != nil {
			return nil, err
		}

		personGoogleAccountAuths = append(personGoogleAccountAuths, *personGoogleAccountAuth)
	}

	return &personGoogleAccountAuths, nil
}

func ConvertToFinInstUpdates(finInstUpdates *[]dpfm_api_processing_formatter.FinInstUpdates) (*[]FinInst, error) {
	finInsts := make([]FinInst, 0)

	for _, data := range *finInstUpdates {
		finInst, err := TypeConverter[*FinInst](data)
		if err != nil {
			return nil, err
		}

		finInsts = append(finInsts, *finInst)
	}

	return &finInsts, nil
}

func ConvertToAccountingUpdates(accountingUpdates *[]dpfm_api_processing_formatter.AccountingUpdates) (*[]Accounting, error) {
	accountings := make([]Accounting, 0)

	for _, data := range *accountingUpdates {
		accounting, err := TypeConverter[*Accounting](data)
		if err != nil {
			return nil, err
		}

		accountings = append(accountings, *accounting)
	}

	return &accountings, nil
}

func ConvertToGeneral(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.General = &sub_func_complementer.General{
		BusinessPartner:				input.General.BusinessPartner,
		BusinessPartnerType:			input.General.BusinessPartnerType,
		BusinessPartnerFullName:		input.General.BusinessPartnerFullName,
		BusinessPartnerName:			input.General.BusinessPartnerName,
		Industry:						input.General.Industry,
		LegalEntityRegistration:		input.General.LegalEntityRegistration,
		Country:						input.General.Country,
		Language:						input.General.Language,
		Currency:						input.General.Currency,
		Representative:           	    input.General.Representative,
		PhoneNumber:           		    input.General.PhoneNumber,
		OrganizationBPName1:			input.General.OrganizationBPName1,
		OrganizationBPName2:			input.General.OrganizationBPName2,
		OrganizationBPName3:			input.General.OrganizationBPName3,
		OrganizationBPName4:			input.General.OrganizationBPName4,
		Tag1:							input.General.Tag1,
		Tag2:							input.General.Tag2,
		Tag3:							input.General.Tag3,
		Tag4:							input.General.Tag4,
		OrganizationFoundationDate:		input.General.OrganizationFoundationDate,
		OrganizationLiquidationDate:	input.General.OrganizationLiquidationDate,
		BusinessPartnerBirthplaceName:	input.General.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:		input.General.BusinessPartnerDeathDate,
		GroupBusinessPartnerName1:		input.General.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:		input.General.GroupBusinessPartnerName2,
		AddressID:						input.General.AddressID,
		BusinessPartnerIDByExtSystem:	input.General.BusinessPartnerIDByExtSystem,
		BusinessPartnerIsBlocked:		input.General.BusinessPartnerIsBlocked,
		CertificateAuthorityChain:		input.General.CertificateAuthorityChain,
		UsageControlChain:				input.General.UsageControlChain,
		CreationDate:					input.General.CreationDate,
		LastChangeDate:					input.General.LastChangeDate,
		IsMarkedForDeletion:			input.General.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToRole(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var roles []sub_func_complementer.Role

	roles = append(
		roles,
		sub_func_complementer.Role{
			BusinessPartner:		*input.General.BusinessPartner,
			BusinessPartnerRole:	input.General.Role[0].BusinessPartnerRole,
			ValidityStartDate:		input.General.Role[0].ValidityStartDate,
			ValidityEndDate:		input.General.Role[0].ValidityEndDate,
			CreationDate:			input.General.Role[0].CreationDate,
			LastChangeDate:			input.General.Role[0].LastChangeDate,
			IsMarkedForDeletion:	input.General.Role[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Role = &roles

	return subfuncSDC
}

func ConvertToPerson(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var persons []sub_func_complementer.Person

	persons = append(
		persons,
		sub_func_complementer.Person{
			BusinessPartner:			*input.General.BusinessPartner,
			BusinessPartnerType:		input.General.Person[0].BusinessPartnerType,
			FirstName:					input.General.Person[0].FirstName,
			LastName:					input.General.Person[0].LastName,
			FullName:					input.General.Person[0].FullName,
			MiddleName:					input.General.Person[0].MiddleName,
			NickName:					input.General.Person[0].NickName,
			Gender:						input.General.Person[0].Gender,
			Language:					input.General.Person[0].Language,
			CorrespondenceLanguage:		input.General.Person[0].CorrespondenceLanguage,
			BirthDate:					input.General.Person[0].BirthDate,
			Nationality:				input.General.Person[0].Nationality,
			EmailAddress:				input.General.Person[0].EmailAddress,
			MobilePhoneNumber:			input.General.Person[0].MobilePhoneNumber,
			ProfileComment:				input.General.Person[0].ProfileComment,
			PreferableLocalSubRegion:	input.General.Person[0].PreferableLocalSubRegion,
			PreferableLocalRegion:		input.General.Person[0].PreferableLocalRegion,
			PreferableCountry:			input.General.Person[0].PreferableCountry,
			ActPurpose:					input.General.Person[0].ActPurpose,
			CreationDate:				input.General.Person[0].CreationDate,
			LastChangeDate:				input.General.Person[0].LastChangeDate,
			IsMarkedForDeletion:		input.General.Person[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Person = &persons

	return subfuncSDC
}

func ConvertToAddress(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var addresss []sub_func_complementer.Address

	addresss = append(
		addresss,
		sub_func_complementer.Address{
			BusinessPartner:	*inut.General.BusinessPartner,
			AddressID:   		input.General.Address[0].AddressID,
			PostalCode:  		input.General.Address[0].PostalCode,
			LocalSubRegion: 	input.General.Address[0].LocalSubRegion,
			LocalRegion: 		input.General.Address[0].LocalRegion,
			Country:     		input.General.Address[0].Country,
			GlobalRegion: 		input.General.Address[0].GlobalRegion,
			TimeZone:	 		input.General.Address[0].TimeZone,
			District:    		input.General.Address[0].District,
			StreetName:  		input.General.Address[0].StreetName,
			CityName:    		input.General.Address[0].CityName,
			Building:    		input.General.Address[0].Building,
			Floor:       		input.General.Address[0].Floor,
			Room:        		input.General.Address[0].Room,
			XCoordinate: 		input.General.Address[0].XCoordinate,
			YCoordinate: 		input.General.Address[0].YCoordinate,
			ZCoordinate: 		input.General.Address[0].ZCoordinate,
			Site:		 		input.General.Address[0].Site,
		},
	)

	subfuncSDC.Message.Address = &addresss

	return subfuncSDC
}

func ConvertToSNS(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var sNSs []sub_func_complementer.SNS

	sNSs = append(
		sNSs,
		sub_func_complementer.SNS{
			BusinessPartner:		*inut.General.BusinessPartner,
			BusinessPartnerType:	input.General.SNS[0].BusinessPartnerType,
			XURL:  					input.General.SNS[0].XURL,
			InstagramURL:     		input.General.SNS[0].InstagramURL,
			TikTokURL:         		input.General.SNS[0].TikTokURL,
			CreationDate:        	input.General.SNS[0].CreationDate,
			LastChangeDate:      	input.General.SNS[0].LastChangeDate,
			IsMarkedForDeletion:	input.General.SNS[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.SNS = &sNSs

	return subfuncSDC
}

func ConvertToGPS(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var gPSs []sub_func_complementer.GPS

	gPSs = append(
		gPSs,
		sub_func_complementer.GPS{
			BusinessPartner:		*inut.General.BusinessPartner,
			BusinessPartnerType:	input.General.GPS[0].BusinessPartnerType,
			XCoordinate:     		input.General.GPS[0].XCoordinate,
			YCoordinate:     		input.General.GPS[0].YCoordinate,
			ZCoordinate:     		input.General.GPS[0].ZCoordinate,
			LocalSubRegion:  		input.General.GPS[0].LocalSubRegion,
			LocalRegion:     		input.General.GPS[0].LocalRegion,
			Country:         		input.General.GPS[0].Country,
			CreationDate:        	input.General.GPS[0].CreationDate,
			CreationTime:        	input.General.GPS[0].CreationTime,
			LastChangeDate:      	input.General.GPS[0].LastChangeDate,
			LastChangeTime:      	input.General.GPS[0].LastChangeTime,
			IsMarkedForDeletion:	input.General.GPS[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.GPS = &gPSs

	return subfuncSDC
}

func ConvertToRank(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var ranks []sub_func_complementer.Rank

	ranks = append(
		ranks,
		sub_func_complementer.Rank{
			BusinessPartner:		*inut.General.BusinessPartner,
			RankType:				input.General.Rank[0].RankType,
			Rank:					input.General.Rank[0].Rank,
			ValidityStartDate:		input.General.Rank[0].ValidityStartDate,
			ValidityEndDate:		input.General.Rank[0].ValidityEndDate,
			CreationDate:			input.General.Rank[0].CreationDate,
			LastChangeDate:			input.General.Rank[0].LastChangeDate,
			IsMarkedForDeletion:	input.General.Rank[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Rank = &ranks

	return subfuncSDC
}

func ConvertToPersonOrganization(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var personOrganizations []sub_func_complementer.PersonOrganization

	personOrganizations = append(
		personOrganizations,
		sub_func_complementer.PersonOrganization{
			BusinessPartner:				*inut.General.BusinessPartner,
			BusinessPartnerType:			input.General.PersonOrganization[0].BusinessPartnerType,
			BusinessPartnerOrganization:	input.General.PersonOrganization[0].BusinessPartnerOrganization,
			CreationDate:					input.General.PersonOrganization[0].CreationDate,
			LastChangeDate:					input.General.PersonOrganization[0].LastChangeDate,
			IsMarkedForDeletion:			input.General.PersonOrganization[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.PersonOrganization = &personOrganizations

	return subfuncSDC
}

func ConvertToPersonMobilePhoneAuth(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var personMobilePhoneAuths []sub_func_complementer.PersonMobilePhoneAuth

	personMobilePhoneAuths = append(
		personMobilePhoneAuths,
		sub_func_complementer.PersonMobilePhoneAuth{
			BusinessPartner:				*inut.General.BusinessPartner,
			BusinessPartnerType:			input.General.PersonMobilePhoneAuth[0].BusinessPartnerType,
			MobilePhoneNumber:				input.General.PersonMobilePhoneAuth[0].MobilePhoneNumber,
			MobilePhoneIsAuthenticated:		input.General.PersonMobilePhoneAuth[0].MobilePhoneIsAuthenticated,
			CreationDate:					input.General.PersonMobilePhoneAuth[0].CreationDate,
			LastChangeDate:					input.General.PersonMobilePhoneAuth[0].LastChangeDate,
			IsMarkedForDeletion:			input.General.PersonMobilePhoneAuth[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.PersonMobilePhoneAuth = &personMobilePhoneAuths

	return subfuncSDC
}

func ConvertToPersonGoogleAccountAuth(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var personGoogleAccountAuths []sub_func_complementer.PersonGoogleAccountAuth

	personGoogleAccountAuths = append(
		personGoogleAccountAuths,
		sub_func_complementer.PersonGoogleAccountAuth{
			BusinessPartner:				*inut.General.BusinessPartner,
			BusinessPartnerType:			input.General.PersonGoogleAccountAuth[0].BusinessPartnerType,
			EmailAddress:					input.General.PersonGoogleAccountAuth[0].EmailAddress,
			GoogleID:						input.General.PersonGoogleAccountAuth[0].GoogleID,
			GoogleAccountIsAuthenticated:	input.General.PersonGoogleAccountAuth[0].GoogleAccountIsAuthenticated,
			CreationDate:					input.General.PersonGoogleAccountAuth[0].CreationDate,
			LastChangeDate:					input.General.PersonGoogleAccountAuth[0].LastChangeDate,
			IsMarkedForDeletion:			input.General.PersonGoogleAccountAuth[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.PersonGoogleAccountAuth = &personGoogleAccountAuths

	return subfuncSDC
}

func ConvertToPersonInstagramAuth(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var personInstagramAuths []sub_func_complementer.PersonInstagramAuth

	personInstagramAuths = append(
		personInstagramAuths,
		sub_func_complementer.PersonInstagramAuth{
			BusinessPartner:				*inut.General.BusinessPartner,
			BusinessPartnerType:			input.General.PersonInstagramAuth[0].BusinessPartnerType,
			InstagramID:					input.General.PersonInstagramAuth[0].InstagramID,
			InstagramUserName:				input.General.PersonInstagramAuth[0].InstagramUserName,
			InstagramHasProfilePricture:	input.General.PersonInstagramAuth[0].InstagramHasProfilePricture,
			InstagramProfilePrictureURI:	input.General.PersonInstagramAuth[0].InstagramProfilePrictureURI,
			InstagramIsPrivate:				input.General.PersonInstagramAuth[0].InstagramIsPrivate,
			InstagramIsPublished:			input.General.PersonInstagramAuth[0].InstagramIsPublished,
			InstagramIsAuthenticated:		input.General.PersonInstagramAuth[0].InstagramIsAuthenticated,
			CreationDate:					input.General.PersonInstagramAuth[0].CreationDate,
			LastChangeDate:					input.General.PersonInstagramAuth[0].LastChangeDate,
			IsMarkedForDeletion:			input.General.PersonInstagramAuth[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.PersonInstagramAuth = &personInstagramAuths

	return subfuncSDC
}

func ConvertToFinInst(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var finInsts []sub_func_complementer.FinInst

	finInsts = append(
		finInsts,
		sub_func_complementer.FinInst{
			BusinessPartner:			*inut.General.BusinessPartner,
			FinInstIdentification:		input.General.FinInst[0].FinInstIdentification,
			FinInstCountry:				input.General.FinInst[0].FinInstCountry,
			FinInstCode:				input.General.FinInst[0].FinInstCode,
			FinInstBranchCode:			input.General.FinInst[0].FinInstBranchCode,
			FinInstFullCode:			input.General.FinInst[0].FinInstFullCode,
			FinInstName:				input.General.FinInst[0].FinInstName,
			FinInstBranchName:			input.General.FinInst[0].FinInstBranchName,
			SWIFTCode:					input.General.FinInst[0].SWIFTCode,
			InternalFinInstCustomerID:	input.General.FinInst[0].InternalFinInstCustomerID,
			InternalFinInstAccountID:	input.General.FinInst[0].InternalFinInstAccountID,
			FinInstControlKey:			input.General.FinInst[0].FinInstControlKey,
			FinInstAccountName:			input.General.FinInst[0].FinInstAccountName,
			FinInstAccount:				input.General.FinInst[0].FinInstAccount,
			HouseBank:					input.General.FinInst[0].HouseBank,
			HouseBankAccount:			input.General.FinInst[0].HouseBankAccount,
			CreationDate:				input.General.FinInst[0].CreationDate,
			LastChangeDate:				input.General.FinInst[0].LastChangeDate,
			IsMarkedForDeletion:		input.General.FinInst[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.FinInst = &finInsts

	return subfuncSDC
}

func ConvertToAccounting(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var accountings []sub_func_complementer.Accounting

	accountings = append(
		accountings,
		sub_func_complementer.Accounting{
			BusinessPartner:			*inut.General.BusinessPartner,
			ChartOfAccounts:     		input.General.Accounting[0].ChartOfAccounts,
			FiscalYearVariant:			input.General.Accounting[0].FiscalYearVariant,
			CreationDate:        		input.General.Accounting[0].CreationDate,
			LastChangeDate:      		input.General.Accounting[0].LastChangeDate,
			IsMarkedForDeletion: 		input.General.Accounting[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Accounting = &accountings

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func getSystemDatePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("2006-01-02")
	return &res
}

func getSystemTimePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("15:04:05")
	return &res
}
