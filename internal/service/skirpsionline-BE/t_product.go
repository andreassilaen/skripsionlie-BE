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

func (s Service) GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error) {
	headers, err := s.skirpsionlineBE.GetAllProduct(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllProduct]")
	}

	return headers, err
}


func (s Service) InsertProduct(ctx context.Context, header SBeEntity.InsertProduct) (string, error) {
	var (
		result string
		err    error
	)


	result, err = s.skirpsionlineBE.InsertProduct(ctx, header.InsertProductBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertProduct]")
	} else {
		result = "Sukses InsertProduct"
	}
	return result, err

	// _, err = s.skirpsionlineBE.GetAllCategory(ctx)
	
}
