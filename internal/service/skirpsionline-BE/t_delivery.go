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

func (s Service) GetAllDelivery(ctx context.Context) ([]SBeEntity.T_Delivery, error) {
	headers, err := s.skirpsionlineBE.GetAllDelivery(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllDelivery]")
	}

	return headers, err
}

func (s Service) GetDeliverByEmpId(ctx context.Context, empId string) ([]SBeEntity.T_Delivery, error) {
	headers, err := s.skirpsionlineBE.GetDeliverByEmpId(ctx, empId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetDeliverByEmpId]")
	}

	return headers, err
}
