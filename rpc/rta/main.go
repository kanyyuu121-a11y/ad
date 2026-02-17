package main

import (
	advertise "advertiseproject/kitex_gen/advertiseproject/advertise/adservice"
	"log"
)

func main() {
	svr := advertise.NewServer(new(AdServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
