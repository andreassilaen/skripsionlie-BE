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

func (d Data) GetAllOrder(ctx context.Context) ([]SBeEntity.TH_Order, error) {
	var (
		header  SBeEntity.TH_Order
		headers []SBeEntity.TH_Order
	)

	row, err := (*d.stmt)[getAllOrder].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllOrder][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllOrder][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}