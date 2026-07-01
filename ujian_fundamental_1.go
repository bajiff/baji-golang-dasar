package main

import (
	"fmt"
)

type Pesanan struct {
	NamaMenu string
	Jumlah int
	TotalHarga int
}

func ProsesPesanan(namaMenu string, jumlah int, daftarMenu map[string]int) (Pesanan, error) {
	harga, ada := daftarMenu[namaMenu] 
	if !ada {
		return Pesanan{}, fmt.Errorf("Maaf, menu %s tidak tersedia!", namaMenu)
	} 
		harga *= jumlah
		return Pesanan{NamaMenu: namaMenu,Jumlah: jumlah,TotalHarga: harga}, nil
}

func HitungTotalKeranjang(keranjang []Pesanan) int {
	total := 0

	for i, namaKeranjang := range keranjang {
		fmt.Println("Pajak ke-%d namanya: %s", i, namaKeranjang)
	}  
}

func ujianFundamentalSatu(){
	daftarMenu := map[string]int{"Kopi": 1000, "Es Teh": 1500, "Gooday": 4000, "Ketoprak": 20000}
	
	// var keranjang []Pesanan 
	// ProsesPesanan("makan",1,daftarMenu)
	
	cobaPesan, err := ProsesPesanan("Kopi",1,daftarMenu)
	if err != nil {
		fmt.Println("Ada error:",err.Error())
	} else {
		// cobaPesan = append(, daftarMenu).
		fmt.Println("Berhasil: ", cobaPesan)
	}

	cobaPesanEsTeh, err := ProsesPesanan("Es Teh",1,daftarMenu)
	if err != nil {
		fmt.Println("Ada error:",err.Error())
	} else {
		fmt.Println("Berhasil: ", cobaPesanEsTeh)
	}

	cobaPesanNgaco, err := ProsesPesanan("Jus Uranium",1,daftarMenu)
	if err != nil {
		fmt.Println("Ada error:",err.Error())
	} else {
		fmt.Println("Berhasil: ", cobaPesanNgaco)
	}
	
	
	for i, harga := range daftarMenu {
		fmt.Printf("Nama: %s Harga: %d\n", i,harga)
	}
}