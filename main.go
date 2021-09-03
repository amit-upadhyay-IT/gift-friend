package main

import (
	"birthdayGiftVoucher/api"
	"log"
)

func main() {

	server := api.NewServer()

	err := server.Start("0.0.0.0:9090")
	if err != nil {
		log.Fatal("can not start server: ", err)
	}
}
