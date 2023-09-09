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
	var finInst *[]dpfm_api_output_formatter.FinInst
	var role *[]dpfm_api_output_formatter.Role
	var accounting *[]dpfm_api_output_formatter.Accounting
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalCreateSql(nil, mtx, input, output, errs, log)
		case "FinInst":
			finInst = c.finInstCreateSql(nil, mtx, input, output, errs, log)
		case "Role":
			role = c.roleCreateSql(nil, mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:    general,
		FinInst:    finInst,
		Role:       role,
		Accounting: accounting,
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
	var finInst *[]dpfm_api_output_formatter.FinInst
	var role *[]dpfm_api_output_formatter.Role
	var accounting *[]dpfm_api_output_formatter.Accounting
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		case "FinInst":
			finInst = c.finInstUpdateSql(mtx, input, output, errs, log)
		case "Role":
			role = c.roleUpdateSql(mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:    general,
		FinInst:    finInst,
		Role:       role,
		Accounting: accounting,
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
	// data_platform_business_partner_general_dataの更新
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
	// data_platform_business_partner_fin_inst_dataの更新
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
	// data_platform_business_partner_role_dataの更新
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
	// data_platform_business_partner_fin_inst_dataの更新
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

func finInstIsUpdate(finInst *dpfm_api_processing_formatter.FinInstUpdates) bool {
	businessPartner := finInst.BusinessPartner
	finInstIdentification := finInst.FinInstIdentification
	validityEndDate := finInst.ValidityEndDate
	validityStartDate := finInst.ValidityStartDate

	return !(businessPartner == 0 || finInstIdentification == 0 || validityEndDate == "" || validityStartDate == "")
}

func roleIsUpdate(role *dpfm_api_processing_formatter.RoleUpdates) bool {
	businessPartner := role.BusinessPartner
	businessPartnerRole := role.BusinessPartnerRole
	validityEndDate := role.ValidityEndDate
	validityStartDate := role.ValidityStartDate

	return !(businessPartner == 0 || businessPartnerRole == "" || validityEndDate == "" || validityStartDate == "")
}

func accountingIsUpdate(accounting *dpfm_api_processing_formatter.AccountingUpdates) bool {
	businessPartner := accounting.BusinessPartner

	return !(businessPartner == 0)
}
