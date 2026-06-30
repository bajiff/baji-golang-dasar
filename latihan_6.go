// ! latihan_6.go
package main

import (
	"fmt"
)

func buatPesanan(namaMenu string, jumlah int) string {
	return fmt.Sprintf("Pesanan Anda: %d porsi %s sedang disiapkan.\n", jumlah, namaMenu)
}

func hitungPajak(totalHarga int) int {
	tax := totalHarga * 10 / 100
	totalHargaSetelahTax := tax + totalHarga
	fmt.Printf("Total kembalian: %d\n", totalHargaSetelahTax)
	return totalHargaSetelahTax
}

func latihanEnam() {
	fmt.Println(buatPesanan("Ketoprak Torium", 10))
	hitungPajak(50000)
}
