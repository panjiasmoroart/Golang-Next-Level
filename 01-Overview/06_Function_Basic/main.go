package main

import "fmt"

// alurnya sepertin ini
// func main() ==[called]=> func cetak() ==[called]=> func hello()

func main() {
	// Function adalah sekumpulan baris kode yang dibungkus dengan nama tertentu sehingga bisa digunakan berulang kali
	// Dengan menggunakan fungsi, kita bisa menerapkan prinsip DRY yaitu Dont Repeat Your Self

	// Cara Membuat Function
	// Di bahasa pemrograman Go, secara tidak langsung kita sebenarnya telah mendefinisikan sebuah fungsi, yaitu main.
	// Hal ini karena fungsi main adalah fungsi utama dan fungsi yang akan dijalankan pertama kali oleh Go

	cetak()
}

func cetak() {
	fmt.Println("fungsi cetak() dipanggil")
	hello()
	// cetakNama("Panji Asmoro")
	fmt.Println("========================")

	var names = []string{"Rey", "Jo", "Panji"}
	var ages = []int{10, 13, 20}
	for _, name := range names {
		cetakNama(name)
	}

	for _, age := range ages {
		cetakUmur(age)
	}

	fmt.Println("========================")
	for _, name := range names {
		cetakMerge(name)
	}

	for _, age := range ages {
		cetakMerge(age)
	}

	fmt.Println("========================")
	for _, name := range names {
		cetakNama2(name)
	}

	for _, age := range ages {
		cetakUmur2(age)
	}

	fmt.Println("========================")
	value1 := 10
	value2 := 15
	total := addition(value1, value2)
	cetakInterface(total)

	fmt.Println("========================")
	fmt.Println(combineString("Hello", "World"))
}

func hello() {
	fmt.Println("Hello World")
}

// Parameter
// Pada fungsi, kita bisa memberikan parameter didalamnya.
// Caranya, dengan memberikan parameter pada tanda ()
func cetakNama(text string) {
	fmt.Println(text)
}

func cetakUmur(age int) {
	fmt.Println(age)
}

// fungsi dibuat agar bisa melakukan prinsip DRY ya? Kalau dilihat kembali, bukannya
// fungsi cetakNama dan cetakUmur melakukan hal yang sama, yaitu fmt.Println() ya
// dalam behaviour-nya, maka kita bisa menggabungkannya menjadi 1 fungsi
func cetakMerge(arg interface{}) {
	fmt.Println(arg)
}

// Secara singkat, interface adalah sebuah tipe data yang nilai defaultnya adalah nil. Interface
// akan berisi jika sudah diberi nilai. Dan interface bisa menyesuaikan dengan value yang dimasukinya

// Contoh diatas, itu bisa kita gabungkan jika kode di dalam fungsinya sama.
// Namun jika berbeda, maka kita tidak disarankan untuk menggabungkannya contoh :
func cetakNama2(text string) {
	fmt.Println(text)
}

func cetakUmur2(age int) {
	fmt.Println(age, "tahun")
}

// Fungsi yang mengembalikan Nilai
// Jika fungsi sebelumnya, kita hanya menampilkan sesuatu, nah sebuah fungsi juga bisa mengembalikan sebuah nilai. Contoh :
func addition(num1 int, num2 int) int {
	return num1 + num2
}

func cetakInterface(arg interface{}) {
	fmt.Println(arg)
}

// Pada fungsi addition diatas, kita memberikan sebuah nilai sebagai return nya. Formatnya adalah
// func nama_fungsi(params <tipedata>) <tipe_data_return> {
//     return <tipe_data_return>
// }

// Jadi, kita perlu mendefinisikan tipedata apa yang akan kita kembalikan.
// Jika kita mengembalikan sebuah string, maka tipe datanya adalah sebuah string
func combineString(str1 string, str2 string) string {
	return str1 + " " + str2
}
