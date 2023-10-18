package skirpsionlineBE

import (
	"context"
	// "fmt"
	// "strconv"

	// "strconv"
	// "fmt"
	// "log"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)






func (s Service) GetTranByCartId(ctx context.Context, cartId string) ([]SBeEntity.TH_Transaction, error) {
	headers, err := s.skirpsionlineBE.GetTranByCartId(ctx, cartId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetTranByCartId]")
	}

	return headers, err
}