package models

import (
	"net/http"

	"project-2-Herwindams24/db"
)
type DaftarInstansi struct {
	Id 				int 		`json:"id"`
	namaInstansi	string		`json:"namaInstansi" validate:"required"`
	Kota 			string 		`json:"id_Kota" validate:"required"`
	Telfon			string		`json:"telfon" validate:"required"`
}

func FetchAllDaftarInstansi() (Response, error) {
	var obj DaftarInstansi
	var arrobj []DaftarInstansi
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM daftarInstansi"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.namaInstansi, &obj.Kota, &obj.Telfon)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}