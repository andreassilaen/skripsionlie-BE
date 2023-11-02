package skirpsionlineBE

import (
	"context"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	"log"
)




func (s Service) GetDetailTranByTraId(ctx context.Context, traId string) ([]SBeEntity.TD_Transaction, error) {
	details, err := s.skirpsionlineBE.GetDetailTranByTraId(ctx, traId)

	if err != nil {
		return details, errors.Wrap(err, "[SERVICE][GetDetailTranByTraId]")
	}

	return details, err
}




func (s Service) InsertDetailTransaction(ctx context.Context, header SBeEntity.InsertDetailTransaction) (string, error) {
	var (
		result string
		err    error
	)
	
	result, err = s.skirpsionlineBE.InsertDetailTran(ctx, header.InsertDetailTransactionBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertDetailTransaction]")
	} else {
		result = "Sukses InsertDetailTransaction"
	}
	return result, err
	
}