package main

import (
	"auth2/router"
	"fmt"
	"auth2/models"
)

func main() {
	DB, err := models.InitDB()
	if err != nil {
		fmt.Printf("DB Error!")
		return
	}
	defer DB.Close()
	r := router.InitRouter()
	r.Run()
}
