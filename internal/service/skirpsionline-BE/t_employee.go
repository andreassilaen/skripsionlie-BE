package skirpsionlineBE

import (
	"context"
	"log"
	"strconv"
	// "fmt"
	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"
	"skripsi-online-BE/pkg/errors"
	// "log"
	// "strconv"
	// "strings"
	// "job-order-be/pkg/errors"
	// "log"
)

func (s Service) GetAllEmployee(ctx context.Context) ([]SBeEntity.T_Employee, error) {
	headers, err := s.skirpsionlineBE.GetAllEmployee(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetAllEmployee]")
	}

	return headers, err
}

func (s Service) GetEmpLastData(ctx context.Context) (SBeEntity.T_Employee, error) {
	header, err := s.skirpsionlineBE.GetEmpLastData(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][GetEmpLastData]")
	}

	return header, err
}



func (s Service) GetEmpByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Employee, error) {
	headers, err := s.skirpsionlineBE.GetEmpByLogin(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetEmpByLogin]")
	}

	return headers, err
}


func (s Service) InsertEmployee(ctx context.Context, header SBeEntity.InsertEmployee) (string, error) {
	var (
		result string
		err    error
	)

	last, err:= s.skirpsionlineBE.GetEmpLastData(ctx)
	
    leftFour:= last.EmpId[3:]
    log.Println("leftFour = ",leftFour)

    test, _ := strconv.Atoi(leftFour)
	log.Println("test = ", test)

	test +=1
	log.Println("test =>", test)
	testing := strconv.Itoa(test)
	header.InsertEmployeeBody.EmpId = "emp" + testing
	log.Println("hasil akhir : ",header.InsertEmployeeBody.EmpId)

	result, err = s.skirpsionlineBE.InsertEmployee(ctx, header.InsertEmployeeBody)
	log.Println("header Service = ", header)
	if err != nil {
		result = "Gagal Insert"
		return result, errors.Wrap(err, "[Service][InsertEmployee]")
	} else {
		result = "Sukses InsertEmployee = " + header.InsertEmployeeBody.EmpId
	}
	return result, err
	
}