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

func (s Service) GetAllOrder(ctx context.Context) ([]SBeEntity.T_Order, error) {
	headers, err := s.skirpsionlineBE.GetAllOrder(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllOrder]")
	}

	return headers, err
}






func (s Service) InsertOrder(ctx context.Context, header SBeEntity.InsertOrder) (string, error) {
	var (
		result string
		err    error
	)

	result, err = s.skirpsionlineBE.InsertOrder(ctx, header.InsertOrderBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertOrder]")
	} else {
		result = "Sukses InsertOrder"
	}
	return result, err
	
}










