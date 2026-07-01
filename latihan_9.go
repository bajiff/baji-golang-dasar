// ! latihan_9.go

package main

import (
	"errors"
	"fmt"
)

// ? Latihan
func bagiDua(angka int) (int, error) {
	if angka == 0 {
		return 0, errors.New("Tidak bisa membagi 0")
	}
	return angka / 2, nil
}

// ? Real case
func validasiPesan(pesan string) (string, error) {
	if pesan == "" {
		return "", errors.New("Pesanan tidak boleh kosong, Koki bingung!")
	}
	wadahPesan := fmt.Sprintf("Memasak pesanan Anda: %s", pesan)

	return wadahPesan, nil
}

func validasiNama(namaAnda string) (string, error) {
	wadahNama := fmt.Sprintf("Halo : %s", namaAnda)

	if namaAnda == "" {
		return "", errors.New("Nama Anda tidak valid")
	}
	return wadahNama, nil
}

func latihanSembilan() {
	defer fmt.Println("Dapur ditutup!")
	fmt.Println("Welcome to Latihan 9")
	testBagiDua, err := bagiDua(0)

	if err != nil {
		fmt.Println("Waduh error: ", err.Error())
	} else {
		fmt.Println("Berhasil: ", testBagiDua)
	}

	hasilMenu, err := validasiPesan("")

	if err != nil {
		fmt.Println("Waduh error: ", err.Error())
	} else {
		fmt.Println(hasilMenu)
	}

	hasilMenu, err = validasiPesan("Nasi Goreng")
	if err != nil {
		fmt.Println("Waduh error: ", err.Error())
	} else {
		fmt.Println(hasilMenu)
	}
	identity, err := validasiNama("Meki")

	if err != nil {
		fmt.Println("Waduh error pak: ", err.Error())
	} else {
		fmt.Println(identity)
	}
}
