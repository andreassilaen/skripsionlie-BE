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

func (s Service) GetProductInJOinTHTDCartProdByProdId(ctx context.Context, custId string, cartId string, prodId string) ([]SBeEntity.JoinTHTDCartProd, error) {
	headers, err := s.skirpsionlineBE.GetProductInJOinTHTDCartProdByProdId(ctx, custId, cartId, prodId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetProductInJOinTHTDCartProdByProdId]")
	}

	return headers, err
}

func (s Service) UpdateQtyDetailJoinTHTDCart(ctx context.Context, header SBeEntity.UpdateQtyDetailJoinTHTDCartProd) (string, error) {
	var (
		result string
		err    error
	)

	result, err = s.skirpsionlineBE.UpdateQtyDetailJoinTHTDCart(ctx, header.UpdateQtyDetailJoinTHTDCartProdBody)
	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[Service][UpdateQtyDetailJoinTHTDCart]")
	} else {
		result = "Sukses UpdateQtyDetailJoinTHTDCart"
	}
	return result, err
}

func (s Service) GetJoinTHTraRekByCusId(ctx context.Context, custId string) ([]SBeEntity.JoinTHTraRek, error) {
	headers, err := s.skirpsionlineBE.GetJoinTHTraRekByCusId(ctx, custId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetCustById]")
	}

	return headers, err
}

func (s Service) GetJoinOrdTHTDTraProdByOrdId(ctx context.Context, ordId string) ([]SBeEntity.JoinOrdTHTDTraProdByOrdId, error) {
	headers, err := s.skirpsionlineBE.GetJoinOrdTHTDTraProdByOrdId(ctx, ordId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinOrdTHTDTraProdByOrdId]")
	}

	return headers, err
}

func (s Service) GetJoinOrdTHTraByCustId(ctx context.Context, custId string) ([]SBeEntity.JoinOrdTHTraByCustId, error) {
	headers, err := s.skirpsionlineBE.GetJoinOrdTHTraByCustId(ctx, custId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinOrdTHTraByCustId]")
	}

	return headers, err
}

func (s Service) GetJoinOrdTHTraDelByCustId(ctx context.Context, custId string) ([]SBeEntity.JoinOrdTHTraDelByCustId, error) {
	header, err := s.skirpsionlineBE.GetJoinOrdTHTraDelByCustId(ctx, custId)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetJoinOrdTHTraDelByCustId]")
	}

	return header, err
}

func (s Service) GetCountDashboardAdmin(ctx context.Context) (SBeEntity.CountTHTraOrdDel, error) {
	header, err := s.skirpsionlineBE.GetCountDashboardAdmin(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetCountDashboardAdmin]")
	}

	return header, err
}

func (s Service) GetReportOrdTHTraByOrdDate(ctx context.Context, startDate string, endDate string) ([]SBeEntity.JoinReportOrdTHTra, error) {
	header, err := s.skirpsionlineBE.GetReportOrdTHTraByOrdDate(ctx, startDate, endDate)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetReportOrdTHTraByOrdDate]")
	}

	return header, err
}

func (s Service) GetReportOrdTHTraDelByOrdDate(ctx context.Context, custId string, startDate string, endDate string) ([]SBeEntity.JoinReportOrdTHTraDel, error) {
	header, err := s.skirpsionlineBE.GetReportOrdTHTraDelByOrdDate(ctx, custId, startDate, endDate)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetReportOrdTHTraDelByOrdDate]")
	}

	return header, err
}

func (s Service) GetDetailReportByOrdId(ctx context.Context, ordId int) (SBeEntity.JoinDetailReport, error) {
	header, err := s.skirpsionlineBE.GetDetailReportByOrdId(ctx, ordId)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetDetailReportByOrdId]")
	}

	return header, err
}

func (s Service) GetJoinTDTranProdCustByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTranProdCustByTraId, error) {
	header, err := s.skirpsionlineBE.GetJoinTDTranProdCustByTraId(ctx, traId)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetJoinTDTranProdCustByTraId]")
	}

	return header, err
}
