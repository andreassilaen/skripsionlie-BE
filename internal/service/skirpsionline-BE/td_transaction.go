package skirpsionlineBE

import (
	"context"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)




func (s Service) GetDetailTranByTraId(ctx context.Context, traId string) ([]SBeEntity.TD_Transaction, error) {
	details, err := s.skirpsionlineBE.GetDetailTranByTraId(ctx, traId)

	if err != nil {
		return details, errors.Wrap(err, "[SERVICE][GetDetailTranByTraId]")
	}

	return details, err
}