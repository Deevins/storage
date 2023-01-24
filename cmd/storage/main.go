package main

import (
	"fmt"
	"github.com/deevins/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("HEllo"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file uploaded \n", file)

	fileByID, _ := st.GetByID(file.ID)

	fmt.Println("file found", fileByID)
}
