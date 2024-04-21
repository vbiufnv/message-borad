package main

import (
	"log"
	"message-borad/api"
	"message-borad/dao"
)

func main() {
	err := dao.ConDB()
	if err != nil {
		log.Fatal(err)
	}

	defer dao.DB.Close()

	api.SetRouter()

}
