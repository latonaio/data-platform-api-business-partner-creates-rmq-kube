package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-business-partner-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-business-partner-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var generalCreates *dpfm_api_output_formatter.General
	for _, fn := range accepter {
		switch fn {
		case "General":
			c.generalCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
			generalCreates = dpfm_api_output_formatter.ConvertToGeneralCreates(subfuncSDC)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General: generalCreates,
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
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General: general,
	}

	return data
}

func (c *DPFMAPICaller) generalCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_business_partner_general_data?????????
	generalData := subfuncSDC.Message.General
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "BusinessPartnerGeneral", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToGeneralCreates(subfuncSDC)

	return data
}

func (c *DPFMAPICaller) generalUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	req := dpfm_api_processing_formatter.ConvertToGeneralUpdates(input.General)

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "BusinessPartnerGeneral", "runtime_session_id": 123})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("General Data cannot insert")
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "General Data cannot insert"
		return nil
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToGeneralUpdates(req)

	return data
}
