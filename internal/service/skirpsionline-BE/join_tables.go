package skirpsionlineBE


import (
	"context"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "log"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

func (s Service) GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error) {
	headers, err := s.skirpsionlineBE.GetJoinAdmCust(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinAdmCust]")
	}

	return headers, err
}


func (s Service) GetJoinOrdCustTHTra(ctx context.Context) ([]SBeEntity.JoinOrdCustTHTra, error) {
	headers, err := s.skirpsionlineBE.GetJoinOrdCustTHTra(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinOrdCustTHTra]")
	}

	return headers, err
}

func (s Service) GetJoinOrdCustTHTraByOrdId(ctx context.Context, ordId int) ([]SBeEntity.JoinOrdCustTHTra, error) {
	headers, err := s.skirpsionlineBE.GetJoinOrdCustTHTraByOrdId(ctx, ordId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinOrdCustTHTraByOrdId]")
	}

	return headers, err
}


func (s Service) GetJoinTDTraProdByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTraProdByTraId, error) {
	headers, err := s.skirpsionlineBE.GetJoinTDTraProdByTraId(ctx, traId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinTDTraProdByTraId]")
	}

	return headers, err
}


func (s Service) GetListJoinTHTDCartProdByCustIdAndCartId(ctx context.Context, custId string, cartId string) ([]SBeEntity.JoinTHTDCartProd, error) {
	headers, err := s.skirpsionlineBE.GetListJoinTHTDCartProdByCustIdAndCartId(ctx, custId, cartId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetListJoinTHTDCartProdByCustIdAndCartId]")
	}

	return headers, err
}
