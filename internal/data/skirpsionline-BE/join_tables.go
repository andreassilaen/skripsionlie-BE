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