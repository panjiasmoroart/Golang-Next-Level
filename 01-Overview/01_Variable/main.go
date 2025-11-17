package main

import "fmt"

// Variable Global
// Seperti yang telah dijelaskan sebelumnya, variable global adalah sebuah variable yg bisa di akses dari dalam blok kode
var memberName string = "Variable Global"

func main() {
	// Deklarasi secara explisit dapat menggunakan var
	var name string
	name = "Panji Asmoro"

	var age int = 22

	fmt.Printf("Halo, perkenalkan namaku %s dan aku berumur %d tahun\n", name, age)

	// Deklarasi secara implisit tidak memerlukan keyword var, Kita cukup menambahkan : sebelum =
	nameImpl := "Implisit"
	ageImpl := 26

	fmt.Printf("Halo, perkenalkan namaku %s dan aku berumur %d tahun\n", nameImpl, ageImpl)

	// Selain var, kita bisa menggunakan keyword make dan new untuk deklarasi sebuah variable
	// Keyword new akan membuat sebuah pointer dari tipe data yang digunakan
	newPointer := new(string)
	// arterisk menghasilkan sebuah value, maka kita perlu me-replace value
	// tersebut dengan value yang kita inginkan (sesuai dengan tipe data sebelumnya)
	*newPointer = "Aku Pointer"

	fmt.Printf("Ini adalah pointer :%v dan ini adalah valuenya : %s", newPointer, *newPointer)

	// Keyword make
	// Mirip dengan new, cuma make tidak mereference ke alamat memori.
	// Keyword ini digunakan untuk membuat variable tertentu, seperti
	// Channel, Slice, Map, Array

	// hal menari dari go bisa melakukan deklarasi variable, tidak hanya 1 variable
	//  var, hanya berlaku untuk variable yang mempunyai tipe data yang sama
	var firstPerson, lastPerson string = "Muhammad", "Ikhsan"
	//  jika kita ingin deklarasi variable yang berbeda sekaligus
	bookName, price := "GolangAdvanced", 150000
	// deklarasi multiple variable, itu tidak terbatas hanya untuk 2 variable. Kita bisa membuat lebih dari itu
	// name, age, isSingle := "NewBie", 22, true
	fmt.Println("")
	fmt.Println(firstPerson, lastPerson)
	fmt.Println(bookName, price)

	// Scope Variable -> ruang lingkup (scope) dari sebuah variable
	// Local Variable
	// variable yang berada di dalam sebuah blok kode tertentu. Variable
	// jenis ini hanya bisa di akses oleh blok kode tertentu dan anakannya
	// num := 2
	// if num > 1 {
	// 	text := "Lebih besar dari 1"
	// } else {
	// 	text := "Lebih kecil atau sama dengan 1"
	// }
	// fmt.Println(text)
	// Outpt : error: undefined: text
	// Variable text adalah sebuah variable lokal. Jadi hanya bisa di akses pada blok kode yang mendefinisikannya
	// yang bisa mengakses variable text adalah blok kode if dan else. Karena, variable text di definisikan pada blok tersebut
	// Jadi, text tidak bisa di akses dari luar blok kode

	// Untuk memperbaikinya, kita tinggal meletakkan fmt.Println(text) di dalam blok if - else sehingga menjadi :
	num := 2
	if num > 1 {
		text := "Lebih besar dari 1"
		fmt.Println(text)
	} else {
		text := "Lebih kecil atau sama dengan 1"
		fmt.Println(text)
	}

	// Namun, variabel lokal bisa menjadi sebuah variable global, jika dilihat dari scope dibawahnya.
	// Pada contoh sebelumnya, kita membuat sebuah variable num. Nah variable ini adalah variable lokal
	// yang hanya bisa di akses pada blok kode func main(). Namun, num menjadi global
	// bagi blok kode if-else karena mereka bisa mengakses variable num

	number := 2
	if number > 1 {
		fmt.Println(number, "lebih besar dari 1")
	} else {
		fmt.Println(number, "lebih kecil dan sama dari 1")

	}

	fmt.Println(memberName, "dipanggil dari fungsi main")
	sayHello()
}

func sayHello() {
	fmt.Println(memberName, "dipanggil dari fungsi sayHello")
}
