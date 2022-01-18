package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type data_karyawan struct {
	Id     string
	Nama   string
	Umur   int
	Posisi string
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/tugas16_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sql_tampil() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	// var kelas = 2
	rows, err := db.Query("Select * from tb_karyawan")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []data_karyawan

	for rows.Next() {
		var each = data_karyawan{}

		var err = rows.Scan(&each.Id, &each.Nama, &each.Umur, &each.Posisi)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each.Id, each.Nama, each.Umur, each.Posisi)
	}

}

/* menambah data
func sql_tambah() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("insert into tb_karyawan values (?,?,?,?)", "K5", "Irene", 19, "HRD")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("tambah data berhasil")
}
*/

/* mengubah data
func ubah_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("update tb_karyawan set umur = ? where Id = ?", 20, "K3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Ubah Berhasil")
}
*/

/* hapus data
func hapus_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("delete from tb_karyawan where id =? ", K5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("hapus Berhasil")
}
*/

func main() {
	// sql_tambah()
	// ubah_data()
	// hapus_data()
	sql_tampil()
}
