package models

import (
	"net/http"

	"project-2-Herwindams24/db"
)

type DaftarInstansi struct {
	Id           int    `json:"id"`
	namaInstansi string `json:"namaInstansi" validate:"required"`
	Kota         string `json:"kota" validate:"required"`
	Telfon       string `json:"telfon" validate:"required"`
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

func StoreDaftarInstansi(namaInstansi string, Kota string, telfon string) (Response, error) {

	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT daftarInstansi (namaInstansi, kota, telfon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namaInstansi, Kota, telfon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateDaftarInstansi(id int, namaInstansi string, Kota string, telfon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE daftarInstansi SET namaInstansi = ?, Kota = ?, telfon = ? WHERE id = ?"


	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namaInstansi, Kota, telfon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil
}
