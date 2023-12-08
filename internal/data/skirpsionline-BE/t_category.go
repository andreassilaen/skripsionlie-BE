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

func (d Data) GetAllCategory(ctx context.Context) ([]SBeEntity.T_Category, error) {
	var (
		header  SBeEntity.T_Category
		headers []SBeEntity.T_Category
	)

	row, err := (*d.stmt)[getAllCategory].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllCategory][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllCategory][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}


func (d Data) InsertCategory(ctx context.Context, ctgType string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertCategory].ExecContext(ctx,
		ctgType,
	)

	if err != nil {
		result = "Gagal Insert Data"
		return result, errors.Wrap(err, "[DATA][insertCategory]")
	}
	result = " Data Sukses"
	return result, err
}