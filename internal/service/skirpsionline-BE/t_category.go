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

func (s Service) GetAllCategory(ctx context.Context) ([]SBeEntity.T_Category, error) {
	headers, err := s.skirpsionlineBE.GetAllCategory(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllCategory]")
	}

	return headers, err
}


func (s Service) InsertCategory(ctx context.Context, ctgType string) (string, error) {
	headers, err := s.skirpsionlineBE.InsertCategory(ctx, ctgType)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][InsertCategory]")
	}

	return headers, err
}