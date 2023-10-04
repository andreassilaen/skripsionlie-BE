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

func (d Data) GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error) {
	var (
		header  SBeEntity.T_Product
		headers []SBeEntity.T_Product
	)

	row, err := (*d.stmt)[getAllProduct].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllProduct][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllProduct][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}




func (d Data) InsertProduct(ctx context.Context, header SBeEntity.T_Product) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertProduct].ExecContext(ctx,
		header.ProdId,
		header.CtgId,
		header.ProdName,
		header.ProdStock,
		header.ProdPrice,
		header.ProdDesc,)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][insertProduct]")
	}
	result = " Sukses Data"
	return result, err
}

