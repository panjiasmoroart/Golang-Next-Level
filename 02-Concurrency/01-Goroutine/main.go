package main

import (
	"fmt"
	"runtime"
	"time"
)

func speak(total int, message string) {
	for i := 0; i < total; i++ {
		fmt.Println("Speak ", message, "- (", i+1, ")")
	}
}

func main() {
	// set max processor yang akan digunakan
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	// running in new goroutine
	go speak(2, "Goroutine 1")
	// running in new goroutine
	go speak(2, "Goroutine 2")
	// running in new goroutine
	go speak(2, "Goroutine 3")

	// running in main goroutine
	speak(2, "Hello")

	// kita menggunakan time.Sleep() untuk melakukan blocker agar proses yang berjalan
	// di asynchronous bisa segera selesai sebelum main goroutine selesai

	// Cara seperti itu sangat tidak dianjurkan, karena tidak flexible. Misalnya, saat kita melakukan sleep 1 detik,
	// ternyata goroutine nya jalan di 2 detik. Artinya, kita tidak akan mendapatkan hasil dari goroutine nya
	//  muncul Waitgroup yang berfungsi untuk melakukan waiting secara otomatis, sampai async nya selesai.
	time.Sleep(1 * time.Second)
}
