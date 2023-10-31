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

func (d Data) GetAllRekening(ctx context.Context) ([]SBeEntity.T_Rekening, error) {
	var (
		header  SBeEntity.T_Rekening
		headers []SBeEntity.T_Rekening
	)

	row, err := (*d.stmt)[getAllRekening].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getAllRekening][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getAllRekening][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Rekening : ", headers)

	defer row.Close()
	return headers, nil
}


func (d Data) GetRekByRekId(ctx context.Context, rekId int) ([]SBeEntity.T_Rekening, error) {
	var (
		header  SBeEntity.T_Rekening
		headers []SBeEntity.T_Rekening
	)

	row, err := (*d.stmt)[getRekByRekId].QueryxContext(ctx, rekId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetRekByRekId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetRekByRekId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Rekening by Id : ", headers)

	defer row.Close()
	return headers, nil
}