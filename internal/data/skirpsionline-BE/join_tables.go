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



func (d Data) GetJoinOrdCustTHTra(ctx context.Context,) ([]SBeEntity.JoinOrdCustTHTra, error) {
	var (
		header  SBeEntity.JoinOrdCustTHTra
		headers []SBeEntity.JoinOrdCustTHTra
	)

	row, err := (*d.stmt)[getJoinOrdCustTHTra].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getJoinOrdCustTHTra][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getJoinOrdCustTHTra][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master JoinOrdCustTHTra  : ", headers)

	defer row.Close()
	return headers, nil
}


func (d Data) GetJoinOrdCustTHTraByOrdId(ctx context.Context, ordId int) ([]SBeEntity.JoinOrdCustTHTra, error) {
	var (
		header  SBeEntity.JoinOrdCustTHTra
		headers []SBeEntity.JoinOrdCustTHTra
	)

	row, err := (*d.stmt)[getJoinOrdCustTHTraByOrdId].QueryxContext(ctx, ordId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetJoinOrdCustTHTraByOrdId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetJoinOrdCustTHTraByOrdId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master JoinOrdCustTHTraByOrdId  : ", headers)

	defer row.Close()
	return headers, nil
}


func (d Data) GetJoinTDTraProdByTraId(ctx context.Context, traId string) ([]SBeEntity.JoinTDTraProdByTraId, error) {
	var (
		header  SBeEntity.JoinTDTraProdByTraId
		headers []SBeEntity.JoinTDTraProdByTraId
	)

	row, err := (*d.stmt)[getJoinTDTraProdByTraId].QueryxContext(ctx, traId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getJoinTDTraProdByTraId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getJoinTDTraProdByTraId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master getJoinTDTraProdByTraId  : ", headers)

	defer row.Close()
	return headers, nil
}