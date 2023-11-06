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

func (d Data) GetAllDelivery(ctx context.Context) ([]SBeEntity.T_Delivery, error) {
	var (
		header  SBeEntity.T_Delivery
		headers []SBeEntity.T_Delivery
	)

	row, err := (*d.stmt)[getAllDelivery].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getAllDelivery][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getAllDelivery][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Delivery : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) GetDeliverByEmpId(ctx context.Context, empId string) ([]SBeEntity.T_Delivery, error) {
	var (
		header  SBeEntity.T_Delivery
		headers []SBeEntity.T_Delivery
	)

	row, err := (*d.stmt)[getDeliveryByEmpId].QueryxContext(ctx, empId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetDeliverByEmpId][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetDeliverByEmpId][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Delivery : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) InsertDeliveryProcess(ctx context.Context, header SBeEntity.T_Delivery2) (string, error) {
	var (
		result string
		err    error
	)



	_, err = (*d.stmt)[insertDeliveryProcess].ExecContext(ctx,
		header.EmpId,
		header.OrdId,
	)

	if err != nil {
		result = "Gagal update Data"
		return result, errors.Wrap(err, "[DATA][insertDeliveryProcess]")
	}
	result = " Sukses Data"
	return result, err
}


func (d Data) UpdateDeliveryDone(ctx context.Context, ordId string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateDeliveryDone].ExecContext(ctx, ordId)

	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[DATA][updateDeliveryDone]")
	}
	result = "Sukses Update updateDeliveryDone"

	return result, err
}