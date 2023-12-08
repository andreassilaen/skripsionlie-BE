package skirpsionlineBE

import (
	"context"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "log"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

func (s Service) GetAllRekening(ctx context.Context) ([]SBeEntity.T_Rekening, error) {
	headers, err := s.skirpsionlineBE.GetAllRekening(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllRekening]")
	}

	return headers, err
}


func (s Service) GetRekByRekId(ctx context.Context, rekId int) ([]SBeEntity.T_Rekening, error) {
	headers, err := s.skirpsionlineBE.GetRekByRekId(ctx, rekId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetRekByRekId]")
	}

	return headers, err
}



func (s Service) InsertRekening(ctx context.Context, rekBank string, rekNumber int, rekName string) (string, error) {
	headers, err := s.skirpsionlineBE.InsertRekening(ctx, rekBank, rekNumber, rekName)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][InsertRekening]")
	}

	return headers, err
}