package skirpsionlineBE

import (
	"context"
	// "strconv"
	// "fmt"
	"log"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

func (s Service) GetCustByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Customer, error) {
	headers, err := s.skirpsionlineBE.GetCustByLogin(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetMasterPack]")
	}

	return headers, err
}

func (s Service) InsertCustomer(ctx context.Context, newId string, header SBeEntity.InsertCustomer) (string, error) {
	var (
		result string
		err    error
	)

	// total, err := s.skirpsionlineBE.GetCountCust(ctx)

	// word := "adm001"

    // // Take left three characters
    // leftThree := word[5:]

    // // Print the result
    // fmt.Println(leftThree)

    // test, _ = strconv.Atoi(leftThree)


	result, err = s.skirpsionlineBE.InsertCustomer(ctx, header.InsertCustomerBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertProduct]")
	} else {
		result = "Sukses InserCustomer"
	}
	return result, err

	// _, err = s.skirpsionlineBE.GetAllCategory(ctx)
	
}