package skirpsionlineBE

import (
	"context"
	// "fmt"
	// "strconv"

	// "strconv"
	// "fmt"
	"log"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

func (s Service) GetAllHeaderTran(ctx context.Context) ([]SBeEntity.TH_Transaction, error) {
	headers, err := s.skirpsionlineBE.GetAllHeaderTran(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllHeaderTran]")
	}

	return headers, err
}

func (s Service) GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error) {
	headers, err := s.skirpsionlineBE.GetTranByCartId(ctx, cartId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetTranByCartId]")
	}

	return headers, err
}



func (s Service) InsertHeaderTran(ctx context.Context, header SBeEntity.InsertHeaderTransaction) (string, error) {
	var (
		result string
		err    error
	)
	
	result, err = s.skirpsionlineBE.InsertHeaderTran(ctx, header.InsertHeaderTransactionBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertHeaderTran]")
	} else {
		result = "Sukses InsertHeaderTran"
	}
	return result, err
	
}

func (s Service) GetAllHeaderTranByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Transaction, error) {
	headers, err := s.skirpsionlineBE.GetAllHeaderTranByCustId(ctx, custId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllHeaderTranByCustId]")
	}

	return headers, err
}
