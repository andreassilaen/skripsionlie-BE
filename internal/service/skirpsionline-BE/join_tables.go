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

func (s Service) GetJoinAdmCust(ctx context.Context) ([]SBeEntity.JoinAdmCust, error) {
	headers, err := s.skirpsionlineBE.GetJoinAdmCust(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetJoinAdmCust]")
	}

	return headers, err
}