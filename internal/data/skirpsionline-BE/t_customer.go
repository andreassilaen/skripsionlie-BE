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

func (d Data) GetCustByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Customer, error) {
	var (
		header  SBeEntity.T_Customer
		headers []SBeEntity.T_Customer
	)

	row, err := (*d.stmt)[getCustByLogin].QueryxContext(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetCustByLogin][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetCustByLogin][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Customer : ", headers)

	defer row.Close()
	return headers, nil
}

// func (d Data) GetCustLastData(ctx context.Context) {int, error} {

// }


func (d Data) GetCountCust(ctx context.Context) (int, error) {
	var (
		// header  SBeEntity.T_Customer
		// headers []SBeEntity.T_Customer
		result	int
	)

	row, err := (*d.stmt)[getCountCust].QueryxContext(ctx)

	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetCustByLogin][Query]")
	}

	// for row.Next() {
	// 	err = row.StructScan(&header)
	// 	if err != nil {
	// 		return headers, errors.Wrap(err, "[DATA][GetCustByLogin][Query]")
	// 	}
	// 	headers = append(headers, header)
	// }
	log.Println("Banyak Customer : ", result)

	defer row.Close()
	return result, nil
}



func (d Data) InsertCustomer(ctx context.Context, header SBeEntity.T_Customer) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertCustomer].ExecContext(ctx,
		header.CustId,
		header.CustName,
		header.CustUserName,
		header.CustPassWord,
		header.CustPhone,
		header.CustEmail,
		header.CustAddress,
	)

	if err != nil {
		result = "Gagal Insert Data"
		return result, errors.Wrap(err, "[DATA][insertProduct]")
	}
	result = " Data Sukses"
	return result, err
}

