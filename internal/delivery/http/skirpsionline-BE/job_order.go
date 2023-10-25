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
	GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error)
	InsertCustomer(ctx context.Context, header SBeEntity.InsertCustomer) (string, error)
	GetAdmLastData(ctx context.Context) (SBeEntity.T_Admin, error)
	GetCustLastData(ctx context.Context) (SBeEntity.T_Customer, error)
	GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error)
	InsertAdmin(ctx context.Context, header SBeEntity.InsertAdmin) (string, error)
	CheckUser(ctx context.Context, header SBeEntity.B_ChekUser) (interface{}, SBeEntity.B_Role, error)

	GetCartByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error)
	InsertHeaderCart(ctx context.Context, header SBeEntity.InsertHeaderCart) (string, error)

	InsertDetailCart(ctx context.Context, header SBeEntity.InsertDetailCart) (string, error)

	GetAllEmployee(ctx context.Context) ([]SBeEntity.T_Employee, error)
	GetEmpLastData(ctx context.Context) (SBeEntity.T_Employee, error)
	GetEmpByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Employee, error)
	InsertEmployee(ctx context.Context, header SBeEntity.InsertEmployee) (string, error)

	GetAllHeaderTran(ctx context.Context) ([]SBeEntity.TH_Transaction, error)
	GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error)

	UpdateCustomerById(ctx context.Context, header SBeEntity.UpdateCustomerById, cusId string) (string, error)



	GetAllOrder(ctx context.Context) ([]SBeEntity.T_Order, error) 
	InsertOrder(ctx context.Context, header SBeEntity.InsertOrder) (string, error)




	GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error)

	TokenUser(ctx context.Context) error

	InsertJoinHeaderDetailCart(ctx context.Context, header SBeEntity.InsertJoinHeaderDetailCart) (interface{}, error) 
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
