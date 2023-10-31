package skirpsionlineBE

import (
	"context"
	"log"

	// "github.com/jmoiron/sqlx"
	// "github.com/opentracing/opentracing-go"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// jaegerLog "skripsi-online-BE/pkg/log"
	"skripsi-online-BE/pkg/errors"
)

func (d Data) GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error) {
	var (
		header  SBeEntity.JoinAdmCust
		headers []SBeEntity.JoinAdmCust
	)

	row, err := (*d.stmt)[getJoinAdmCust].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetJoinAdmCust][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetJoinAdmCust][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetJoinOrdCustTHTra(ctx context.Context) ([]SBeEntity.JoinOrdCustTHTra, error) {
	var (
		header  SBeEntity.JoinOrdCustTHTra
		headers []SBeEntity.JoinOrdCustTHTra
	)

	row, err := (*d.stmt)[getJoinOrdCustTHTra].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getJoinOrdCustTHTra][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getJoinOrdCustTHTra][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master JoinOrdCustTHTra  : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetJoinOrdCustTHTraByOrdId(ctx context.Context, ordId int) ([]SBeEntity.JoinOrdCustTHTra, error) {
	var (
		header  SBeEntity.JoinOrdCustTHTra
		headers []SBeEntity.JoinOrdCustTHTra
	)

	row, err := (*d.stmt)[getJoinOrdCustTHTraByOrdId].QueryxContext(ctx, ordId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetJoinOrdCustTHTraByOrdId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetJoinOrdCustTHTraByOrdId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master JoinOrdCustTHTraByOrdId  : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetJoinTDTraProdByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTraProdByTraId, error) {
	var (
		header  SBeEntity.JoinTDTraProdByTraId
		headers []SBeEntity.JoinTDTraProdByTraId
	)

	row, err := (*d.stmt)[getJoinTDTraProdByTraId].QueryxContext(ctx, traId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getJoinTDTraProdByTraId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getJoinTDTraProdByTraId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master getJoinTDTraProdByTraId  : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetListJoinTHTDCartProdByCustIdAndCartId(ctx context.Context, custId string, cartId string) ([]SBeEntity.JoinTHTDCartProd, error) {
	var (
		header  SBeEntity.JoinTHTDCartProd
		headers []SBeEntity.JoinTHTDCartProd
	)

	row, err := (*d.stmt)[getListJoinTHTDCartProdByCustIdAndCartId].QueryxContext(ctx, custId, cartId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getListJoinTHTDCartProdByCustIdAndCartId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getListJoinTHTDCartProdByCustIdAndCartId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master JoinTHTDCartProdByCustIdAndCartId : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetProductInJOinTHTDCartProdByProdId(ctx context.Context, custId string, cartId string, prodId string) ([]SBeEntity.JoinTHTDCartProd, error) {
	var (
		header  SBeEntity.JoinTHTDCartProd
		headers []SBeEntity.JoinTHTDCartProd
	)

	row, err := (*d.stmt)[getProductInJoinTHTDCartProdByProdId].QueryxContext(ctx, custId, cartId, prodId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getProductInJoinTHTDCartProdByProdId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getProductInJoinTHTDCartProdByProdId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master GetProductInJOinTHTDCartProdByProdId : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) UpdateQtyDetailJoinTHTDCart(ctx context.Context, header SBeEntity.JoinTHTDCartProd2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateQtyDetailJoinTHTDCart].ExecContext(ctx,
		header.CartDtlQty,
		header.CustId,
		header.CartId,
		header.ProdId,
	)

	println("ceking data CardtlQty => ", header.CartDtlQty)

	if err != nil {
		result = "Gagal Update updateQtyDetailJoinTHTDCart"
		return result, errors.Wrap(err, "[DATA][updateQtyDetailJoinTHTDCart]")
	}
	result = "Sukses Update updateQtyDetailJoinTHTDCart"

	return result, err
}
