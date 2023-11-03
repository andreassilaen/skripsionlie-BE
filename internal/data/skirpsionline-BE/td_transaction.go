package skirpsionlineBE

import (
	"context"
	"log"

	// "github.com/jmoiron/sqlx"
	// "github.com/opentracing/opentracing-go"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// jaegerLog "skripsi-online-BE/pkg/log"
	"skripsi-online-BE/pkg/errors"

	"github.com/jmoiron/sqlx"
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


func (d Data) InsertDetailTran(ctx context.Context, header SBeEntity.TD_Transaction2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertDetailTran].ExecContext(ctx,
		header.TraId,
		header.ProdId,
		header.TraDtlQty,
		header.TraDtlPrice,
		header.TraDtlAmount,
	)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][InsertDetailTran]")
	}
	result = " Sukses Data"
	return result, err
}



func (d Data) NewInsertDetailTran(ctx context.Context, user []SBeEntity.TD_Transaction2) error {

    for _, v := range user {
        query, args, err := sqlx.In(qInsertDetailTran,
            v.TraId, v.ProdId, v.TraDtlQty, v.TraDtlPrice, v.TraDtlAmount)
        if err != nil {
            return errors.Wrap(err, "[DATA][OtherInsertDetailTran]")
        }
        _, err = d.db.ExecContext(ctx, query, args...)
        if err != nil {
            return errors.Wrap(err, "[DATA][OtherInsertDetailTran]")
        }
    }
    return nil
}