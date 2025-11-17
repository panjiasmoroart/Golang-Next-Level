package main

import "fmt"

func main() {
	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message

	// output => Mobil BMW berwarna Black

	result := showString1(car["name"], car["color"])
	fmt.Println(result) // output => Mobil BMW berwarna Black
}

func showString1(param1, param2 string) string {
	return "Mobil " + param1 + " berwarna " + param2
}
