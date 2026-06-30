package main

import (
	"fmt"
)

func latihanTiga() {

	var totalTagihan int = 7000
	var tresholdUp = 100000
	var tresholdDown = 50000
	var statusPesanan = "ready"

	// ? Logic
	if totalTagihan > tresholdUp {
		fmt.Printf("Dapet discount 10%% \n")
	} else if totalTagihan > tresholdDown {
		fmt.Printf("Dapet discount 5%% \n")
	} else {
		fmt.Print("Tidak dapat discount \n")
	}

	switch statusPesanan {
	case "coocking":
		fmt.Print("Koki sedang memasak dengan api besar!.\n")
	case "ready":
		fmt.Print("Ting! Pesanan siap dihidangkan ke pelanggan.\n")
	case "pending":
		fmt.Print("Pesanan sedang disiapkan, koki sedang cuci tangan.\n")
	default:
		fmt.Print("Status pesanan tidak valid.\n")
	}
}
