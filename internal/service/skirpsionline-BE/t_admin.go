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

func (s Service) GetAllAdmin(ctx context.Context) ([]SBeEntity.T_Admin, error) {
	headers, err := s.skirpsionlineBE.GetAllAdmin(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllAdmin]")
	}

	return headers, err
}


func (s Service) GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error) {
	headers, err := s.skirpsionlineBE.GetAdmByLogin(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAdmByLogin]")
	}

	return headers, err
}