package skirpsionlineBE

import (
	"context"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	"log"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

// func (s Service) GetAllCart(ctx context.Context) ([]SBeEntity.TH_Cart, error) {
// 	headers, err := s.skirpsionlineBE.GetAllCart(ctx)

// 	if err != nil {
// 		return headers, errors.Wrap(err, "[SERVICE][GetAllCart]")
// 	}

// 	return headers, err
// }


// func (s Service) GetCartByCustId(ctx context.Context, custId string) ([]SBeEntity.TH_Cart, error) {
// 	headers, err := s.skirpsionlineBE.GetCartByCustId(ctx, custId)

// 	if err != nil {
// 		return headers, errors.Wrap(err, "[SERVICE][GetCartByCustId]")
// 	}

// 	return headers, err
// }

func (s Service) InsertDetailCart(ctx context.Context, header SBeEntity.InsertDetailCart) (string, error) {
	var (
		result string
		err    error
	)
	
	result, err = s.skirpsionlineBE.InsertDetailCart(ctx, header.InsertDetailCartBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertDetailCart]")
	} else {
		result = "Sukses InsertDetailCart"
	}
	return result, err
	
}






