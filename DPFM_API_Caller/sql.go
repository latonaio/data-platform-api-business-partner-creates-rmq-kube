package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var role *[]dpfm_api_output_formatter.Role
	var person *[]dpfm_api_output_formatter.Person
	var address *[]dpfm_api_output_formatter.Address
	var sNS *[]dpfm_api_output_formatter.SNS
	var gPS *[]dpfm_api_output_formatter.GPS
	var rank *[]dpfm_api_output_formatter.Rank
	var personOrganization *[]dpfm_api_output_formatter.PersonOrganization
	var personMobilePhoneAuth *[]dpfm_api_output_formatter.PersonMobilePhoneAuth
	var personGoogleAccountAuth *[]dpfm_api_output_formatter.PersonGoogleAccountAuth
	var personInstagramAuth *[]dpfm_api_output_formatter.PersonInstagramAuth
	var finInst *[]dpfm_api_output_formatter.FinInst
	var accounting *[]dpfm_api_output_formatter.Accounting
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalCreateSql(nil, mtx, input, output, errs, log)
		case "Role":
			role = c.roleCreateSql(nil, mtx, input, output, errs, log)
		case "Person":
			person = c.personCreateSql(nil, mtx, input, output, errs, log)
		case "Address":
			address = c.addressCreateSql(nil, mtx, input, output, errs, log)
		case "SNS":
			sNS = c.sNSCreateSql(nil, mtx, input, output, errs, log)
		case "GPS":
			gPS = c.gPSCreateSql(nil, mtx, input, output, errs, log)
		case "Rank":
			rank = c.rankCreateSql(nil, mtx, input, output, errs, log)
		case "PersonOrganization":
			personOrganization = c.personOrganizationCreateSql(nil, mtx, input, output, errs, log)
		case "PersonMobilePhoneAuth":
			personMobilePhoneAuth = c.personMobilePhoneAuthCreateSql(nil, mtx, input, output, errs, log)
		case "PersonGoogleAccountAuth":
			personGoogleAccountAuth = c.personGoogleAccountAuthCreateSql(nil, mtx, input, output, errs, log)
		case "PersonInstagramAuth":
			personInstagramAuth = c.personInstagramAuthCreateSql(nil, mtx, input, output, errs, log)
		case "FinInst":
			finInst = c.finInstCreateSql(nil, mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:					general,
		Role:						role,
		Person:						person,
		Address:					address,
		SNS:						sNS,
		GPS:						gPS,
		Rank:						rank,
		PersonOrganization:			personOrganization,
		PersonMobilePhoneAuth:		personMobilePhoneAuth,
		PersonGoogleAccountAuth:	personGoogleAccountAuth,
		PersonInstagramAuth:		personInstagramAuth,
		FinInst:					finInst,
		Accounting:					accounting,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var role *[]dpfm_api_output_formatter.Role
	var person *[]dpfm_api_output_formatter.Person
	var address *[]dpfm_api_output_formatter.Address
	var sNS *[]dpfm_api_output_formatter.SNS
	var gPS *[]dpfm_api_output_formatter.GPS
	var rank *[]dpfm_api_output_formatter.Rank
	var personOrganization *[]dpfm_api_output_formatter.PersonOrganization
	var personMobilePhoneAuth *[]dpfm_api_output_formatter.PersonMobilePhoneAuth
	var personGoogleAccountAuth *[]dpfm_api_output_formatter.PersonGoogleAccountAuth
	var personInstagramAuth *[]dpfm_api_output_formatter.PersonInstagramAuth
	var finInst *[]dpfm_api_output_formatter.FinInst
	var accounting *[]dpfm_api_output_formatter.Accounting
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		case "Role":
			role = c.roleUpdateSql(mtx, input, output, errs, log)
		case "Person":
			person = c.personUpdateSql(mtx, input, output, errs, log)
		case "Address":
			address = c.addressUpdateSql(mtx, input, output, errs, log)
		case "SNS":
			sNS = c.sNSUpdateSql(mtx, input, output, errs, log)
		case "GPS":
			gPS = c.gPSUpdateSql(mtx, input, output, errs, log)
		case "Rank":
			rank = c.rankUpdateSql(mtx, input, output, errs, log)
		case "PersonOrganization":
			personOrganization = c.personOrganizationUpdateSql(mtx, input, output, errs, log)
		case "PersonMobilePhoneAuth":
			personMobilePhoneAuth = c.personMobilePhoneAuthUpdateSql(mtx, input, output, errs, log)
		case "PersonGoogleAccountAuth":
			personGoogleAccountAuth = c.personGoogleAccountAuthUpdateSql(mtx, input, output, errs, log)
		case "PersonInstagramAuth":
			personInstagramAuth = c.personInstagramAuthUpdateSql(mtx, input, output, errs, log)
		case "FinInst":
			finInst = c.finInstUpdateSql(mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:					general,
		Role:						role,
		Person:						person,
		Address:					address,
		SNS:						sNS,
		GPS:						gPS,
		Rank:						rank,
		PersonOrganization:			personOrganization,
		PersonMobilePhoneAuth:		personMobilePhoneAuth,
		PersonGoogleAccountAuth:	personGoogleAccountAuth,
		PersonInstagramAuth:		personInstagramAuth,
		FinInst:					finInst,
		Accounting:					accounting,
	}

	return data
}

func (c *DPFMAPICaller) generalCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	generalData := input.General
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "BusinessPartnerGeneral", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "General Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneralCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) roleCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Role {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.Role {
		input.General.Role[i].BusinessPartner = input.General.BusinessPartner
		roleData := input.General.Role[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": roleData, "function": "BusinessPartnerRole", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Role Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToRoleCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Person {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.Person {
		input.General.Person[i].BusinessPartner = input.General.BusinessPartner
		personData := input.General.Person[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personData, "function": "BusinessPartnerPerson", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Person Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.Address {
		input.General.Address[i].BusinessPartner = input.General.BusinessPartner
		addressData := input.General.Address[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "BusinessPartnerAddress", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Address Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) sNSCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SNS {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.SNS {
		input.General.SNS[i].BusinessPartner = input.General.BusinessPartner
		sNSData := input.General.SNS[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": sNSData, "function": "BusinessPartnerSNS", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "SNS Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSNSCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) gPSCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.GPS {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.GPS {
		input.General.GPS[i].BusinessPartner = input.General.BusinessPartner
		gPSData := input.General.GPS[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": gPSData, "function": "BusinessPartnerGPS", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "GPS Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGPSCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) rankCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Rank {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.Rank {
		input.General.Rank[i].BusinessPartner = input.General.BusinessPartner
		rankData := input.General.Rank[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": rankData, "function": "BusinessPartnerRank", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Rank Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToRankCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personOrganizationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonOrganization {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.PersonOrganization {
		input.General.PersonOrganization[i].BusinessPartner = input.General.BusinessPartner
		personOrganizationData := input.General.PersonOrganization[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personOrganizationData, "function": "BusinessPartnerPersonOrganization", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "PersonOrganization Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonOrganizationCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personMobilePhoneAuthCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonMobilePhoneAuth {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.PersonMobilePhoneAuth {
		input.General.PersonMobilePhoneAuth[i].BusinessPartner = input.General.BusinessPartner
		personMobilePhoneAuthData := input.General.PersonMobilePhoneAuth[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personMobilePhoneAuthData, "function": "BusinessPartnerPersonMobilePhoneAuth", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "PersonMobilePhoneAuth Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonMobilePhoneAuthCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personGoogleAccountAuthCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonGoogleAccountAuth {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.PersonGoogleAccountAuth {
		input.General.PersonGoogleAccountAuth[i].BusinessPartner = input.General.BusinessPartner
		personGoogleAccountAuthData := input.General.PersonGoogleAccountAuth[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personGoogleAccountAuthData, "function": "BusinessPartnerPersonGoogleAccountAuth", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "PersonGoogleAccountAuth Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonGoogleAccountAuthCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personInstagramAuthCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonInstagramAuth {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.PersonInstagramAuth {
		input.General.PersonInstagramAuth[i].BusinessPartner = input.General.BusinessPartner
		personInstagramAuthData := input.General.PersonInstagramAuth[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personInstagramAuthData, "function": "BusinessPartnerPersonInstagramAuth", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "PersonInstagramAuth Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonInstagramAuthCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) finInstCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FinInst {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.FinInst {
		input.General.FinInst[i].BusinessPartner = input.General.BusinessPartner
		finInstData := input.General.FinInst[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": finInstData, "function": "BusinessPartnerFinInst", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Fin Inst Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToFinInstCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) accountingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.Accounting {
		input.General.Accounting[i].BusinessPartner = input.General.BusinessPartner
		accountingData := input.General.Accounting[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "BusinessPartnerAccounting", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Accounting Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAccountingCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) generalUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	general := input.General
	generalData := dpfm_api_processing_formatter.ConvertToGeneralUpdates(general)

	sessionID := input.RuntimeSessionID
	if generalIsUpdate(generalData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "BusinessPartnerGeneral", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "General Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneralUpdates(general)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) roleUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Role {
	req := make([]dpfm_api_processing_formatter.RoleUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, role := range general.Role {
		roleData := *dpfm_api_processing_formatter.ConvertToRoleUpdates(general, role)

		if roleIsUpdate(&roleData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": roleData, "function": "BusinessPartnerRole", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Role Data cannot update"
				return nil
			}
		}
		req = append(req, roleData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToRoleUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Person {
	req := make([]dpfm_api_processing_formatter.PersonUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, person := range general.Person {
		personData := *dpfm_api_processing_formatter.ConvertToPersonUpdates(general, person)

		if personIsUpdate(&personData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personData, "function": "BusinessPartnerPerson", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Person Data cannot update"
				return nil
			}
		}
		req = append(req, personData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) addressUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	req := make([]dpfm_api_processing_formatter.AddressUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, address := range general.Address {
		addressData := *dpfm_api_processing_formatter.ConvertToAddressUpdates(general, address)

		if addressIsUpdate(&addressData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": addressData, "function": "BusinessPartnerAddress", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Address Data cannot update"
				return nil
			}
		}
		req = append(req, addressData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAddressUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) sNSUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SNS {
	req := make([]dpfm_api_processing_formatter.SNSUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, sNS := range general.SNS {
		sNSData := *dpfm_api_processing_formatter.ConvertToSNSUpdates(general, sNS)

		if sNSIsUpdate(&sNSData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": sNSData, "function": "BusinessPartnerSNS", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "SNS Data cannot update"
				return nil
			}
		}
		req = append(req, sNSData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSNSUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) gPSUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.GPS {
	req := make([]dpfm_api_processing_formatter.GPSUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, gPS := range general.GPS {
		gPSData := *dpfm_api_processing_formatter.ConvertToGPSUpdates(general, gPS)

		if gPSIsUpdate(&gPSData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": gPSData, "function": "BusinessPartnerGPS", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "GPS Data cannot update"
				return nil
			}
		}
		req = append(req, gPSData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGPSUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) rankUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Rank {
	req := make([]dpfm_api_processing_formatter.RankUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, rank := range general.Rank {
		rankData := *dpfm_api_processing_formatter.ConvertToRankUpdates(general, rank)

		if rankIsUpdate(&rankData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": rankData, "function": "BusinessPartnerRank", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Rank Data cannot update"
				return nil
			}
		}
		req = append(req, rankData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToRankUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personOrganizationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonOrganization {
	req := make([]dpfm_api_processing_formatter.PersonOrganizationUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, personOrganization := range general.PersonOrganization {
		personOrganizationData := *dpfm_api_processing_formatter.ConvertToPersonOrganizationUpdates(general, personOrganization)

		if personOrganizationIsUpdate(&personOrganizationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personOrganizationData, "function": "BusinessPartnerPersonOrganization", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "PersonOrganization Data cannot update"
				return nil
			}
		}
		req = append(req, personOrganizationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonOrganizationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personMobilePhoneAuthUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonMobilePhoneAuth {
	req := make([]dpfm_api_processing_formatter.PersonMobilePhoneAuthUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, personMobilePhoneAuth := range general.PersonMobilePhoneAuth {
		personMobilePhoneAuthData := *dpfm_api_processing_formatter.ConvertToPersonMobilePhoneAuthUpdates(general, personMobilePhoneAuth)

		if personMobilePhoneAuthIsUpdate(&personMobilePhoneAuthData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personMobilePhoneAuthData, "function": "BusinessPartnerPersonMobilePhoneAuth", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "PersonMobilePhoneAuth Data cannot update"
				return nil
			}
		}
		req = append(req, personMobilePhoneAuthData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonMobilePhoneAuthUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personGoogleAccountAuthUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonGoogleAccountAuth {
	req := make([]dpfm_api_processing_formatter.PersonGoogleAccountAuthUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, personGoogleAccountAuth := range general.PersonGoogleAccountAuth {
		personGoogleAccountAuthData := *dpfm_api_processing_formatter.ConvertToPersonGoogleAccountAuthUpdates(general, personGoogleAccountAuth)

		if personGoogleAccountAuthIsUpdate(&personGoogleAccountAuthData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personGoogleAccountAuthData, "function": "BusinessPartnerPersonGoogleAccountAuth", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "PersonGoogleAccountAuth Data cannot update"
				return nil
			}
		}
		req = append(req, personGoogleAccountAuthData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonGoogleAccountAuthUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) personInstagramAuthUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PersonInstagramAuth {
	req := make([]dpfm_api_processing_formatter.PersonInstagramAuthUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, personInstagramAuth := range general.PersonInstagramAuth {
		personInstagramAuthData := *dpfm_api_processing_formatter.ConvertToPersonInstagramAuthUpdates(general, personInstagramAuth)

		if personInstagramAuthIsUpdate(&personInstagramAuthData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": personInstagramAuthData, "function": "BusinessPartnerPersonInstagramAuth", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "PersonInstagramAuth Data cannot update"
				return nil
			}
		}
		req = append(req, personInstagramAuthData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPersonInstagramAuthUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) finInstUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.FinInst {
	req := make([]dpfm_api_processing_formatter.FinInstUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, finInst := range general.FinInst {
		finInstData := *dpfm_api_processing_formatter.ConvertToFinInstUpdates(general, finInst)

		if finInstIsUpdate(&finInstData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": finInstData, "function": "BusinessPartnerFinInst", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Fin Inst Data cannot update"
				return nil
			}
		}
		req = append(req, finInstData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToFinInstUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) accountingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	req := make([]dpfm_api_processing_formatter.AccountingUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, accounting := range general.Accounting {
		accountingData := *dpfm_api_processing_formatter.ConvertToAccountingUpdates(general, accounting)

		if accountingIsUpdate(&accountingData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "BusinessPartnerAccounting", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Accounting Data cannot update"
				return nil
			}
		}
		req = append(req, accountingData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAccountingUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func generalIsUpdate(general *dpfm_api_processing_formatter.GeneralUpdates) bool {
	businessPartner := general.BusinessPartner

	return !(businessPartner == 0)
}

func roleIsUpdate(role *dpfm_api_processing_formatter.RoleUpdates) bool {
	businessPartner := role.BusinessPartner
	businessPartnerRole := role.BusinessPartnerRole

	return !(businessPartner == 0 || businessPartnerRole == "")
}

func personIsUpdate(person *dpfm_api_processing_formatter.PersonUpdates) bool {
	businessPartner := person.BusinessPartner

	return !(businessPartner == 0)
}

func addressIsUpdate(address *dpfm_api_processing_formatter.AddressUpdates) bool {
	businessPartner := address.BusinessPartner
	addressID := address.AddressID

	return !(businessPartner == 0 || addressID == 0 )
}

func sNSIsUpdate(sNS *dpfm_api_processing_formatter.SNSUpdates) bool {
	businessPartner := sNS.BusinessPartner

	return !(businessPartner == 0)
}

func gPSIsUpdate(gPS *dpfm_api_processing_formatter.GPSUpdates) bool {
	businessPartner := gPS.BusinessPartner

	return !(businessPartner == 0)
}

func personOrganizationIsUpdate(personOrganization *dpfm_api_processing_formatter.PersonOrganizationUpdates) bool {
	businessPartner := personOrganization.BusinessPartner

	return !(businessPartner == 0)
}

func personMobilePhoneAuthIsUpdate(personMobilePhoneAuth *dpfm_api_processing_formatter.PersonMobilePhoneAuthUpdates) bool {
	businessPartner := personMobilePhoneAuth.BusinessPartner

	return !(businessPartner == 0)
}

func personGoogleAccountAuthIsUpdate(personGoogleAccountAuth *dpfm_api_processing_formatter.PersonGoogleAccountAuthUpdates) bool {
	businessPartner := personGoogleAccountAuth.BusinessPartner

	return !(businessPartner == 0)
}

func personInstagramAuthIsUpdate(personInstagramAuth *dpfm_api_processing_formatter.PersonInstagramAuthUpdates) bool {
	businessPartner := personInstagramAuth.BusinessPartner

	return !(businessPartner == 0)
}

func finInstIsUpdate(finInst *dpfm_api_processing_formatter.FinInstUpdates) bool {
	businessPartner := finInst.BusinessPartner
	finInstIdentification := finInst.FinInstIdentification

	return !(businessPartner == 0 || finInstIdentification == 0)
}

func accountingIsUpdate(accounting *dpfm_api_processing_formatter.AccountingUpdates) bool {
	businessPartner := accounting.BusinessPartner

	return !(businessPartner == 0)
}
