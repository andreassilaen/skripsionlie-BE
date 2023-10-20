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

	row, err := (*d.stmt)[getCartByCustId].QueryxContext(ctx,custId)

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
















