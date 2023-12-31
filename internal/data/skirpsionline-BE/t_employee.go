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

func (d Data) GetAllEmployee(ctx context.Context) ([]SBeEntity.T_Employee, error) {
	var (
		header  SBeEntity.T_Employee
		headers []SBeEntity.T_Employee
	)

	row, err := (*d.stmt)[getAllEmployee].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllEmployee][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllEmployee][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Employee : ", headers)

	defer row.Close()
	return headers, nil
}



func (d Data) GetEmpLastData(ctx context.Context) (SBeEntity.T_Employee, error) {
	var (
		header  SBeEntity.T_Employee
		// headers []SBeEntity.T_Customer
	)

	row, err := (*d.stmt)[getEmpLastData].QueryxContext(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[DATA][GetEmpLastData][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return header, errors.Wrap(err, "[DATA][GetEmpLastData][Query]")
		}
		// headers = append(headers, header)
	}
	log.Println("Data GetEmpLastData : ", header)
	defer row.Close()
	return header, nil
}

func (d Data) GetEmpByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Employee, error) {
	var (
		header  SBeEntity.T_Employee
		headers []SBeEntity.T_Employee
	)

	row, err := (*d.stmt)[getEmpByLogin].QueryxContext(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetEmpByLogin][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetEmpByLogin][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Employee : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) InsertEmployee(ctx context.Context, header SBeEntity.T_Employee) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertEmployee].ExecContext(ctx,
		header.EmpName,
		header.EmpUserName,
		header.EmpPassWord,
		header.EmpPhone,
		header.EmpEmail,
		header.EmpAddress,
	)

	if err != nil {
		result = "Gagal Insert Data"
		return result, errors.Wrap(err, "[DATA][insertEmployee]")
	}
	result = " Data Sukses"
	return result, err
}



func (d Data) UpdateEmployeeById(ctx context.Context, header SBeEntity.T_Employee2, empId string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateEmployeeById].ExecContext(ctx,
		header.EmpName,
		header.EmpUserName,
		header.EmpPassWord,
		header.EmpPhone,
		header.EmpEmail,
		header.EmpAddress,
		empId)

	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[DATA][UpdateCustomerById]")
	}
	result = "Sukses Update " + empId

	return result, err
}



func (d Data) GetEmpById(ctx context.Context, userId string) ([]SBeEntity.T_Employee, error) {
	var (
		header  SBeEntity.T_Employee
		headers []SBeEntity.T_Employee
	)

	row, err := (*d.stmt)[getEmpById].QueryxContext(ctx, userId)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][getEmpById][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][getEmpById][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Employee : ", headers)

	defer row.Close()
	return headers, nil
}