// ! latihan_7.go
package main

import "fmt"

type Pegawai struct {
	Nama      string
	Energi    int
	Jabatan   string
	Gaji      int
	IsActive  bool
	IsMarried bool
	Tunjangan int
	Usia      int
	Skills    []string
}

func kerjaLembur(p *Pegawai) int {
	p.Energi = p.Energi - 30
	fmt.Printf("Pegawai %s sisa Energi %d\n", p.Nama, p.Energi)
	return p.Energi
}

func latihanTujuh() {
	pegawaiBaru := Pegawai{
		Nama:      "Baji",
		Energi:    100,
		Jabatan:   "CEO",
		Gaji:      50000000,
		IsActive:  true,
		IsMarried: false,
		Tunjangan: 40000000,
		Usia:      21,
		Skills:    []string{"Javascript", "ReactJS", "Tailwindcss", "Axios", "Golang"},
  }
	fmt.Println(&pegawaiBaru)
	kerjaLembur(&pegawaiBaru)
	fmt.Println(&pegawaiBaru.Energi)

}
