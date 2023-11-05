package skirpsionlineBE

import (
	"context"
	"log"

	// "github.com/jmoiron/sqlx"
	// "github.com/opentracing/opentracing-go"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	// jaegerLog "skripsi-online-BE/pkg/log"
	"skripsi-online-BE/pkg/errors"
)

func (d Data) GetAllOrder(ctx context.Context) ([]SBeEntity.T_Order, error) {
	var (
		header  SBeEntity.T_Order
		headers []SBeEntity.T_Order
	)

	row, err := (*d.stmt)[getAllOrder].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getAllOrder][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getAllOrder][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Order: ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) InsertOrder(ctx context.Context, header SBeEntity.T_Order2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertOrder].ExecContext(ctx,
		header.AdmId,
		header.TraId,
		// header.ProdLastupdate,
	)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][insertProduct]")
	}
	result = " Sukses Data"
	return result, err
}

func (d Data) InsertOrderAcc(ctx context.Context, header SBeEntity.T_Order2) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertOrderAcc].ExecContext(ctx,
		header.AdmId,
		header.TraId,
		// header.ProdLastupdate,
	)

	if err != nil {
		result = "Gagal insertOrderAcc"
		return result, errors.Wrap(err, "[DATA][insertOrderAcc]")
	}
	result = " Sukses insertOrderAcc"
	return result, err
}

func (d Data) UpdateOrderOnDeliveryYes(ctx context.Context, ordId int) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateOrderOnDeliveryYes].ExecContext(ctx, ordId)

	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[DATA][updateOrderOnDeliveryYes]")
	}
	result = "Sukses Update updateOrderOnDeliveryYes"

	return result, err
}
