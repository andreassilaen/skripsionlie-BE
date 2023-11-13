package skirpsionlineBE

import (
	"context"
	"errors"

	"skripsi-online-BE/internal/entity"
	"skripsi-online-BE/internal/entity/auth"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	jaegerLog "skripsi-online-BE/pkg/log"

	"github.com/opentracing/opentracing-go"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	// GetTableSelisih(ctx context.Context, ptid int, code string, tglsttk string, sttktype string) ([]sttkEntity.GetTableSelisih, error)

	GetAllAdmin(ctx context.Context) ([]SBeEntity.T_Admin, error)
	GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error)
	GetAllCategory(ctx context.Context) ([]SBeEntity.T_Category, error)
	InsertProduct(ctx context.Context, header SBeEntity.T_Product2) (string, error)
	GetAllCart(ctx context.Context) ([]SBeEntity.TH_Cart, error)
	GetCustByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Customer, error)
	GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error)
	GetCustById(ctx context.Context, custId string) ([]SBeEntity.T_Customer, error)
	GetAdmById(ctx context.Context, userId string) ([]SBeEntity.T_Admin, error)
	GetEmpById(ctx context.Context, userId string) ([]SBeEntity.T_Employee, error)
	GetProdById(ctx context.Context, prodId string) ([]SBeEntity.T_Product, error)
	InsertCustomer(ctx context.Context, header SBeEntity.T_Customer) (string, error)
	GetCountCust(ctx context.Context) (int, error)
	GetAdmLastData(ctx context.Context) (SBeEntity.T_Admin, error)
	GetCustLastData(ctx context.Context) (SBeEntity.T_Customer, error)
	GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error)
	GetHeaderCartLastData(ctx context.Context) (SBeEntity.TH_Cart, error)
	InsertAdmin(ctx context.Context, header SBeEntity.T_Admin) (string, error)

	GetCartByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error)
	InsertHeaderCart(ctx context.Context, header SBeEntity.TH_Cart2) (string, error)
	UpdateHeaderCartPeyed(ctx context.Context, cartId int) (string, error)

	InsertDetailCart(ctx context.Context, header SBeEntity.TD_Cart2) (string, error)
	NewInsertDetailCart(ctx context.Context, user []SBeEntity.TD_Cart2) error

	GetAllEmployee(ctx context.Context) ([]SBeEntity.T_Employee, error)
	GetEmpLastData(ctx context.Context) (SBeEntity.T_Employee, error)
	GetEmpByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Employee, error)
	InsertEmployee(ctx context.Context, header SBeEntity.T_Employee) (string, error)

	GetAllHeaderTran(ctx context.Context) ([]SBeEntity.TH_Transaction, error)
	GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error)
	GetHeaderCartNotPayedCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error)
	GetHeaderTranLastDataByCusId(ctx context.Context, custId string) (SBeEntity.TH_Transaction, error)
	InsertHeaderTran(ctx context.Context, header SBeEntity.TH_Transaction2) (string, error)
	GetAllHeaderTranByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Transaction, error) 
	UpdateTHTranChecked(ctx context.Context, traId int) (string, error)

	GetDetailTranByTraId(ctx context.Context, traId string) ([]SBeEntity.TD_Transaction, error)
	InsertDetailTran(ctx context.Context, header SBeEntity.TD_Transaction2) (string, error)
	NewInsertDetailTran(ctx context.Context, user []SBeEntity.TD_Transaction2) error /// buat sendiri

	UpdateAdminById(ctx context.Context, header SBeEntity.T_Admin2, admId string) (string, error)
	UpdateEmployeeById(ctx context.Context, header SBeEntity.T_Employee2, empId string) (string, error)
	UpdateCustomerById(ctx context.Context, header SBeEntity.T_Customer2, cusId string) (string, error)
	UpdateProdById(ctx context.Context, header SBeEntity.T_Product3, prodId string) (string, error)

	GetAllOrder(ctx context.Context) ([]SBeEntity.T_Order, error)
	InsertOrder(ctx context.Context, header SBeEntity.T_Order2) (string, error)
	InsertOrderAcc(ctx context.Context, header SBeEntity.T_Order2) (string, error)
	UpdateOrderOnDeliveryYes(ctx context.Context, ordId int) (string, error)

	GetAllDelivery(ctx context.Context) ([]SBeEntity.T_Delivery, error)
	GetDeliverByEmpId(ctx context.Context, empId string) ([]SBeEntity.T_Delivery, error)
	InsertDeliveryProcess(ctx context.Context, header SBeEntity.T_Delivery2) (string, error)
	UpdateDeliveryDone(ctx context.Context, ordId string) (string, error)

	GetAllRekening(ctx context.Context) ([]SBeEntity.T_Rekening, error)
	GetRekByRekId(ctx context.Context, rekId int) ([]SBeEntity.T_Rekening, error)

	GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error)

	GetJoinOrdCustTHTra(ctx context.Context) ([]SBeEntity.JoinOrdCustTHTra, error)
	GetJoinOrdCustTHTraByOrdId(ctx context.Context, ordId int) ([]SBeEntity.JoinOrdCustTHTra, error)
	GetJoinTHTraRekByCusId(ctx context.Context, custId string) ([]SBeEntity.JoinTHTraRek, error)
	GetJoinTDTraProdByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTraProdByTraId, error)
	GetJoinOrdTHTDTraProdByOrdId(ctx context.Context, ordId string) ([]SBeEntity.JoinOrdTHTDTraProdByOrdId, error)
	GetListJoinTHTDCartProdByCustIdAndCartId(ctx context.Context, custId string, cartId string) ([]SBeEntity.JoinTHTDCartProd, error)
	GetProductInJOinTHTDCartProdByProdId(ctx context.Context, custId string, cartId string, prodId string) ([]SBeEntity.JoinTHTDCartProd, error)
	GetJoinOrdTHTraByCustId(ctx context.Context, custId string) ([]SBeEntity.JoinOrdTHTraByCustId, error)
	UpdateQtyDetailJoinTHTDCart(ctx context.Context, header SBeEntity.JoinTHTDCartProd2) (string, error)
	GetCountDashboardAdmin(ctx context.Context) (SBeEntity.CountTHTraOrdDel, error)


	DeleteProductByProdId(ctx context.Context, prodId int) (string, error) 
}

// AuthData ...
type AuthData interface {
	CheckAuth(ctx context.Context, _token, code string) (auth.Auth, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	skirpsionlineBE Data
	authData        AuthData
	tracer          opentracing.Tracer
	logger          jaegerLog.Factory
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(skripsionlineData Data, authData AuthData, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	// Assign variable dari parameter ke object
	return Service{
		skirpsionlineBE: skripsionlineData,
		authData:        authData,
		tracer:          tracer,
		logger:          logger,
	}
}

func (s Service) checkPermission(ctx context.Context, _permissions ...string) error {
	claims := ctx.Value(entity.ContextKey("claims"))
	if claims != nil {
		actions := claims.(entity.ContextValue).Get("permissions").(map[string]interface{})
		for _, action := range actions {
			permissions := action.([]interface{})
			for _, permission := range permissions {
				for _, _permission := range _permissions {
					if permission.(string) == _permission {
						return nil
					}
				}
			}
		}
	}
	return errors.New("401 unauthorized")
}
