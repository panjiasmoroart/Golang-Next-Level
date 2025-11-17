package main

import (
	"fmt"
	"time"
)

func main() {
	// Percabangan - Switch Case Statements
	// Alternatif dari if, Switch case ini mirip seperti IF - ELSE IF - ELSE,
	// namun dibalut dengan lebih simple
	// time := "am"

	// switch time {
	// case "pm":
	// 	fmt.Println("Malam")
	// case "am":
	// 	fmt.Println("Pagi")
	// default:
	// 	fmt.Println("Maaf, waktunya salah !")
	// }

	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Weekend Saturday")
	case time.Sunday:
		fmt.Println("Weekend Monday")
	default:
		fmt.Println("Weekdays")
	}

	//  case time.Saturday dan case time.Sunday itu mempunyai nilai yang sama, kita bisa menyingkatnya :
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekdays")
	}

	// kita bisa membuat switch case seperti ini
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Subuh - Pagi")
	case t.Hour() > 12:
		fmt.Println("Siang - Malam")
	default:
		fmt.Println("Tidak tau jam berapa")
	}

}
