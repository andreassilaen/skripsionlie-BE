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
	InsertCustomer(ctx context.Context, header SBeEntity.T_Customer) (string, error)
	GetCountCust(ctx context.Context) (int, error)
	GetAdmLastData(ctx context.Context) (SBeEntity.T_Admin, error)
	GetCustLastData(ctx context.Context) (SBeEntity.T_Customer, error)
	GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error)
	GetHeaderCartLastData(ctx context.Context) (SBeEntity.TH_Cart, error)
	InsertAdmin(ctx context.Context, header SBeEntity.T_Admin) (string, error)

	GetCartByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error)
	InsertHeaderCart(ctx context.Context, header SBeEntity.TH_Cart2) (string, error)

	InsertDetailCart(ctx context.Context, header SBeEntity.TD_Cart2) (string, error)
	NewInsertDetailCart(ctx context.Context, user []SBeEntity.TD_Cart2) error

	GetAllEmployee(ctx context.Context) ([]SBeEntity.T_Employee, error)
	GetEmpLastData(ctx context.Context) (SBeEntity.T_Employee, error)
	GetEmpByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Employee, error)
	InsertEmployee(ctx context.Context, header SBeEntity.T_Employee) (string, error)

	GetAllHeaderTran(ctx context.Context) ([]SBeEntity.TH_Transaction, error)
	GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error)

	UpdateCustomerById(ctx context.Context, header SBeEntity.T_Customer2, cusId string) (string, error)

	GetAllOrder(ctx context.Context) ([]SBeEntity.T_Order, error)
	InsertOrder(ctx context.Context, header SBeEntity.T_Order2) (string, error)

	GetAllDelivery(ctx context.Context) ([]SBeEntity.T_Delivery, error)
	GetDeliverByEmpId(ctx context.Context, empId string) ([]SBeEntity.T_Delivery, error)

	GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error)
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
