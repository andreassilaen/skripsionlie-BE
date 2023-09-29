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

func (d Data) GetAllAdmin(ctx context.Context) ([]SBeEntity.T_Admin, error) {
	var (
		header  SBeEntity.T_Admin
		headers []SBeEntity.T_Admin
	)

	row, err := (*d.stmt)[getAllAdmin].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllAdmin][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllAdmin][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Admin : ", headers)

	defer row.Close()
	return headers, nil
}



