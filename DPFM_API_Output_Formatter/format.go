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
