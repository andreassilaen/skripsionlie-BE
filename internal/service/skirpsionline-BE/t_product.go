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

func (s Service) GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error) {
	headers, err := s.skirpsionlineBE.GetAllProduct(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllProduct]")
	}

	return headers, err
}


func (s Service) InsertProduct(ctx context.Context, header SBeEntity.InsertProduct) (string, error) {
	var (
		result string
		err    error
	)
	
	result, err = s.skirpsionlineBE.InsertProduct(ctx, header.InsertProductBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertProduct]")
	} else {
		result = "Sukses InsertProduct"
	}
	return result, err

	// _, err = s.skirpsionlineBE.GetAllCategory(ctx)
	
}

// func (s Service) InsertProduct(ctx context.Context, header SBeEntity.InsertProduct) (string, error) {
// 	var (
// 		result string
// 		err    error
// 	)

// 	last, err:= s.skirpsionlineBE.GetProdLastData(ctx)

	

// 	// word := "adm001"

//     // Take left three characters
//     leftThree := last.ProdId[3:]

//     // Print the result
//     log.Println("leftThree",leftThree)

//     test, _ := strconv.Atoi(leftThree)
// 	log.Println("test = ", test)

// 	test +=1
// 	log.Println("test =>", test)
// 	testing := strconv.Itoa(test)
// 	header.InsertProductBody.ProdId = "pro" + testing
// 	log.Println("hasil akhir : ",header.InsertProductBody.ProdId)
	

// 	result, err = s.skirpsionlineBE.InsertProduct(ctx, header.InsertProductBody)
// 	log.Println("header Service = ", header)
// 	if err != nil {
// 		result = "Gagal Insert"
// 		return result, errors.Wrap(err, "[Service][InsertProduct]")
// 	} else {
// 		result = "Sukses InsertProduct"
// 	}
// 	return result, err

// 	// _, err = s.skirpsionlineBE.GetAllCategory(ctx)
	
// }


func (s Service) GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error) {
	header, err := s.skirpsionlineBE.GetProdLastData(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetProdLastData]")
	}

	return header, err
}