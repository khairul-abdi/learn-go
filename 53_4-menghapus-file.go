package main

import (
	"fmt"
	"os"
)

var path = "/home/klik/Documents/Learn/GOLANG/Learn From Doc/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		fmt.Println("MASUK")
		return
	}

	fmt.Println("==> file berhasil di delete")
}

func main() {
	deleteFile()
}
