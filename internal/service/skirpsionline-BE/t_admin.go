package skirpsionlineBE

import (
	"context"
	"log"
	// "strconv"
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

func (s Service) GetAdmLastData(ctx context.Context) (SBeEntity.T_Admin, error) {
	header, err := s.skirpsionlineBE.GetAdmLastData(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetAdmLastData]")
	}

	return header, err
}



func (s Service) GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error) {
	headers, err := s.skirpsionlineBE.GetAdmByLogin(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAdmByLogin]")
	}

	return headers, err
}


func (s Service) InsertAdmin(ctx context.Context, header SBeEntity.InsertAdmin) (string, error) {
	var (
		result string
		err    error
	)

	result, err = s.skirpsionlineBE.InsertAdmin(ctx, header.InsertAdminBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertAdmin]")
	} else {
		result = "Sukses InsertAdmin" 
	}
	return result, err
	
}