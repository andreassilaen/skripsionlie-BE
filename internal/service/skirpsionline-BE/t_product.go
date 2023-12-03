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

func (s Service) GetProdById(ctx context.Context, prodId string) ([]SBeEntity.T_Product, error) {
	headers, err := s.skirpsionlineBE.GetProdById(ctx, prodId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetProdById]")
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

}

func (s Service) UpdateProdById(ctx context.Context, header SBeEntity.UpdateProdById, prodId string) (string, error) {
	var (
		result string
		err    error
	)

	result, err = s.skirpsionlineBE.UpdateProdById(ctx, header.UpdateProdByIdBody, prodId)
	log.Println("prodId =>", prodId)
	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[Service][UpdateProdById]")
	} else {
		result = "Sukses UpdateProdById"
	}
	return result, err
}

func (s Service) GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error) {
	header, err := s.skirpsionlineBE.GetProdLastData(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetProdLastData]")
	}

	return header, err
}

func (s Service) DeleteProductByProdId(ctx context.Context, prodId int) (string, error) {
	headers, err := s.skirpsionlineBE.DeleteProductByProdId(ctx, prodId)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][DeleteProductByProdId]")
	}

	return headers, err
}

// func (s Service) UpdateProdStockById(ctx context.Context, header SBeEntity.UpdateProdStockById, prodId int) (string, error) {
// 	var (
// 		result string
// 		err    error
// 	)

// 	result, err = s.skirpsionlineBE.UpdateProdStockById(ctx, header.ProdStockBody, prodId)
// 	log.Println("admId =>", prodId)
// 	if err != nil {
// 		result = "Gagal Update"
// 		return result, errors.Wrap(err, "[Service][UpdateProdStockById]")
// 	} else {
// 		result = "Sukses UpdateProdStockById"
// 	}
// 	return result, err
// }

func (s Service) UpdateProdStockById(ctx context.Context, prodStock int, prodId int) (string, error) {
	var (
		result string
		err    error
	)

	result, err = s.skirpsionlineBE.UpdateProdStockById(ctx, prodStock, prodId)
	log.Println("prodId =>", prodId)
	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[Service][UpdateProdStockById]")
	} else {
		result = "Sukses UpdateProdStockById"
	}
	return result, err
}
