package skirpsionlineBE

import (
	"context"
	"log"
	"strconv"

	// "github.com/jmoiron/sqlx"
	// "github.com/opentracing/opentracing-go"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// jaegerLog "skripsi-online-BE/pkg/log"
	"skripsi-online-BE/pkg/errors"
)

func (d Data) GetAllProduct(ctx context.Context) ([]SBeEntity.T_Product, error) {
	var (
		header  SBeEntity.T_Product
		headers []SBeEntity.T_Product
	)

	row, err := (*d.stmt)[getAllProduct].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllProduct][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllProduct][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetProdById(ctx context.Context, prodId string) ([]SBeEntity.T_Product, error) {
	var (
		header  SBeEntity.T_Product
		headers []SBeEntity.T_Product
	)

	row, err := (*d.stmt)[getProdById].QueryxContext(ctx, prodId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getProdById][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getProdById][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Product by Id : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetProdLastData(ctx context.Context) (SBeEntity.T_Product, error) {
	var (
		header SBeEntity.T_Product
		// headers []SBeEntity.T_Customer
	)

	row, err := (*d.stmt)[getProdLastData].QueryxContext(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[DATA][GetProdLastData][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return header, errors.Wrap(err, "[DATA][GetProdLastData][Query]")
		}
		// headers = append(headers, header)
	}
	log.Println("Data GetProdLastData : ", header)
	defer row.Close()
	return header, nil
}

func (d Data) InsertProduct(ctx context.Context, header SBeEntity.T_Product2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertProduct].ExecContext(ctx,
		header.AdmId,
		header.CtgId,
		header.ProdName,
		header.ProdDesc,
		header.ProdPrice,
		header.ProdStock,
		header.ProdImage,
		// header.ProdLastupdate,
	)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][insertProduct]")
	}
	result = " Sukses Data"
	return result, err
}

func (d Data) UpdateProdById(ctx context.Context, header SBeEntity.T_Product3, prodId string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateProdById].ExecContext(ctx,
		header.AdmId,
		header.ProdName,
		header.ProdDesc,
		header.ProdPrice,
		header.ProdStock,
		header.ProdImage,
		prodId,
	)

	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[DATA][updateProdById]")
	}
	result = "Sukses Update " + prodId

	return result, err
}

func (d Data) DeleteProductByProdId(ctx context.Context, prodId int) (string, error) {
	var (
		result string
		err    error
	)

	prod_id := strconv.Itoa(prodId)

	_, err = (*d.stmt)[deleteProductByProdId].ExecContext(ctx, prodId)
	if err != nil {
		result = "Gagal delete Data => " + prod_id
		return result, errors.Wrap(err, "[DATA][DeleteProductByProdId]")
	}
	result = " Sukses delete Data => " + prod_id
	return result, err
}

// func (d Data) UpdateProdStockById(ctx context.Context, header SBeEntity.ProdStock, prodId int) (string, error) {
// 	var (
// 		result string
// 		err    error
// 	)

// 	Id := strconv.Itoa(prodId)

// 	_, err = (*d.stmt)[updateProdStockById].ExecContext(ctx,
// 		header.ProdStock,
// 		prodId)

// 	if err != nil {
// 		result = "Gagal Update Stock"
// 		return result, errors.Wrap(err, "[DATA][UpdateProdStockById]")
// 	}
// 	result = "Sukses Update stock prod_id => " + Id

// 	return result, err
// }

func (d Data) UpdateProdStockById(ctx context.Context, prodStock int, prodId int) (string, error) {
	var (
		result string
		err    error
	)

	Id := strconv.Itoa(prodId)

	_, err = (*d.stmt)[updateProdStockById].ExecContext(ctx,
		prodStock,
		prodId)

	if err != nil {
		result = "Gagal Update Stock"
		return result, errors.Wrap(err, "[DATA][UpdateProdStockById]")
	}
	result = "Sukses Update stock prod_id => " + Id

	return result, err
}
