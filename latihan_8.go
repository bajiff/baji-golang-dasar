// ! latihan_8.go
package main

import "fmt"

type Pekerja struct {
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

type RobotResto struct {
	Nama        string
	Baterai     int
	Model       string
	FungsiUtama string
}

type KaryawanResto interface {
	Perkenalan() string
	kerjaNyantuy(kadarKortisol int) int
	GetBaterai()
	GetEnergi()
}

func (p *Pekerja) kerjaNyantuy(kadarKortisol int) int {
	p.Energi -= kadarKortisol
	fmt.Printf("Pekerja %s, sisa energi: %d - %d = %d\n",p.Nama, p.Energi+kadarKortisol, kadarKortisol, p.Energi)
	return p.Energi
}

func (p *Pekerja) Perkenalan() string {
	return fmt.Sprintf("Halow nama saya %s, jabatan saya %s\n", p.Nama, p.Jabatan)
}

func (p *Pekerja) GetEnergi() {
	fmt.Printf("Pekerja %s, sisa energi: %d\n",p.Nama, p.Energi)
}

func (p *Pekerja) GetBaterai() {
	fmt.Printf("Nama: %s tidak punya baterai\n", p.Nama)
}

func (r *RobotResto) kerjaNyantuy(kadarKortisol int) int {
	r.Baterai -= (kadarKortisol * 2)
	fmt.Printf("Robot %s, sisa baterai: %d - %d = %d\n", r.Nama, r.Baterai+(kadarKortisol*2), kadarKortisol*2, r.Baterai)
	return r.Baterai
}

func (r *RobotResto) GetBaterai() {
	fmt.Printf("Robot %s, sisa baterai: %d\n",r.Nama, r.Baterai)
}
func (r *RobotResto) GetEnergi() {
	fmt.Printf("Robot: %s tidak punya Energi\n", r.Nama)
}

func (r *RobotResto) Perkenalan() string {
	return fmt.Sprintf("Halow saya %s, Model: %s, Sisa baterai: %d\n", r.Nama, r.Model, r.Baterai)
}

func latihanDelapan() {
	priaTampan := Pekerja{
		Nama:      "Mr. Baji Ganteng",
		Energi:    100,
		Jabatan:   "Chief Executive Officer",
		Gaji:      7000000,
		IsActive:  true,
		IsMarried: false,
		Tunjangan: 45000000,
		Usia:      21,
		Skills:    []string{"Javascript", "ReactJS", "Tailwindcss", "Axios", "Golang", "Laravel", "PostgreSQL", "RabbitMQ", "Redis", "Public Speaking"},
	}

	robotGacor := RobotResto{
		Nama:        "Mewky",
		Baterai:     8000,
		Model:       "mewky-1.0",
		FungsiUtama: "Untuk Membantu",
	}

	var majikan KaryawanResto = &robotGacor
	fmt.Printf("%s\n",majikan.Perkenalan())
	fmt.Printf("%d\n",majikan.kerjaNyantuy(30))
	majikan.GetEnergi()
	majikan.GetBaterai()
	fmt.Println("")


	var karyawan KaryawanResto = &priaTampan
	fmt.Printf("%s\n",karyawan.Perkenalan())
	fmt.Printf("%d\n",karyawan.kerjaNyantuy(30))
	karyawan.GetEnergi()
	karyawan.GetBaterai()


	// daftarPekerja := []KaryawanResto{&priaTampan, &robotGacor}

	// for i, karyawan := range daftarPekerja {
	// 	fmt.Println("==============================")
	// 	fmt.Printf("No-%d, Nama: %s ", i+1, karyawan.Perkenalan())
	// 	karyawan.kerjaNyantuy(10)
	// 	fmt.Println("==============================")
	// }
}
