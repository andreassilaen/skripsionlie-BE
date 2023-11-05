package skirpsionlineBE

import (
	"context"
	// "fmt"
	"log"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
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

func (s Service) InsertDeliveryProcess(ctx context.Context, header SBeEntity.InsertDelivery) (string, error) {
	var (
		result string
		err    error
	)

	result, err = s.skirpsionlineBE.InsertDeliveryProcess(ctx, header.InsertDeliveryBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertDeliveryProcess]")
	} else {
		// result = "Sukses InsertCustomer = " + header.InsertCustomerBody.CustId
		result = "Sukses InsertCustomer"
	}
	return result, err

	// _, err = s.skirpsionlineBE.GetAllCategory(ctx)

}
