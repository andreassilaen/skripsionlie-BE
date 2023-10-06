package skirpsionlineBE

import (
	"context"
	// "fmt"
	"strconv"

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

func (s Service) GetCustLastData(ctx context.Context) (SBeEntity.T_Customer, error) {
	header, err := s.skirpsionlineBE.GetCustLastData(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetCustLastData]")
	}

	return header, err
}


func (s Service) InsertCustomer(ctx context.Context, header SBeEntity.InsertCustomer) (string, error) {
	var (
		result string
		err    error
	)

	// total, err := s.skirpsionlineBE.GetCountCust(ctx)

	last, err:= s.skirpsionlineBE.GetCustLastData(ctx)

	

	// word := "adm001"

    // Take left three characters
    leftThree := last.CustId[4:]

    // Print the result
    log.Println("leftThree",leftThree)

    test, _ := strconv.Atoi(leftThree)
	log.Println("test = ", test)

	test +=1
	log.Println("test =>", test)
	testing := strconv.Itoa(test)
	header.InsertCustomerBody.CustId = "cust" + testing
	log.Println("hasil akhir : ",header.InsertCustomerBody.CustId)

	result, err = s.skirpsionlineBE.InsertCustomer(ctx, header.InsertCustomerBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertCustomer]")
	} else {
		result = "Sukses InsertCustomer = " + header.InsertCustomerBody.CustId
	}
	return result, err

	// _, err = s.skirpsionlineBE.GetAllCategory(ctx)
	
}