package skirpsionlineBE

// func (d Data) GetMasterJadwal(ctx context.Context, ptid int, code string, sttktype string) ([]sttkEntity.GetMasterJadwal, error) {
// 	var (
// 		jadwal  sttkEntity.GetMasterJadwal
// 		jadwals []sttkEntity.GetMasterJadwal
// 		err     error
// 	)
// 	log.Println("data masterjadwal object", ptid, code, sttktype)
// 	rows, err := (*d.stmt)[getMasterJadwal].QueryxContext(ctx, ptid, code, sttktype)
// 	if err != nil {
// 		return jadwals, errors.Wrap(err, "[DATA] [GetMasterJadwal]")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.StructScan(&jadwal); err != nil {
// 			return jadwals, errors.Wrap(err, "[DATA] [GetMasterJadwal]")
// 		}
// 		jadwals = append(jadwals, jadwal)
// 	}
// 	return jadwals, err
// }
