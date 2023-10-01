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

func (s Service) GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error) {
	headers, err := s.skirpsionlineBE.GetAllProduct(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllProduct]")
	}

	return headers, err
}
