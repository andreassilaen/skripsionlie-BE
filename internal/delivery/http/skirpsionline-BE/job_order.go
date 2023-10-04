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
	GetAllOrder(ctx context.Context) ([]SBeEntity.TH_Order, error)
	GetCustByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Customer, error)
	GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error)
	InsertCustomer(ctx context.Context, header SBeEntity.InsertCustomer) (string, error) 

	


	GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error)

	TokenUser(ctx context.Context) error
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
