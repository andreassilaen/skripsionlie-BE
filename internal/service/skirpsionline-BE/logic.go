package skirpsionlineBE

import (
	"context"
	"math"
	"strconv"

	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"

	// "skripsi-online-BE/pkg/errors"

	// "google.golang.org/genproto/googleapis/gapic/metadata"
	"log"
	// "github.com/google/martian/body"
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
		err error
		// cekUser interface{}
		result  SBeEntity.B_Role
		result2 interface{}
		// result2		string
	)

	leftFour := header.UserName[:3]
	log.Println("leftFour = ", leftFour)

	tAdmin, err := s.GetAdmByLogin(ctx, header.UserName, header.PassWord)

	tEmployee, err := s.GetEmpByLogin(ctx, header.UserName, header.PassWord)

	tCustomer, err := s.GetCustByLogin(ctx, header.UserName, header.PassWord)

	if len(tAdmin) != 0 {
		// if tAdmin.UserName == header.UserName && t
		result.Role = "adm"
		result2 = tAdmin
		// return result, err
	}

	if len(tEmployee) != 0 {
		result.Role = "emp"
		result2 = tEmployee
		// return result, err
	}

	if len(tCustomer) != 0 {
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

func (s Service) InsertJoinHeaderDetailCart(ctx context.Context, header SBeEntity.InsertJoinHeaderDetailCart) (interface{}, error) {
	var (
		err error
		// cekUser interface{}
		result SBeEntity.B_Role
		// result2		interface{}
		result3 string
		// body    SBeEntity.TH_Cart

		// detailCart       SBeEntity.TD_Cart2
		// insertDetailCart []SBeEntity.TD_Cart2
		updateQtyTHTDCart SBeEntity.UpdateQtyDetailJoinTHTDCartProd
	)

	carNotPayed, err := s.skirpsionlineBE.GetHeaderCartNotPayedCustId(ctx, header.HeaderCartBody.CustId)
	log.Println("carNotPayed => ", carNotPayed)
	if err != nil {
		result3 = "Gagal get Data carNotPayed"
		return result, errors.Wrap(err, "[DATA][InsertHeaderCart]")
	}

	if len(carNotPayed) == 0 {
		println("masuk len(carNotPayed = 0)")
		_, err = s.skirpsionlineBE.InsertHeaderCart(ctx, header.HeaderCartBody)
		if err != nil {
			result3 = "Gagal insert Data"
			return result, errors.Wrap(err, "[DATA][InsertHeaderCart]")
		}

		carNotPayed, err := s.skirpsionlineBE.GetHeaderCartNotPayedCustId(ctx, header.HeaderCartBody.CustId)
		header.DetailCartBody.CartId = carNotPayed[0].CartId
		header.DetailCartBody.CartDtlQty = 1

		_, err = s.skirpsionlineBE.InsertDetailCart(ctx, header.DetailCartBody)
		if err != nil {
			result3 = "Gagal insert Data"
			return result, errors.Wrap(err, "[DATA][InsertDetailCart]")
		}

		log.Printf("InsertDetailCart >>>>> %+v", header.DetailCartBody)
	}

	if len(carNotPayed) != 0 {
		println("masuk len(carNotPayed != 0)")
		checkProdInCartNotPayed, err := s.skirpsionlineBE.GetProductInJOinTHTDCartProdByProdId(ctx, header.HeaderCartBody.CustId, carNotPayed[0].CartId, header.DetailCartBody.ProdId)
		if err != nil {
			result3 = "Gagal get Data carNotPayed"
			return result, errors.Wrap(err, "[DATA][InsertHeaderCart]")
		}
		if len(checkProdInCartNotPayed) == 0 {
			println("masuk len(checkProdInCartNotPayed = 0)")
			header.DetailCartBody.CartId = carNotPayed[0].CartId
			header.DetailCartBody.CartDtlQty = 1
			_, err = s.skirpsionlineBE.InsertDetailCart(ctx, header.DetailCartBody)
			if err != nil {
				result3 = "Gagal insert Data"
				return result, errors.Wrap(err, "[DATA][InsertDetailCart]")
			}
		} else {
			println("masuk else, update qty doang di detail")
			updateQtyTHTDCart.UpdateQtyDetailJoinTHTDCartProdBody.CartId = carNotPayed[0].CartId
			updateQtyTHTDCart.UpdateQtyDetailJoinTHTDCartProdBody.CustId = header.HeaderCartBody.CustId
			updateQtyTHTDCart.UpdateQtyDetailJoinTHTDCartProdBody.ProdId = checkProdInCartNotPayed[0].ProdId
			updateQtyTHTDCart.UpdateQtyDetailJoinTHTDCartProdBody.CartDtlQty = checkProdInCartNotPayed[0].CartDtlQty + 1
			_, err = s.skirpsionlineBE.UpdateQtyDetailJoinTHTDCart(ctx, updateQtyTHTDCart.UpdateQtyDetailJoinTHTDCartProdBody)
			if err != nil {
				result3 = "Gagal update qty"
				return result, errors.Wrap(err, "[DATA][updateQtyDetailCart]")
			}

			log.Printf(">>>>>>>>>>>> updateQtyTHTDCart %+v", updateQtyTHTDCart)
		}
	}

	result3 = "Berhasil Insert or Update qty Detail"

	// _, err = s.skirpsionlineBE.InsertHeaderCart(ctx, header.HeaderCartBody)
	// if err != nil {
	// 	result3 = "Gagal insert Data"
	// 	return result, errors.Wrap(err, "[DATA][InsertHeaderCart]")
	// }

	// body, err = s.skirpsionlineBE.GetHeaderCartLastData(ctx)
	// if err != nil {
	// 	result3 = "Gagal getlast Data"
	// 	return result, errors.Wrap(err, "[DATA][InsertDetailCart]")
	// }
	// header.DetailCartBody.CartId = body.CartId

	// var i = 0
	// var a =10
	// for x := range  {

	// }

	// log.Println("cel len header.DetailCartBody => ", len(header.DetailCartBody))
	// if len(header.DetailCartBody) >= 1 {
	// 	for x := range header.DetailCartBody {
	// 		detailCart = SBeEntity.TD_Cart2{
	// 			CartId:     body.CartId,
	// 			ProdId:     header.DetailCartBody[x].ProdId,
	// 			CartDtlQty: header.DetailCartBody[x].CartDtlQty,
	// 		}

	// 		log.Println("cek x => ", x)
	// 		log.Println("cek detailCart => ", detailCart)

	// 		insertDetailCart = append(insertDetailCart, detailCart)

	// 	}

	// }

	// log.Println("insertDetailCart", insertDetailCart)
	// limitzI := 50
	// totalzI := len(insertDetailCart)
	// countzI := int(math.Ceil(float64(totalzI) / float64(limitzI)))
	// for i := 0; i < countzI; i++ {
	// 	startzI := limitzI * i
	// 	endzI := limitzI * (i + 1)
	// 	if endzI > totalzI {
	// 		endzI = totalzI
	// 	}
	// 	tempUpdatez := insertDetailCart[startzI:endzI]
	// 	err = s.skirpsionlineBE.NewInsertDetailCart(ctx, tempUpdatez)
	// 	if err != nil {
	// 		log.Println(err, "[Service][InsertDetailCart]")
	// 		// return result, errors.Wrap(err, "[Service][InsertDetailCart]")
	// 	}
	// }
	// log.Println("masokDetail-3")

	// // _, err = s.skirpsionlineBE.InsertDetailCart(ctx, header.DetailCartBody)
	// // if err != nil {
	// // 	result3 = "Gagal insert Data"
	// // 	return result, errors.Wrap(err, "[DATA][InsertDetailCart]")
	// // }
	// // result3 = "sukses Insert Header & Detail cart"

	return result3, err
}

func (s Service) InsertJoinHeaderDetailTran(ctx context.Context, header SBeEntity.InsertJoinHeaderDetailTran) (interface{}, error) {
	var (
		err error
		// cekUser interface{}
		result SBeEntity.B_Role
		// result2		interface{}
		result3 string
		body    SBeEntity.TH_Transaction

		detailTran       SBeEntity.TD_Transaction2
		insertDetailTran []SBeEntity.TD_Transaction2
	)

	_, err = s.skirpsionlineBE.InsertHeaderTran(ctx, header.HeaderTranBody)
	if err != nil {
		result3 = "Gagal insert Data"
		return result, errors.Wrap(err, "[DATA][InsertHeaderTran]")
	}

	body, err = s.skirpsionlineBE.GetHeaderTranLastDataByCusId(ctx, header.HeaderTranBody.CustId)
	log.Printf("body => %+v", body)
	if err != nil {
		result3 = "Gagal getlast Data"
		return result, errors.Wrap(err, "[DATA][GetHeaderTranLastDataByCusId]")
	}

	cartNo, _ := strconv.Atoi(header.HeaderTranBody.CartId)
	_, err = s.skirpsionlineBE.UpdateHeaderCartPeyed(ctx, cartNo)
	log.Printf("cartNo => %+v", cartNo)
	if err != nil {
		result3 = "Gagal updateCartPayed"
		return result, errors.Wrap(err, "[DATA][UpdateHeaderCartPeyed]")
	}
	// body, err = s.skirpsionlineBE.GetHeaderCartLastData(ctx)
	// if err != nil {
	// 	result3 = "Gagal getlast Data"
	// 	return result, errors.Wrap(err, "[DATA][InsertDetailCart]")
	// }

	log.Println("cel len header.DetailTranBody => ", len(header.DetailTranBody))
	if len(header.DetailTranBody) >= 1 {
		for x := range header.DetailTranBody {
			detailTran = SBeEntity.TD_Transaction2{
				TraId:        body.TranId,
				ProdId:       header.DetailTranBody[x].ProdId,
				TraDtlQty:    header.DetailTranBody[x].TraDtlQty,
				TraDtlPrice:  header.DetailTranBody[x].TraDtlPrice,
				TraDtlAmount: header.DetailTranBody[x].TraDtlAmount,
			}

			log.Println("cek x => ", x)
			log.Println("cek detailCart => ", detailTran)

			insertDetailTran = append(insertDetailTran, detailTran)

		}

	}

	log.Println("insertDetailTran", insertDetailTran)
	limitzI := 50
	totalzI := len(insertDetailTran)
	countzI := int(math.Ceil(float64(totalzI) / float64(limitzI)))
	for i := 0; i < countzI; i++ {
		startzI := limitzI * i
		endzI := limitzI * (i + 1)
		if endzI > totalzI {
			endzI = totalzI
		}
		tempUpdatez := insertDetailTran[startzI:endzI]
		err = s.skirpsionlineBE.NewInsertDetailTran(ctx, tempUpdatez)
		if err != nil {
			log.Println(err, "[Service][InsertDetailCart]")
			// return result, errors.Wrap(err, "[Service][InsertDetailCart]")
		}
	}
	log.Println("masokDetail-3")

	return result3, err
}

// LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL

// func (s Service) InsertEmployee(ctx context.Context, header SBeEntity.InsertEmployee) (string, error) {
// 	var (
// 		result string
// 		err    error
// 	)

// 	result, err = s.skirpsionlineBE.InsertEmployee(ctx, header.InsertEmployeeBody)
// 	log.Println("header Service = ", header)
// 	if err != nil {
// 		result = "Gagal Insert"
// 		return result, errors.Wrap(err, "[Service][InsertEmployee]")
// 	} else {
// 		result = "Sukses InsertEmployee"
// 	}
// 	return result, err

// }

// func (s Service) InsertMainCart(ctx context.Context, headers SBeEntity.MainCart) ([]SBeEntity.JoinTHTDCartProd,error) {
// 	var (
// 		header		SBeEntity.JoinTDTraProdByTraId
// 		headers		[]S
// 	)
// }

// func (s Service) InsertCartMain(ctx context.Context, header SBeEntity.InsertJoinHeaderDetailCart, cusId string) (interface{}, error) {
// 	var (
// 		err error
// 		// cekUser interface{}
// 		// result SBeEntity.B_Role
// 		// result2		interface{}
// 		result3 string
// 		// body    SBeEntity.TH_Cart

// 		// detailCart       SBeEntity.TD_Cart2
// 		// insertDetailCart []SBeEntity.TD_Cart2
// 	)

// 	_, err = s.skirpsionlineBE.GetHeaderCartNotPayedCustId(ctx, cusId)
// 	if err != nil {
// 		baru, err = s.skirpsionlineBE.InsertHeaderCart(ctx, header.HeaderCartBody)
// 	}

// 	// _, err = s.skirpsionlineBE.GetDetailTranByTraId(ctx,traId)

// 	return result3, err
// }

func (s Service) GetAllProductCartNotPayedByCusId(ctx context.Context, custId string) ([]SBeEntity.JoinTHTDCartProd, error) {
	var (
		headers []SBeEntity.JoinTHTDCartProd
	)
	checkCart, err := s.skirpsionlineBE.GetHeaderCartNotPayedCustId(ctx, custId)
	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][JOIN][GetHeaderCartNotPayedCustId]")
	}

	getListCart, err := s.skirpsionlineBE.GetListJoinTHTDCartProdByCustIdAndCartId(ctx, custId, checkCart[0].CartId)
	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][JOIN][GetListJoinTHTDCartProdByCustIdAndCartId]")
	}

	headers = getListCart
	return headers, err

}

func (s Service) GetUserMainByIdAndRole(ctx context.Context, userId string, role string) (interface{}, error) {
	var (
		// cek  SBeEntity.T_UserMain
		// header  SBeEntity.T_UserMain
		header interface{}

		err error
	)

	if role == "cus" {
		header, err = s.skirpsionlineBE.GetCustById(ctx, userId)
		if err != nil {
			return header, errors.Wrap(err, "[SERVICE][GetUserMainByIdAndRole][GetCustById]")
		}
	} else if role == "adm" {
		header, err = s.skirpsionlineBE.GetAdmById(ctx, userId)
		if err != nil {
			return header, errors.Wrap(err, "[SERVICE][GetUserMainByIdAndRole][GetAdmById]")
		}
	} else if role == "emp" {
		header, err = s.skirpsionlineBE.GetEmpById(ctx, userId)
		if err != nil {
			return header, errors.Wrap(err, "[SERVICE][GetUserMainByIdAndRole][GetEmpById]")
		}
	}

	log.Printf(" header usermain => %+v", header)

	return header, err
}

func (s Service) UpdateUserMain(ctx context.Context, body SBeEntity.UpdateUserMain, userId string, role string) (interface{}, error) {
	var (
		// cek  SBeEntity.T_UserMain
		// header  SBeEntity.T_UserMain
		header  interface{}
		err     error
		bodyCus SBeEntity.T_Customer2
		bodyAdm SBeEntity.T_Admin2
		bodyEmp SBeEntity.T_Employee2
	)

	if role == "cus" {
		log.Printf(" Masuk role CUS !!!!!!!!")
		bodyCus.CustName = body.UpdateUserMainBody.UserName
		bodyCus.CustUserName = body.UpdateUserMainBody.UserUserName
		bodyCus.CustPassWord = body.UpdateUserMainBody.UserPassWord
		bodyCus.CustEmail = body.UpdateUserMainBody.UserEmail
		bodyCus.CustPhone = body.UpdateUserMainBody.UserPhone
		bodyCus.CustAddress = body.UpdateUserMainBody.UserAddress
		header, err = s.skirpsionlineBE.UpdateCustomerById(ctx, bodyCus, userId)
		if err != nil {
			return header, errors.Wrap(err, "[SERVICE][UpdateUserMain][UpdateCustomerById]")
		}
	} else if role == "adm" {
		log.Printf(" Masuk role ADM !!!!!!!!")
		bodyAdm.AdmName = body.UpdateUserMainBody.UserName
		bodyAdm.AdmUserName = body.UpdateUserMainBody.UserUserName
		bodyAdm.AdmPassWord = body.UpdateUserMainBody.UserPassWord
		bodyAdm.AdmEmail = body.UpdateUserMainBody.UserEmail
		bodyAdm.AdmPhone = body.UpdateUserMainBody.UserPhone
		bodyAdm.AdmAddress = body.UpdateUserMainBody.UserAddress
		header, err = s.skirpsionlineBE.UpdateAdminById(ctx, bodyAdm, userId)
		if err != nil {
			return header, errors.Wrap(err, "[SERVICE][UpdateUserMain][UpdateAdminById]")
		}
	} else if role == "emp" {
		log.Printf(" Masuk role EMP !!!!!!!!")
		bodyEmp.EmpName = body.UpdateUserMainBody.UserName
		bodyEmp.EmpUserName = body.UpdateUserMainBody.UserUserName
		bodyEmp.EmpPassWord = body.UpdateUserMainBody.UserPassWord
		bodyEmp.EmpEmail = body.UpdateUserMainBody.UserEmail
		bodyEmp.EmpPhone = body.UpdateUserMainBody.UserPhone
		bodyEmp.EmpAddress = body.UpdateUserMainBody.UserAddress
		header, err = s.skirpsionlineBE.UpdateEmployeeById(ctx, bodyEmp, userId)
		if err != nil {
			return header, errors.Wrap(err, "[SERVICE][UpdateUserMain][UpdateAdminById]")
		}
	}

	log.Printf(" header usermain => %+v", header)

	return header, err
}

func (s Service) InsertOrderByAdminCancelMain(ctx context.Context, header SBeEntity.InsertOrder) (interface{}, error) {
	var (
		// header SBeEntity.
		// header 		SBeEntity.InsertOrder
		result string
		err    error
	)

	tra_id, _ := strconv.Atoi(header.InsertOrderBody.TraId)

	_, err = s.skirpsionlineBE.UpdateTHTranChecked(ctx, tra_id)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][InsertOrderByAdminCancelMain][UpdateTHTranChecked]")
	}

	result, err = s.skirpsionlineBE.InsertOrder(ctx, header.InsertOrderBody)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertOrderByAdminCancelMain][InsertOrder]")
	} else {
		result = "Sukses InsertOrderByAdminCancelMain"
	}
	return result, err

}

func (s Service) UpdateStock(ctx context.Context, traid string) (string, error) {
	var (
		result string
	)

	header, err := s.skirpsionlineBE.GetJoinTDTranProdCustByTraId(ctx, traid)
	if err != nil {
		log.Println("GagalGetData")
		result = "Gagal Get Data"
		return result, errors.Wrap(err, "[DATA][GetProductStockDetail]")
	}
	log.Println("header GetJoinTDTranProdCustByTraId luar err", header)

	for x := range header {
		_, err = s.skirpsionlineBE.UpdateProdStockById(ctx, header[x].ProdStock-header[x].TraDtlQty, header[x].ProdId)
		if err != nil {
			log.Println("GagalUpdateStock ->", x)
			return result, errors.Wrap(err, "[DATA][UpdateStock]")
		}
		log.Println("header UpdateProdStockById dalam for", header[x].ProdStock-header[x].TraDtlQty, header[x].ProdId)
	}

	result = "Berhasil"

	return result, err

}

func (s Service) InsertOrderByAdminAccMain(ctx context.Context, header SBeEntity.InsertOrder) (interface{}, error) {
	var (
		// header SBeEntity.
		// header 		SBeEntity.InsertOrder
		result string
		err    error
	)

	tra_id, _ := strconv.Atoi(header.InsertOrderBody.TraId)

	_, err = s.skirpsionlineBE.UpdateTHTranChecked(ctx, tra_id)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][InsertOrderByAdminAccMain][UpdateTHTranChecked]")
	}

	_, err = s.UpdateStock(ctx, header.InsertOrderBody.TraId)

	result, err = s.skirpsionlineBE.InsertOrderAcc(ctx, header.InsertOrderBody)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertOrderByAdminAccMain][InsertOrder]")
	} else {
		result = "Sukses InsertOrderByAdminAccMain"
	}
	return result, err

}

// func (s Service) GetDataOrderDeliveryByCustId(ctx context.Context, custId string) (interface{}, error) {
// 	var (
// 		header     interface{}
// 		err        error
// 	)

// 		api1, _ := s.GetJoinOrdTHTraByCustId(ctx, custId)
// 		api2, _ := s.GetJoinOrdTHTraDelByCustId(ctx, custId)
// 		if len(api1) =0 {
// 		header = api2
// 		}
// 		if len(api2) =0 {
// 		header = api1
// 		}

// 		append

// 	return header, err
// 	}
