package skirpsionlineBE

import (
	"context"
	"log"
	"strconv"

	// "strconv"
	// "strings"

	// "github.com/jmoiron/sqlx"
	// "github.com/opentracing/opentracing-go"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// jaegerLog "skripsi-online-BE/pkg/log"
	"skripsi-online-BE/pkg/errors"
)

func (d Data) GetAllCart(ctx context.Context) ([]SBeEntity.TH_Cart, error) {
	var (
		header  SBeEntity.TH_Cart
		headers []SBeEntity.TH_Cart
	)

	row, err := (*d.stmt)[getAllCart].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllCart][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllCart][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetCartByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error) {
	var (
		header  SBeEntity.TH_Cart
		headers []SBeEntity.TH_Cart
	)

	row, err := (*d.stmt)[getCartByCustId].QueryxContext(ctx, custId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetCartByCustId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetCartByCustId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetHeaderCartLastData(ctx context.Context) (SBeEntity.TH_Cart, error) {
	var (
		header SBeEntity.TH_Cart
		// headers []SBeEntity.T_Customer
	)

	row, err := (*d.stmt)[getHeaderCartLastData].QueryxContext(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[DATA][getHeaderCartLastData][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return header, errors.Wrap(err, "[DATA][getHeaderCartLastData][Query]")
		}
		// headers = append(headers, header)
	}
	log.Println("Data GetProdLastData : ", header)
	defer row.Close()
	return header, nil
}

func (d Data) InsertHeaderCart(ctx context.Context, header SBeEntity.TH_Cart2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertHeaderCart].ExecContext(ctx,
		header.CustId,
		header.CartTotal,
	)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][insertHeaderCart]")
	}
	result = " Sukses Data"
	return result, err
}

func (d Data) GetHeaderCartNotPayedCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error) {
	var (
		header  SBeEntity.TH_Cart
		headers []SBeEntity.TH_Cart
	)

	row, err := (*d.stmt)[getHeaderCartNotPayedByCustId].QueryxContext(ctx, custId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getHeaderCartNotPayedByCustId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getHeaderCartNotPayedByCustId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Header Cart : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) UpdateHeaderCartPeyed(ctx context.Context, cartId int) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateHeaderCartPayed].ExecContext(ctx, cartId)

	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[DATA][updateHeaderCartPayed]")
	}
	result = "Sukses Update " + strconv.Itoa(cartId)

	return result, err
}
