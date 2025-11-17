package main

import "fmt"

const (
	Plus  string = "+"
	Minus string = "-"
)

func main() {
	// Function Lanjutan, Fungsi dengan multiple return, Fungsi Closure, Fungsi Variadic, Fungsi sebagai parameter
	var sisi float32 = 4
	// Return pertama akan masuk kedalam variable luas, dan return kedua akan masuk kedalam variable keliling
	luas, keliling := calculate(sisi)
	fmt.Println("Luas dan Keliling dari persegi dengan panjang sisi", sisi)
	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)

	// Sebuah anonymus function dimasukkan kedalam variable bernama calculate. Jadi, calculate adalah sebuah function closure
	// fungsi ini hanya bisa digunakan pada blok fungsi pemanggilnya.
	calculateWithClosure := func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println("=========================")
	fmt.Println(calculateWithClosure(1, 3))
	fmt.Println("=========================")
	sorting := func(arr []int) []int {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i] < arr[j] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
		return arr
	}

	var arr = []int{3, 9, 5, 1, 10, 4, 2}
	fmt.Println("Awal : ", arr)                   // Awal :  [3 9 5 1 10 4 2]
	fmt.Println("Setelah diurut :", sorting(arr)) // Setelah diurut : [1 2 3 4 5 9 10]

	fmt.Println("=========================")
	nilai := []int{5, 2, 8, 6, 10, 8, 9, 5, 8}
	kkm := 7
	total, nilaiLulus := graduate(nilai, kkm)
	fmt.Println("Total yang lulus :", total)
	fmt.Println("Nilai yang lulus :", nilaiLulus())

	fmt.Println("=========================")
	fmt.Println(sum(1, 8, 29, 3, 4, 5, 2, 9, 7))
	// Jika data yang ingin kita masukkan adalah sebuah slice, maka kita perlu menambahkan titik tiga (...) setelah nama variable
	var nums = []int{1, 8, 29, 3, 4, 5, 2, 9, 7}
	totalVariadic := sum(nums...)
	fmt.Println(totalVariadic)

	fmt.Println("=========================")
	dataNums := []float32{1, 4, 2, 9}
	result := calculateVariadic(Plus, dataNums...)
	fmt.Println(result)

}

// Fungsi dengan multiple return
// pada bahasa pemrograman Go, kita bisa mengembalikan lebih dari 1 nilai
// func calculate(sisi float32) (float32, float32) {
// 	luas := sisi * sisi
// 	keliling := sisi * 4
// 	return luas, keliling
// }

// Kita juga bisa memberi variable pada return-nya, seperti:
func calculate(sisi float32) (luas float32, keliling float32) {
	// karena sudah di define pada return, maka tidak perlu melakukan inisialisasi lagi
	// tidak perlu melakukan :=
	luas = sisi * sisi
	keliling = sisi * 4
	return luas, keliling
}

// Fungsi Closure
// Closure adalah sebuah fungsi yang mana dapat disimpan pada sebuah variable. Jadi kita bisa membuat sebuah fungsi
// didalam fungsi, atau mengembalikan fungsi

// Fungsi closure tidak memiliki nama (anonymus function) dan biasanya di deklarasikan di dalam sebuah function.
// Fungsi closure digunakan untuk fungsi yang sekali pakai atau digunakan hanya pada blok kode tertentu.

// contoh closure sebagai return
// Pada fungsi graduate, terdapat 2 nilai balik, yaitu dengan tipe data int dan sebuah closure
// Karena closure adalah fungsi, maka untuk menampilkannya dengan cara : nilaiLulus()
func graduate(nilaiArr []int, limit int) (int, func() []int) {
	var studentsGraduate []int

	for _, arr := range nilaiArr {
		if arr >= limit {
			studentsGraduate = append(studentsGraduate, arr)
		}
	}

	return len(studentsGraduate), func() []int {
		return studentsGraduate
	}
}

// Fungsi Variadic
// Fungsi variadic adalah sebuah fungsi yang mempunyai parameter tidak terbatas dengan tipe yang sama
func sum(num ...int) (total int) {
	fmt.Println(num)

	for _, n := range num {
		total += n
	}

	return total
}

// Fungsi variadic ditandai dengan ... pada parameternya. Jika dilhat, maka parameternya akan menampung semua
// data yang dimasukkan kedalam fungsi tersebut. Lalu parameter akan menjadi sebuah slice, sehingga kita perlu
// melakukan perulangan untuk mengolah data didalamnya.

// Bagaimana kalau kita membutuhkan parameter lain?
// Tidak masalah. Namun yang harus diingat adalah, kamu harus membuat parameter variadic ini pada posisi terakhir. Contoh :
func calculateVariadic(operator string, numbers ...float32) (result float32) {
	for _, num := range numbers {
		if operator == Plus {
			result += num
		} else if operator == Minus {
			result -= num
		}
	}

	return result
}

// Pertama, kita mendefinisikan operator yang di digunakan, yaitu Plus dan Minus.
// Lalu pada function calculate, kita meletakkan 2 parameter, yaitu bertipe string dan 1 lagi merupakan variadic.
// Jika variadic kita letakkan di awal, maka akan menyebabkan error : syntax error: cannot use
