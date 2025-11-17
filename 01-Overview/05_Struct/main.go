package main

import "fmt"

// Struct
// adalah sebuah tipe data struktur yang berisikan koleksi dari tipedata. Struct sangat berguna untuk mengelompokkan data.
// Jadi, didalam sebuah struct akan ada property dan atau method yang dibungkus sehingga menciptakan sebuah tipe data baru

// Pada Go, tidak ada mengenal class. Namun, konsep struct ini mirip dengan class pada bahasa OOP lainnya

type Fruit struct {
	// mempunyai 2 Property Name & Weight
	Name   string
	Weight int
}

func main() {
	var fruit1 Fruit
	fruit1.Name = "Apple"
	fruit1.Weight = 1

	fruit2 := Fruit{
		Name:   "Mango",
		Weight: 2,
	}

	var fruit3 = Fruit{
		Name:   "Banana",
		Weight: 1,
	}

	fruit4 := Fruit{"Lemon", 4}

	fmt.Println(fruit1)
	fmt.Println(fruit2)
	fmt.Println(fruit3)
	fmt.Println(fruit4)

	// Untuk menggunakan struct tersebut, ada 4 cara yang mana telah dibuat pada fruit1 ... fruit4.

	// Pada fruit1, melakukan inisialisasi untuk menampung object dari struct Fruit. Sedangkan isinya akan di set kemudian.

	// Pada fruit2 dan fruit3, melakukan inisialisasi langsung dengan memberikan value kepada

	// property nya. Hal ini tidak harus berurutan, kita bisa juga melakukannya dengan cara seperti ini :
	// var fruit5 = Fruit{
	// 	Weight: 1,
	// 	Name:   "Banana",
	// }

	// Pada fruit4, mirip dengan fruit2 dan fruit3. Namun pada fruit4, kita langsung memasukkan value tanpa memberikan
	// property. Hal ini secara otomatis, akan memasukkan setiap valuenya kepada property secara berurutan
}
