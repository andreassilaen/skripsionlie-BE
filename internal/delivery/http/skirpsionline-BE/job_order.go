package skirpsionlineBE

import (
	// "bytes"
	"context"
	jaegerLog "skripsi-online-BE/pkg/log"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// "skripsi-online-BE/pkg/errors"

	"github.com/opentracing/opentracing-go"
)

// IsttkSvc is an interface to sttk Service
type IskripsionlineSvc interface {
	GetAllAdmin(ctx context.Context) ([]SBeEntity.T_Admin, error)
	GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error)
	GetAllCategory(ctx context.Context) ([]SBeEntity.T_Category, error)
	InsertProduct(ctx context.Context, header SBeEntity.InsertProduct) (string, error)
	GetAllCart(ctx context.Context) ([]SBeEntity.TH_Cart, error)
	GetCustByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Customer, error)
	GetCustById(ctx context.Context, custId string) ([]SBeEntity.T_Customer, error)
	GetProdById(ctx context.Context, prodId string) ([]SBeEntity.T_Product, error)
	GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error)
	InsertCustomer(ctx context.Context, header SBeEntity.InsertCustomer) (string, error)
	GetAdmLastData(ctx context.Context) (SBeEntity.T_Admin, error)
	GetCustLastData(ctx context.Context) (SBeEntity.T_Customer, error)
	GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error)
	InsertAdmin(ctx context.Context, header SBeEntity.InsertAdmin) (string, error)
	CheckUser(ctx context.Context, header SBeEntity.B_ChekUser) (interface{}, SBeEntity.B_Role, error)

	GetCartByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error)
	InsertHeaderCart(ctx context.Context, header SBeEntity.InsertHeaderCart) (string, error)
	UpdateHeaderCartPayed(ctx context.Context, cartId int) (string, error)

	InsertDetailCart(ctx context.Context, header SBeEntity.InsertDetailCart) (string, error)

	GetAllEmployee(ctx context.Context) ([]SBeEntity.T_Employee, error)
	GetEmpLastData(ctx context.Context) (SBeEntity.T_Employee, error)
	GetEmpByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Employee, error)
	InsertEmployee(ctx context.Context, header SBeEntity.InsertEmployee) (string, error)

	GetAllHeaderTran(ctx context.Context) ([]SBeEntity.TH_Transaction, error)
	GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error)
	GetHeaderCartNotPayedCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error)
	InsertHeaderTran(ctx context.Context, header SBeEntity.InsertHeaderTransaction) (string, error)
	GetAllHeaderTranByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Transaction, error)
	UpdateTHTranChecked(ctx context.Context, traId int) (string, error)

	GetDetailTranByTraId(ctx context.Context, traId string) ([]SBeEntity.TD_Transaction, error)
	InsertDetailTransaction(ctx context.Context, header SBeEntity.InsertDetailTransaction) (string, error)

	UpdateAdminById(ctx context.Context, header SBeEntity.UpdateAdminById, admId string) (string, error)
	UpdateEmployeeById(ctx context.Context, header SBeEntity.UpdateEmployeeById, empId string) (string, error)
	UpdateCustomerById(ctx context.Context, header SBeEntity.UpdateCustomerById, cusId string) (string, error)
	UpdateProdById(ctx context.Context, header SBeEntity.UpdateProdById, prodId string) (string, error)
	UpdateProdStockById(ctx context.Context, prodStock int, prodId int) (string, error)

	GetAllOrder(ctx context.Context) ([]SBeEntity.T_Order, error)
	InsertOrder(ctx context.Context, header SBeEntity.InsertOrder) (string, error)
	InsertOrderAcc(ctx context.Context, header SBeEntity.InsertOrder) (string, error)
	UpdateOrderOnDeliveryYes(ctx context.Context, ordId int) (string, error)

	GetAllDelivery(ctx context.Context) ([]SBeEntity.T_Delivery, error)
	GetDeliverByEmpId(ctx context.Context, empId string) ([]SBeEntity.T_Delivery, error)
	InsertDeliveryProcess(ctx context.Context, header SBeEntity.InsertDelivery) (string, error)
	UpdateDeliveryDone(ctx context.Context, ordId string) (string, error)

	GetAllRekening(ctx context.Context) ([]SBeEntity.T_Rekening, error)
	GetRekByRekId(ctx context.Context, rekId int) ([]SBeEntity.T_Rekening, error)

	GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error)

	GetJoinOrdCustTHTra(ctx context.Context) ([]SBeEntity.JoinOrdCustTHTra, error)
	GetJoinOrdCustTHTraByOrdId(ctx context.Context, ordId int) ([]SBeEntity.JoinOrdCustTHTra, error)
	GetJoinTDTraProdByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTraProdByTraId, error)
	GetJoinTHTraRekByCusId(ctx context.Context, custId string) ([]SBeEntity.JoinTHTraRek, error)
	GetJoinOrdTHTDTraProdByOrdId(ctx context.Context, ordId string) ([]SBeEntity.JoinOrdTHTDTraProdByOrdId, error)
	GetListJoinTHTDCartProdByCustIdAndCartId(ctx context.Context, custId string, cartId string) ([]SBeEntity.JoinTHTDCartProd, error)
	GetAllProductCartNotPayedByCusId(ctx context.Context, custId string) ([]SBeEntity.JoinTHTDCartProd, error)
	GetProductInJOinTHTDCartProdByProdId(ctx context.Context, custId string, cartId string, prodId string) ([]SBeEntity.JoinTHTDCartProd, error)
	GetJoinOrdTHTraByCustId(ctx context.Context, custId string) ([]SBeEntity.JoinOrdTHTraByCustId, error)
	GetJoinOrdTHTraDelByCustId(ctx context.Context, traId string) ([]SBeEntity.JoinOrdTHTraDelByCustId, error)
	UpdateQtyDetailJoinTHTDCart(ctx context.Context, header SBeEntity.UpdateQtyDetailJoinTHTDCartProd) (string, error)
	GetCountDashboardAdmin(ctx context.Context) (SBeEntity.CountTHTraOrdDel, error)
	GetReportOrdTHTraByOrdDate(ctx context.Context, startDate string, endDate string) ([]SBeEntity.JoinReportOrdTHTra, error)
	GetDetailReportByOrdId(ctx context.Context, ordId int) (SBeEntity.JoinDetailReport, error)
	GetJoinTDTranProdCustByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTranProdCustByTraId, error)
	UpdateStock(ctx context.Context, traid string) (string, error)

	TokenUser(ctx context.Context) error

	GetUserMainByIdAndRole(ctx context.Context, userId string, role string) (interface{}, error)
	UpdateUserMain(ctx context.Context, body SBeEntity.UpdateUserMain, userId string, role string) (interface{}, error)
	InsertOrderByAdminCancelMain(ctx context.Context, header SBeEntity.InsertOrder) (interface{}, error)
	InsertOrderByAdminAccMain(ctx context.Context, header SBeEntity.InsertOrder) (interface{}, error)

	InsertJoinHeaderDetailCart(ctx context.Context, header SBeEntity.InsertJoinHeaderDetailCart) (interface{}, error)
	InsertJoinHeaderDetailTran(ctx context.Context, header SBeEntity.InsertJoinHeaderDetailTran) (interface{}, error)


	InsertCategory(ctx context.Context, ctgType string) (string, error)
	InsertRekening(ctx context.Context, rekBank string, rekNumber int, rekName string) (string, error) 
	
	
	DeleteProductByProdId(ctx context.Context, prodId int) (string, error)
}

type (
	// Handler ...
	Handler struct {
		skripsionlineSvc IskripsionlineSvc
		tracer           opentracing.Tracer
		logger           jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IskripsionlineSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		skripsionlineSvc: is,
		tracer:           tracer,
		logger:           logger,
	}
}
