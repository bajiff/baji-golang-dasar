package main

import "fmt"

func latihanLima() {
	// ? Array
	var prinsipBaji = [4]string{"Do'a", "Usaha", "Ikhtiar", "Tawakal"}
	fmt.Println("Array")
	fmt.Println(prinsipBaji)
	fmt.Printf("%s\n", prinsipBaji)
	fmt.Printf("%#v\n", prinsipBaji)
	fmt.Println("")

	// ? Slice
	hobiBaji := []string{"Coding", "Golang", "Badminton", "Joging"}
	fmt.Println("Slice")
	fmt.Println(hobiBaji)
	fmt.Printf("%s\n", hobiBaji)
	fmt.Printf("%#v\n", hobiBaji)
	fmt.Println("")

	// ? Map
	bajiWealth := map[string]int{"Motor": 1, "Sepatu": 2, "Sendal": 1, "Laptop": 1, "Keyboard": 1, "Smartphone": 1}
	fmt.Println("MAP")
	fmt.Println(bajiWealth)
	fmt.Printf("%v\n", bajiWealth)
	fmt.Printf("%#v\n", bajiWealth)
	fmt.Println("")

	// ? For Range
	kumpulanBumbu := []string{"Merica", "Lengkuas", "Kayu Manis", "Bawang Putih", "Bawang Merah", "Jahe", "Mesiu", "Uranium", "H2O2"}

	for i, namaBumbu := range kumpulanBumbu {
		fmt.Printf("Bumbu ke-%d adalah %s\n", i, namaBumbu)
	}
	
	pesanMejaSatu := []string{"Ketoprak","Mie Goreng", "Mie Samyang"}
	pesanMejaSatu = append(pesanMejaSatu, "Martabak Ramona")
	
	for i, namaPesanan := range pesanMejaSatu {
		fmt.Printf("Nama Pesanan ke ke-%d Adalah: %s\n", i+1, namaPesanan)
	}
	
	daftarHarga := map[string]int{"Seblak":10000,"Capcai":2000,"Es Teh":1500, "Makaroni Basah": 5000}
	
	fmt.Printf("Harga %s adalah: %d\n","Seblak", daftarHarga["Seblak"])

}
