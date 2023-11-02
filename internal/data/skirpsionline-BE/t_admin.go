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

func (d Data) GetAllAdmin(ctx context.Context) ([]SBeEntity.T_Admin, error) {
	var (
		header  SBeEntity.T_Admin
		headers []SBeEntity.T_Admin
	)

	row, err := (*d.stmt)[getAllAdmin].QueryxContext(ctx)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAllAdmin][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAllAdmin][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Admin : ", headers)

	defer row.Close()
	return headers, nil
}



func (d Data) GetAdmLastData(ctx context.Context) (SBeEntity.T_Admin, error) {
	var (
		header  SBeEntity.T_Admin
		// headers []SBeEntity.T_Customer
	)

	row, err := (*d.stmt)[getAdmLastData].QueryxContext(ctx)

	if err != nil {
		return header, errors.Wrap(err, "[DATA][GetAdmLastData][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return header, errors.Wrap(err, "[DATA][GetAdmLastData][Query]")
		}
		// headers = append(headers, header)
	}
	log.Println("Data GetAdmLastData : ", header)
	defer row.Close()
	return header, nil
}

func (d Data) GetAdmByLogin(ctx context.Context, username string, password string) ([]SBeEntity.T_Admin, error) {
	var (
		header  SBeEntity.T_Admin
		headers []SBeEntity.T_Admin
	)

	row, err := (*d.stmt)[getAdmByLogin].QueryxContext(ctx, username, password)

	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetAdmByLogin][Query]")
	}

	for row.Next() {
		err = row.StructScan(&header)
		if err != nil {
			return headers, errors.Wrap(err, "[DATA][GetAdmByLogin][Query]")
		}
		headers = append(headers, header)
	}
	log.Println("Master Admin : ", headers)

	defer row.Close()
	return headers, nil
}

func (d Data) InsertAdmin(ctx context.Context, header SBeEntity.T_Admin) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[insertAdmin].ExecContext(ctx,
		header.AdmName,
		header.AdmUserName,
		header.AdmPassWord,
		header.AdmPhone,
		header.AdmEmail,
		header.AdmAddress,
	)

	if err != nil {
		result = "Gagal Insert Data"
		return result, errors.Wrap(err, "[DATA][insertAdmin]")
	}
	result = " Data Sukses"
	return result, err
}


func (d Data) UpdateAdminById(ctx context.Context, header SBeEntity.T_Admin2, admId string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateAdminById].ExecContext(ctx,
		header.AdmName,
		header.AdmUserName,
		header.AdmPassWord,
		header.AdmPhone,
		header.AdmEmail,
		header.AdmAddress,
		admId)

	if err != nil {
		result = "Gagal Update"
		return result, errors.Wrap(err, "[DATA][updateAdminById]")
	}
	result = "Sukses Update adm " + admId

	return result, err
}