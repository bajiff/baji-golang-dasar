// ! latihan_4.go
package main

import "fmt"

func latihanEmpat() {
	for i := 1; i <= 10; i++ {
		if i == 4 {
			fmt.Print("Pesanan 4 dibatalkan pelanggan, dilewati!\n")
			continue
		} else if i == 8 {
			fmt.Print("Bahan baku habis di pesanan 8, dapur ditutup!\n")
			break
		} else {
			fmt.Printf("Memasak pesanan ke-%d\n", i)
		}
	}
}
