package skirpsionlineBE

import (
	"context"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// "skripsi-online-BE/pkg/errors"

	// "skripsi-online-BE/pkg/errors"

	// "google.golang.org/genproto/googleapis/gapic/metadata"
	"log"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

// func (s Service) CheckUser(ctx context.Context, username string, password string) (interface{}, error) {
// 	var (
// 		err		error
// 		cekUser interface{}
// 	)

// 	cekUser, err = s.skirpsionlineBE.GetAllAdmin(ctx)

// 	if username == "adm_adm"

// 	if procode != "" && status != "" && submitdate != "" { //-------------------------------- 3 Parms-------------
// 		filterArray, err = s.qrData.GetFilterHeaderAll(ctx, procode, status, submitdate)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetFilterHeaderAll]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetFilterHeaderAll]")
// 		}
// 	} else if procode != "" && status != "" && submitdate == "" { //-------------------------------- 2 Parms-------------
// 		filterArray, err = s.qrData.GetFilterHeaderProcodeStatus(ctx, procode, status)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetFilterHeaderProcodeStatus]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetFilterHeaderProcodeStatus]")
// 		}
// 	} else if procode != "" && status == "" && submitdate != "" {
// 		filterArray, err = s.qrData.GetFilterHeaderProcodeSubmitdate(ctx, procode, submitdate)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetFilterHeaderProcodeSubmitdate]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetFilterHeaderProcodeSubmitdate]")
// 		}
// 	} else if procode == "" && status != "" && submitdate != "" {
// 		filterArray, err = s.qrData.GetFilterHeaderStatusSubmitdate(ctx, status, submitdate)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetFilterHeaderStatusSubmitdate]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetFilterHeaderStatusSubmitdate]")
// 		}
// 	} else if procode == "" && status == "" && submitdate == "" { //-------------------------------- 3 Parms-------------
// 		filterArray, err = s.qrData.GetBatchToMonitor(ctx)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetBatchToMonitor]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetBatchToMonitor]")
// 		}
// 	} else if procode != "" && status == "" && submitdate == "" { //-------------------------------- 1 Parms-------------
// 		filterArray, err = s.qrData.GetBatchMonitorByProcode(ctx, procode)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetBatchMonitorByProcode]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetBatchMonitorByProcode]")
// 		}
// 	} else if procode == "" && status != "" && submitdate == "" { //-------------------------------- 1 Parms-------------
// 		filterArray, err = s.qrData.GetBatchMonitorByStatus(ctx, status)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetBatchMonitorByStatus]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetBatchMonitorByStatus]")
// 		}
// 	} else if procode == "" && status == "" && submitdate != "" { //-------------------------------- 1 Parms-------------
// 		filterArray, err = s.qrData.GetBatchMonitorByDate(ctx, submitdate)
// 		if err != nil {
// 			log.Println("[Service][GetFilterHeader][GetBatchMonitorByDate]", err.Error())
// 			return filterArray, errors.Wrap(err, "[Service][GetFilterHeader][GetBatchMonitorByDate]")
// 		}
// 	}
// 	if err != nil {
// 		return filterArray, errors.Wrap(err, "[Service][GetFilterHeader]")
// 	}

// 	return cekUser, nil

// }



func (s Service) CheckUser(ctx context.Context, header SBeEntity.B_ChekUser) (interface{}, SBeEntity.B_Role, error) {
	var (
		err		error
		// cekUser interface{}
		result		SBeEntity.B_Role
		result2		interface{}
		// result2		string
	)

	leftFour:= header.UserName[:3]
    log.Println("leftFour = ",leftFour)

	tAdmin, err := s.GetAdmByLogin(ctx, header.UserName, header.PassWord)

	tEmployee, err := s.GetEmpByLogin(ctx, header.UserName, header.PassWord)

	tCustomer, err := s.GetCustByLogin(ctx, header.UserName, header.PassWord)

	if len(tAdmin) != 0 {
		// if tAdmin.UserName == header.UserName && t
		result.Role = "adm"
			result2 = tAdmin
		// return result, err
	}

	if len (tEmployee) != 0 {
		result.Role = "emp"
	result2 = tEmployee
		// return result, err
	}

	if len (tCustomer) != 0 {
		result.Role = "cus"
	result2 = tCustomer
		// return result, err
	}

	// errs := &SBeEntity.CustomError{Code: 101, Message: "Error 101: This is a custom error with code 101."}

	if len(tAdmin) == 0 && len(tCustomer) == 0 && len(tEmployee) == 0 {
		// error.Error.()
		// return result, errs
		result.Role = "gagal" 
	}

// if len(tAdmin) != 0 {
// 	result2 = tAdmin
// 	// return tAdmin, result, err
	
// }
// if len(tCustomer) != 0 {
// 	result2 = tCustomer
// 	// return tCustomer, result, err
	
// }
return result2, result, err
}