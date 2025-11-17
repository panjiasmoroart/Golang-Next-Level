package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

// menambahkan parameter WaitGroup
// dengan value pointer
func speak(total int, message string, wg *sync.WaitGroup) {
	// check apakah waitgroup nil atau tidak
	if wg != nil {
		// proses untuk menandakan bahwa goroutine telah selesai
		defer wg.Done()
	}

	for i := 0; i < total; i++ {
		fmt.Println("Speak ", message, "- (", i+1, ")")
	}
}

// Atau jika ingin tidak mengubah function speak nya,
func speak3(total int, message string) {
	for i := 0; i < total; i++ {
		fmt.Println("Speak", message, "- (", i+1, ")")
	}

}

// bisa juga dengan cara sepert ini
func sampleWaitGroup2() {
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go speak(2, fmt.Sprintf("Speak Goroutine %v", i+1), &wg)
	}
	speak(2, "Hello", nil)

	wg.Wait()
	fmt.Println("Doneee")
}

func sampleWaitGroup3() {
	// Pertama, kita melakukan inisialisasi
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	for i := 0; i < 3; i++ {
		//  kita perlu menambahkan sebuah state untuk menandakan jumlah goroutine yang akan kita jalankan
		wg.Add(1)
		// generate message diatas
		msg := fmt.Sprintf("Speak Goroutine %v", i+1)

		// gunakan IIFE
		go func(text string) {
			speak3(2, text)
			// untuk menandakan bahwa sebuah goroutine itu telah selesai
			// Method ini berfungsi untuk mengurangi jumlah goroutine yang aktif.
			// Dibalik layarnya, method ini sebenarnya melakukan pengurangan dengan cara
			// wg.Add(-1)
			wg.Done()
			// Jadi, misalnya kita declare bahwa ada 5 goroutine aktif, saat goroutine tersebut
			// telah mengeksekusi wg.Done(), maka akan tinggal 4 goroutine aktif dan begitu seterusnya.
		}(msg)
	}
	speak3(2, "Hello")

	wg.Wait()
	fmt.Println("Doneee")
}

func calculate(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func sampleWaitGroup4() {
	wg := sync.WaitGroup{}
	total := 0
	nums := []int{1, 2, 3, 4, 5}
	wg.Add(1)
	go func(numbers ...int) {
		// nilai total akan di overwrite disini
		total = calculate(numbers...)
		wg.Done()
	}(nums...)

	// prose menunggu sampai goroutine selesai
	wg.Wait()

	// menampilkan nilai total
	log.Println("sampleWaitGroup4 ", total)
}

func main() {
	// membuat object waitgroup
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	// memberitahu, jumlah goroutine yang akan dibuat
	wg.Add(3)

	go speak(2, "Goroutine 1", &wg)
	go speak(2, "Goroutine 2", &wg)
	go speak(2, "Goroutine 3", &wg)

	// karena disini kita tidak membutuhkan goroutine,
	// jadi value waitgroupnya kita isi nil
	speak(2, "Hello", nil)

	// proses menunggu sampai semua waitgroup telah di `done`-in
	wg.Wait()
	fmt.Println("Doneee")

	fmt.Println("======================")

	sampleWaitGroup2()

	fmt.Println("======================")

	sampleWaitGroup3()

	fmt.Println("======================")

	sampleWaitGroup4()
}
