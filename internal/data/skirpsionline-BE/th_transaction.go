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

// func (d Data) GetAllCart(ctx context.Context) ([]SBeEntity.TH_Cart, error) {
// 	var (
// 		header  SBeEntity.TH_Cart
// 		headers []SBeEntity.TH_Cart
// 	)

// 	row, err := (*d.stmt)[getAllCart].QueryxContext(ctx)

// 	if err != nil {
// 		return headers, errors.Wrap(err, "[DATA][GetAllCart][Query]")
// 	}

// 	for row.Next() {
// 		err = row.StructScan(&header)
// 		if err != nil {
// 			return headers, errors.Wrap(err, "[DATA][GetAllCart][Query]")
// 		}
// 		headers = append(headers, header)
// 	}
// 	log.Println("Master Product : ", headers)

// 	defer row.Close()
// 	return headers, nil
// }

func (d Data) GetAllHeaderTran(ctx context.Context) ([]SBeEntity.TH_Transaction, error) {
	var (
		header  SBeEntity.TH_Transaction
		headers []SBeEntity.TH_Transaction
	)

	row, err := (*d.stmt)[getAllHeaderTran].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getAllHeaderTran][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getAllHeaderTran][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error) {
	var (
		header  SBeEntity.TH_Transaction
		headers []SBeEntity.TH_Transaction
	)

	row, err := (*d.stmt)[getTranByCartId].QueryxContext(ctx, cartId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetTranByCartId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetTranByCartId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Transaction : ", headers)

	defer row.Close()
	return headers, nil
}


func (d Data) InsertHeaderTran(ctx context.Context, header SBeEntity.TH_Transaction2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertHeaderTran].ExecContext(ctx,
		header.CartId,
		header.CustId,
		header.RekId,
		header.TranTotal,
		header.TranImg,
	)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][insertHeaderTran]")
	}
	result = " Sukses Data"
	return result, err
}

