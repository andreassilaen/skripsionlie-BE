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




func (d Data) GetDetailTranByTraId(ctx context.Context, traId string) ([]SBeEntity.TD_Transaction, error) {
	var (
		detail  SBeEntity.TD_Transaction
		details []SBeEntity.TD_Transaction
	)

	row, err := (*d.stmt)[getDetailTranByTraId].QueryxContext(ctx,traId)

	if err != nil {
		return details, errors.Wrap(err, "[DATA][getDetailTranByTraId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&detail)
		if err != nil {
			return details, errors.Wrap(err, "[DATA][getDetailTranByTraId][Query]")
		}
		details = append(details, detail)
	}
	log.Println("Detail Transaction : ", details)

	defer row.Close()
	return details, nil
}