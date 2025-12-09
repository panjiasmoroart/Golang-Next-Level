package main

import "fmt"

// Untuk membuat sebuah channel, cukup dengan melakukan
// ch := make(chan <tipe_data>)
// setiap channel itu mempunyai tipe datanya sendiri

// Fungsi Channel
// Fungsi channel selain untuk mengambil value dari sebuah function goroutine,
// adalah untuk berkomunikasi antar goroutine.

func main() {
	// membuat sebuah channel string
	process := make(chan string)

	// masukin ke goroutine
	go process1(process)
	go process2(process)
	go process3(process)

	// print value, terima channel
	fmt.Println("value :", <-process)
	fmt.Println("value :", <-process)
	fmt.Println("value :", <-process)

	fmt.Println("===========================")

	// setup 4 channel
	ch1 := make(chan string)
	ch2 := make(chan int)
	doneCh := make(chan bool)
	resultCh := make(chan int)

	// pada proses1, membutuhkan ke 4 channel tersebut
	go process11(ch1, "-", ch2, doneCh, resultCh)

	// pada proses2 hanya membutuhkan 2 channel, yaitu untuk
	// mengambil nilai operator dan untuk mengirim nilai total
	go process22(ch1, ch2)

	// selalu gunakan "done" untuk menandakan bahwa
	// semua proses telah selesai
	if <-doneCh {
		// menampilkan hasil
		fmt.Println("Total", <-resultCh)
	}
}

func process1(process chan string) {
	msg := "Process 1"
	fmt.Println("process running :", msg)

	// kirim channel
	// proses memasukkan sebuah value ke dalam channel
	process <- msg
}

func process2(process chan string) {
	msg := "Process 2"
	fmt.Println("process running :", msg)

	// proses memasukkan sebuah value ke dalam channel
	process <- msg
}

func process3(process chan string) {
	msg := "Process 3"
	fmt.Println("process running :", msg)

	// proses memasukkan sebuah value ke dalam channel
	process <- msg
}

func process11(ch chan string, kind string, total chan int, done chan bool, result chan int) {
	fmt.Println("running process 1")
	// mengisi nilai ch dengan operator
	// pada saat ini, channel 2 sudah bisa beroperasi
	fmt.Println("process 1: send data into channel ch")
	ch <- kind

	// pada proses ini, process 1 akan melakukan bloking sampai
	// data dari channel total diisi oleh process 2
	fmt.Println("process 1: waiting for get data from process 2")
	totalData := <-total
	fmt.Println("process 1: got data from process 2 =>", totalData)
	fmt.Println("process 1: continue process")
	totalData += 100

	fmt.Println("process 1 done")

	// mengirim nilai ke channel done untuk memberitahukan bahwa
	// process goroutine telah selesai
	fmt.Println("process 1: done")
	done <- true
	result <- totalData
}

func process22(ch chan string, total chan int) {
	fmt.Println("running process 2 ...")
	// menerima nilai dari channel ch
	// proses ini blocking, maka akan nunggu sampai
	// ada value yang di assign ke channel ch (ada di proses 1)
	kind := <-ch
	fmt.Println("process 2: got data from channel ch")

	fmt.Println("process 2: sending value to process 1 ...")
	// just several operation
	switch kind {
	case "+":
		total <- 100 + 10
	case "-":
		total <- 100 - 10
	}
	fmt.Println("process 2: done")
}
