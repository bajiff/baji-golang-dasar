package main

import "fmt"

func latihanDua() {
	const hargaNasiGoreng int = 25000
	var stokBeras int = 50
	pelanggan := hargaNasiGoreng * 3

	checkStokBeras := stokBeras > 0
	fmt.Printf("Pelanggan sekarang %v\n", pelanggan)
	fmt.Printf("Stok beras: %d, apakah stok beras lebih dari 0? %t\n ", stokBeras, checkStokBeras)
}
