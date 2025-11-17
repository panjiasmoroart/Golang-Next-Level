package main

import "fmt"

func main() {
	// Tipe Data
	// Pada Go, ada beberapa tipe data yaitu numerik, string, dan boolean

	// Tipe data string
	// string adalah tipe data yang digunakan untuk membuat sebuah teks. Tipe data ini dicirikan dengan simbol petik dua (")
	var teks string = "Ini adalah sebuah teks"
	fmt.Println(teks)

	// kita bisa menggunakan backtiks (`). Kelebihannya adalah, semua hal yang kita masukkan didalam backtiks,
	// tidak akan di escape. Termasuk \n, kutip dua, kutip satu, baris baru, dan lain lain
	var wordBacktik string = `
		Ini adalah "teks".
		Dan ini dari 'baris baru'.
		Okeee...
	`
	fmt.Println(wordBacktik)

	// Tipe data numerik - Non Decimal (Integer)
	// Pada tipe data integer, ada 2 tipe yaitu signed dan unsigned. signed adalah tipe data
	// yang bisa negatif dan positif sedangkan unsigned hanya menampung variable positif
	// unsigned diawali dengan uint
	// signed diawali dengan int

	// Tips : Jangan sembarangan ngasih tipe data, karena akan berdampak pada alokasi memori.
	// Hal ini di perlukan agar memori dapat bekerja secara optimal
	var ageUint8 uint8 = 22
	var bigNumberInt64 int64 = -9223372036854775807
	fmt.Println(ageUint8)
	fmt.Println(bigNumberInt64)

	// Jika kita membuat variable var age uint64 = 22, bisa saja.
	// Namun memori akan mengalokasikan sebanyak 2^64 (2 pangkat 64),
	// yang mana kita hanya butuh cuma 2^8 (256). Hal ini akan menyebabkan pemborosan pada penggunaan memori

	// Dan juga, jika kita membuat value diluar dari kemampuannya, maka akan muncul error. Contoh :
	// var age uint8 = 256 // maksimal uint8 adalah 255
	// fmt.Printf("%d\n", age)
	// cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)

	//  jika ingin menampilkan bilangan numerik non decimal, bisa menggunakan %d pada fmt.Printf().

	// Tipe data numerik - Decimal (Float)
	// Tidak semua bilangan itu bulat. Nantinya kita juga akan berhadapan dengan bilangan pecahan.
	// Oleh sebab itu, di Go ada 2 jenis yaitu float32 dan float64. Perbedaannya ada pada
	// kapasitas cakupan dari nilai yang ditampungnya
	var contohDecimal float32 = 1.66
	fmt.Printf("%f\n", contohDecimal)   // menampilkan bilangan desimal %f (default 6 angka dibelakang koma)
	fmt.Printf("%.1f\n", contohDecimal) // menampilkan 1 angka decimal
	fmt.Printf("%.3f\n", contohDecimal) // maka akan menampilkan 3 angka decimal
	// %.nf secara otomatis akan melakukan pembulatan ke bilangan terdekat

	// Boolean (bool)
	// Boolean hanya mempunyai 2 data, yaitu true dan false.
	// Tipe data ini biasanya akan kita gunakan pada kondisi percabangan if else
	var checkExists bool = true
	fmt.Println(checkExists)

}
