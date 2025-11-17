package main

import "fmt"

func main() {
	// Percabangan - IF Statements
	// Dalam hidup, selalu ada pilihan. Seperti itupula di dalam pemrograman. Setiap alurnya, pasti mempunyai percabangan

	// Percabangan IF
	// IF statement adalah 1 percabangan. Kondisi ini hanya menjalankan, jika syaratnya terpenuhi

	num := 3
	if num > 2 {
		// kondisi ini akan dipanggil jika
		// syarat terpenuhi
		fmt.Println("Num lebih besar dari 2")
	}

	// akan dijalankan setelah program didalam selesai
	fmt.Println("Program selesai !")

	// Percabangan Else
	// Percabangan ini mempunyai 2 kondisi. Yaitu, kondisi jika terpenuhi,
	// dan kondisi jika tidak terpenuhi. Jika kondisi terpenuhi,
	// maka akan masuk kedalam block code IF dan jika tidak,
	// akan masuk kedalam block code ELSE

	number := 2
	if number > 2 {
		// kondisi ini akan dipanggil jika
		// syarat terpenuhi
		fmt.Println("Num lebih besar dari 2")
	} else {
		// kondisi ini akan dipanggil jika
		// kondisi pertama tidak terpenuhi
		fmt.Println("Num lebih kecil dan sama dari 2")
	}

	// akan dijalankan setelah program didalam selesai
	fmt.Println("Program selesai !")

	age := 18
	var hasKTP bool
	if age >= 17 {
		hasKTP = true
	} else {
		hasKTP = false
	}

	fmt.Printf("Apakah saya mempunyai KTP ? : %v", hasKTP)
	fmt.Println("")

	// bagaimana jika kita ingin membuat banyak percabangan
	// Percabangan IF - ELSE IF - ELSE
	nilai := 80
	var grade string

	if nilai > 80 {
		grade = "A"
	} else if nilai > 70 {
		grade = "B"
	} else if nilai > 60 {
		grade = "C"
	} else if nilai > 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Println("Nilai saya adalah : " + grade)

	// Kali ini, dia akan melakukan pengecekan terhadap nilainya.
	// Jika memenuhi syarat pertama, maka grade = A. Jika tidak,
	// maka kode program akan melakukan pengecekan kepada block else if,
	// dan dilanjutkan seterusnya

	// Ada berapa jumlah else if yang bisa dibuat? Bebas. Tergantung situasi dan kondisi

	// Nested IF (IF di dalam IF)
	gender := "man"
	age2 := 18

	var canWork bool

	if gender == "man" {
		if age2 >= 18 {
			canWork = true
		} else {
			canWork = false
		}
	} else {
		if age2 >= 20 {
			canWork = true
		} else {
			canWork = false
		}
	}

	if !canWork {
		fmt.Println("Saya tidak boleh bekerja")
	} else {
		fmt.Println("Saya boleh bekerja")
	}

	// Latihan
	/**
	  Silahkan buat lakukan pengecekan apakah suatu variable bernilai ganjil atau genap.
	  Contoh :
	  num = 4 => genap
	  num = 5 => ganjil

	  Clue :
	  Menggunakan modulus, dan percabangan
	*/

	// var numMod int

	// if numMod%2 == 0 {
	// 	fmt.Println("Bilangan Ganjil")
	// }

	for num := 1; num <= 5; num++ {
		if num%2 == 0 {
			fmt.Println("Bilangan Genap -> ", num)
		} else {
			fmt.Println("Bilangan Ganjil -> ", num)
		}
	}

}
