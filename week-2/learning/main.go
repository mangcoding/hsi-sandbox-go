package main

import "fmt"

type Siswa struct {
	name    string
	age     int
	address string
	phone   string
	email   string
	point   int
}

func (s Siswa) topUpPoint(point int) int {
	s.point += point
	return s.point
}

func main() {
	// learning struct
	siswa := Siswa{
		name:    "John",
		age:     20,
		address: "Jakarta",
		phone:   "081234567890",
		email:   "john@example.com",
		point:   100,
	}

	fmt.Println(siswa)

	//passing by value, so basically it's a copy of the original value
	var point = siswa.topUpPoint(100)

	fmt.Println(point)
	fmt.Println(siswa.point)

	// learning pointer
	var nugraha *Siswa = &siswa
	nugraha.point = 1000

	// Basically it will print 1000 because we change the point of the siswa struct
	fmt.Println(siswa.point)

}
