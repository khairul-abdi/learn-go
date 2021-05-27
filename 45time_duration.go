package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	waktu := time.Duration(5) * time.Second
	//Gunanya untuk membuat satuan sesuai yang diinginkan (contoh: jika ingin ke detik)
	fmt.Println("WAKTU : ", waktu)

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())
}
